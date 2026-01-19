# Beeper Desktop API Documentation

This document describes the Beeper Desktop HTTP API endpoints that this CLI wraps.

## Base URL

Default: `http://localhost:39867`

The Beeper Desktop application exposes a local HTTP API server on port 39867 (or nearby ports).

## Authentication

Currently, the API does not require authentication when accessed from localhost. The Beeper Desktop app itself handles user authentication.

## Endpoints

### Health Check

**GET** `/health`

Check if the API server is running.

**Response:**
```json
{
  "status": "ok"
}
```

### List Chats

**GET** `/chats`

Retrieve all conversations.

**Response:**
```json
[
  {
    "id": "chat_123",
    "name": "John Doe",
    "participants": ["user1", "user2"],
    "last_message": "Hello there",
    "unread_count": 3,
    "updated_at": "2026-01-19T19:00:00Z"
  }
]
```

### Get Chat

**GET** `/chats/{chat_id}`

Retrieve details of a specific chat.

**Parameters:**
- `chat_id` (path): Unique chat identifier

**Response:**
```json
{
  "id": "chat_123",
  "name": "John Doe",
  "participants": ["user1", "user2"],
  "last_message": "Hello there",
  "unread_count": 3,
  "updated_at": "2026-01-19T19:00:00Z"
}
```

### List Messages

**GET** `/chats/{chat_id}/messages`

Retrieve messages from a specific chat.

**Parameters:**
- `chat_id` (path): Unique chat identifier
- `limit` (query): Maximum number of messages (default: 50)

**Response:**
```json
[
  {
    "id": "msg_456",
    "chat_id": "chat_123",
    "sender": "user1",
    "text": "Hello!",
    "timestamp": "2026-01-19T18:00:00Z",
    "type": "text"
  }
]
```

### Send Message

**POST** `/messages/send`

Send a new message to a chat.

**Request Body:**
```json
{
  "chat_id": "chat_123",
  "message": "Hello from CLI!"
}
```

**Response:**
```json
{
  "message_id": "msg_789",
  "success": true
}
```

### Search Messages

**GET** `/search`

Search for messages across all chats.

**Parameters:**
- `query` (query): Search text
- `limit` (query): Maximum results (default: 100)

**Response:**
```json
[
  {
    "id": "msg_456",
    "chat_id": "chat_123",
    "sender": "user1",
    "text": "Search term found here",
    "timestamp": "2026-01-19T18:00:00Z",
    "type": "text"
  }
]
```

## Data Models

### Chat

| Field | Type | Description |
|-------|------|-------------|
| id | string | Unique chat identifier |
| name | string | Display name of chat |
| participants | []string | List of participant identifiers |
| last_message | string | Preview of most recent message |
| unread_count | int | Number of unread messages |
| updated_at | timestamp | Last activity time |

### Message

| Field | Type | Description |
|-------|------|-------------|
| id | string | Unique message identifier |
| chat_id | string | Parent chat ID |
| sender | string | Sender identifier |
| text | string | Message content |
| timestamp | timestamp | When message was sent |
| type | string | Message type (text, image, etc.) |

## Error Handling

The API returns standard HTTP status codes:

- `200 OK`: Success
- `400 Bad Request`: Invalid parameters
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

Error responses include a message:

```json
{
  "error": "Chat not found"
}
```

## Notes

- This documentation is based on the expected Beeper Desktop API structure
- Actual endpoint paths and response formats may vary
- Some endpoints may require additional authentication or headers
- Check Beeper Desktop documentation for authoritative API details
