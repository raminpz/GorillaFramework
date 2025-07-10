function changeMessage() {
    const messageElement = document.getElementById('message');
    const currentMessage = messageElement.textContent;

    const alternateMessages = [
        "Hello, World!",
        "Welcome to our website!",
        "Enjoy your stay!",
        "Have a great day!",
        "Feel free to explore!"
    ]

    let newMessage;
    do {
        newMessage = alternateMessages[Math.floor(Math.random() * alternateMessages.length)];
    } while (newMessage === currentMessage);

    messageElement.textContent = newMessage;
}