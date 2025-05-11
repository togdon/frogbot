# Frogbot

Everyone needs a frogbot

1.  **Core Functionality**:
    *   It connects to Discord using a bot token specified in the `BOT_TOKEN` environment variable (found in `frogbot/.envrc` or `frogbot/env`)
    *   The main logic is in `frogbot/bot/main.go`, which sets up the Discord session and event handlers
    *   It listens for and responds to various Discord events

2.  **Event Handling**:
    *   **Message Creation (`frogbot/bot/messageCreate.go`)**:
        *   When a message is sent in a channel the bot has access to, it logs the message
        *   If the bot is mentioned (`@Frogbot`), it provides a response. This can be a "What is best in life?" parody or a random quote from Warcraft II, managed by `frogbot/internal/responses/mentions.go`
        *   If a message contains "frog me" (case-insensitive), the bot sends a random frog image from the `frogbot/frogs/` directory
        *   If a message contains "fun", "frog(s)", or "fact", it sends a fun frog fact, sourced from `frogbot/internal/responses/fff.go`
        *   If a message is written in ALL CAPS (and isn't "LOL" or "WTF"), it uses the Gemini generative AI model (via `GEMINI_API_KEY`) to formulate a response. This is handled in `frogbot/internal/responses/yelling.go`
    *   **Message Reaction Add (`frogbot/bot/messageReactionAdd.go`)**:
        *   If a user reacts to a message with the 🐸 emoji, the bot responds with "Ribbit!"
        *   Other reactions are logged
    *   **Message Reaction Remove (`frogbot/bot/messageReactionRemove.go`)**:
        *   Logs when reactions are removed from messages
