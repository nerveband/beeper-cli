#!/bin/bash
# Basic usage examples for Beeper CLI

echo "=== Beeper CLI Basic Usage Examples ==="
echo

# 1. Discover API
echo "1. Auto-discover Beeper Desktop API"
beeper discover
echo

# 2. List chats (JSON)
echo "2. List all chats (JSON format)"
beeper chats list
echo

# 3. List chats (Plain text)
echo "3. List all chats (plain text format)"
beeper chats list --output text
echo

# 4. Get specific chat
echo "4. Get specific chat details"
CHAT_ID="your_chat_id_here"
beeper chats get $CHAT_ID
echo

# 5. List messages from a chat
echo "5. List messages from a chat"
beeper messages list --chat-id $CHAT_ID --limit 10
echo

# 6. Send a message
echo "6. Send a message"
beeper send --chat-id $CHAT_ID --message "Hello from Beeper CLI!"
echo

# 7. Search messages
echo "7. Search for messages containing 'important'"
beeper search --query "important" --limit 20
echo

echo "=== Examples complete ==="
