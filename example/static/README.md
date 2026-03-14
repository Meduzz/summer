# Static WebSocket Demo

## Files
- `index.html`: Main HTML page with embedded styles
- `websocket.js`: WebSocket connection helper (used by index.html)

## Usage
1. Make sure your WebSocket server is running on `ws://localhost:8080/api/ws`
2. Open `static/index.html` in a web browser
3. You should see connection status updates and any received messages

## Notes
- This is an example-level implementation
- Make sure to adjust the wsUrl in websocket.js if your server URL differs
- The page will show connection errors if the server is not running