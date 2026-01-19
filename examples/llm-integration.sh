#!/bin/bash
# LLM-friendly usage examples
# Demonstrates JSON parsing and piping for AI agent integration

echo "=== LLM Integration Examples ==="
echo

# 1. Get chat list as clean JSON for LLM parsing
echo "1. Extract chat IDs and names"
beeper chats list | jq -r '.[] | "\(.id): \(.name)"'
echo

# 2. Find most active chats (by unread count)
echo "2. Find chats with unread messages"
beeper chats list | jq -r '.[] | select(.unread_count > 0) | "\(.name) (\(.unread_count) unread)"'
echo

# 3. Export recent messages to markdown
echo "3. Export recent messages to markdown file"
CHAT_ID="your_chat_id_here"
beeper messages list --chat-id $CHAT_ID --output markdown > messages.md
echo "Exported to messages.md"
echo

# 4. Search and extract specific fields
echo "4. Search for messages and extract sender + text"
beeper search --query "meeting" | jq -r '.[] | "\(.sender): \(.text)"'
echo

# 5. Send message and capture response ID
echo "5. Send message and capture message ID"
MESSAGE_ID=$(beeper send --chat-id $CHAT_ID --message "Status update" | jq -r '.message_id')
echo "Sent message ID: $MESSAGE_ID"
echo

# 6. Get chat participants for group analysis
echo "6. List all participants across chats"
beeper chats list | jq -r '.[] | .participants[]' | sort | uniq
echo

# 7. Batch operations: send to multiple chats
echo "7. Send broadcast message to multiple chats"
CHAT_IDS=("chat1" "chat2" "chat3")
MESSAGE="Broadcast message"
for chat_id in "${CHAT_IDS[@]}"; do
    echo "Sending to $chat_id..."
    beeper send --chat-id $chat_id --message "$MESSAGE"
done
echo

# 8. Generate summary report
echo "8. Generate chat activity summary"
beeper chats list | jq -r '
  "Total chats: \(length)",
  "Unread messages: \([.[] | .unread_count] | add)",
  "Active chats: \([.[] | select(.unread_count > 0)] | length)"
'
echo

echo "=== Examples complete ==="
