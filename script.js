let countdownInterval;
let timeRemaining;

function startTimer() {
    const timeInput = document.getElementById("time").value;
    timeRemaining = parseInt(timeInput);

    if (isNaN(timeRemaining) || timeRemaining <= 0) {
        alert("Please enter a valid number of seconds.");
        return;
    }

    document.getElementById("timer-input").classList.add("hidden");
    document.getElementById("timer-display").classList.remove("hidden");
    document.getElementById("timer-options").classList.add("hidden");
    
    countdownInterval = setInterval(updateCountdown, 1000);
    updateCountdown();
}

function updateCountdown() {
    const days = Math.floor(timeRemaining / (60 * 60 * 24));
    const hours = Math.floor((timeRemaining % (60 * 60 * 24)) / (60 * 60));
    const minutes = Math.floor((timeRemaining % (60 * 60)) / 60);
    const seconds = timeRemaining % 60;

    document.getElementById("countdown").textContent = 
        `${days}d ${hours}h ${minutes}m ${seconds}s`;

    if (timeRemaining <= 0) {
        clearInterval(countdownInterval);
        document.getElementById("timer-options").classList.remove("hidden");
        alert("Time's up!");
    } else {
        timeRemaining -= 1;
    }
}

function resetTimer() {
    clearInterval(countdownInterval);
    timeRemaining = parseInt(document.getElementById("time").value);
    document.getElementById("timer-options").classList.add("hidden");
    countdownInterval = setInterval(updateCountdown, 1000);
}

function newTimer() {
    clearInterval(countdownInterval);
    document.getElementById("timer-display").classList.add("hidden");
    document.getElementById("timer-options").classList.add("hidden");
    document.getElementById("timer-input").classList.remove("hidden");
    document.getElementById("time").value = "";
}
