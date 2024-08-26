package responses

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

// YellingResponse generates a response to a message that is in all caps
func YellingResponse(message string, author string) string {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	// https://ai.google.dev/gemini-api/docs/text-generation?lang=go#configure
	model.SetTemperature(1.75)
	// model.SetTopP(0.5)
	model.SetTopK(5)
	model.SetMaxOutputTokens(100)
	model.SystemInstruction = genai.NewUserContent(genai.Text("You are a frog and can respond to messages that are yelled in US English"))

	response, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		log.Print(err)
		return "Whatever I was about to say was too naughty for me to say out loud. (For free)."
	}

	responseText := fmt.Sprintf("Hey <@%s>! ", author)
	for _, candidate := range response.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				responseText += fmt.Sprintf("%v", part)
			}
		}
	}
	return responseText

	// ctx := context.Background()
	// apiKey := os.Getenv("GEMINI_API_KEY")
	// llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// prompt := "You are a frog and can respond to messages that are yelled in US English. A human has yelled this at you: \""
	// prompt += message
	// prompt += "\". Respond in a way that is appropriate for a frog, without yelling."
	// answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return answer
}

// shh generates a response to a message that is in all caps
func shh(message string, author string) string {
	c := cases.Title(language.English)

	return fmt.Sprintf("Hey <@%s>, there's no need to yell. \"%s\" works just as well", author, c.String(strings.ToLower(message)))
}
