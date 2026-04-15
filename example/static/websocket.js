// WebSocket connection helper
const wsUrl = '/api/ws';

const socket = new WebSocket(wsUrl);

let req1 = {
        "jsonrpc": "2.0",
        "method": "greet",
        "params": {
            "name": "Tester"
        },
        "id": "1"
    }

// let req2 = []

let req2 = {
        "jsonrpc": "2.0",
        "method": "proxy2",
        "params": {
            "name": "Tester"
        },
        "id": "2"
    }

// Helper functions
function logMessage(message) {
    const messagesDiv = document.getElementById('messages');
    const messageDiv = document.createElement('div');
    messageDiv.textContent = message;
    messagesDiv.appendChild(messageDiv);
}

// Connection opened
socket.addEventListener('open', (event) => {
    console.log('WebSocket connected');
    logMessage('Connected to server');
    document.getElementById('status').innerText = 'Status: Connected';
    socket.send(JSON.stringify(req1))
    socket.send(JSON.stringify(req2))
});

// Listen for messages
socket.addEventListener('message', (event) => {
    console.log('Received:', event.data);
    logMessage('Message: ' + event.data);
});

// Connection closed
socket.addEventListener('close', (event) => {
    console.log('WebSocket disconnected');
    logMessage('Disconnected from server');
    document.getElementById('status').innerText = 'Status: Disconnected';
});

// Handle errors
socket.addEventListener('error', (event) => {
    console.error('WebSocket error:', event);
    logMessage('Error: Unable to connect to server');
    document.getElementById('status').innerText = 'Status: Error';
});
