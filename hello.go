package main

import (
	"fmt"
	"math"
	"time"
)

const (
	screenWidth  = 80
	screenHeight = 24
	thetaSpacing = 0.07
	phiSpacing   = 0.02
	radius1      = 1.0  // Radius of the "tube"
	radius2      = 2.0  // Radius of the whole donut
	k1           = 15.0 // Distance from viewer to donut
	k2           = 5.0  // Controls the size of the projection
)

func main() {
	// Create a buffer for the frame
	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		// Initialize output buffer and depth buffer
		output := make([][]rune, screenHeight)
		zBuffer := make([][]float64, screenHeight)
		for i := range output {
			output[i] = make([]rune, screenWidth)
			zBuffer[i] = make([]float64, screenWidth)
			for j := range output[i] {
				output[i][j] = ' '
				zBuffer[i][j] = 0
			}
		}

		// Angle of rotation
		A := 0.0
		B := 0.0

		// Loop through all points on the donut
		for theta := 0.0; theta < 2*math.Pi; theta += thetaSpacing {
			for phi := 0.0; phi < 2*math.Pi; phi += phiSpacing {
				// Calculate 3D coordinates of the point on the donut surface
				cosTheta := math.Cos(theta)
				sinTheta := math.Sin(theta)
				cosPhi := math.Cos(phi)
				sinPhi := math.Sin(phi)

				circleX := radius2 + radius1*cosTheta
				circleY := radius1 * sinTheta

				// Calculate rotated coordinates
				x := circleX*(math.Cos(B)*cosPhi+math.Sin(A)*math.Sin(B)*sinPhi) - circleY*math.Cos(A)*math.Sin(B)
				y := circleX*(math.Sin(B)*cosPhi-math.Sin(A)*math.Cos(B)*sinPhi) + circleY*math.Cos(A)*math.Cos(B)
				z := k1 + circleX*math.Cos(A)*sinPhi + circleY*math.Sin(A)
				ooz := 1 / z // "One over Z" for perspective

				// X and Y projection
				xp := int(float64(screenWidth/2) + k2*ooz*x)
				yp := int(float64(screenHeight/2) - k2*ooz*y)

				// Calculate luminance (brightness)
				luminance := 8 * ((math.Cos(phi) * cosTheta * math.Sin(B)) - (math.Cos(A) * cosTheta * math.Sin(phi)) - (math.Sin(A) * sinTheta) + (math.Cos(B) * math.Cos(A) * sinTheta))
				if luminance > 0 {
					// Depth check and render
					if xp >= 0 && xp < screenWidth && yp >= 0 && yp < screenHeight && ooz > zBuffer[yp][xp] {
						zBuffer[yp][xp] = ooz
						luminanceIndex := int(luminance)
						if luminanceIndex > 11 {
							luminanceIndex = 11
						} else if luminanceIndex < 0 {
							luminanceIndex = 0
						}
						chars := []rune(".,-~:;=!*#$@")
						output[yp][xp] = chars[luminanceIndex]
					}
				}
			}
		}

		// Display the frame
		for i := range output {
			fmt.Println(string(output[i]))
		}

		// Rotate the donut by changing A and B
		A += 0.04
		B += 0.02

		// Add a short delay
		time.Sleep(30 * time.Millisecond)
	}
}
