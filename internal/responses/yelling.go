package responses

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/genai"
)

// YellingResponse generates a response to a message that is in all caps
func YellingResponse(message string, author string) string {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Println(err)
	}

	// Uncomment to dump the list of current models...
	// listconfig := &genai.ListModelsConfig{}

	// models, err := client.Models.List(ctx, listconfig)
	// if err != nil {
	// 	log.Println(err)
	// }

	// for _, model := range models.Items {
	// 	fmt.Println(model.Name)
	// }

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("You are a frog. Your name is FrogBot", genai.RoleUser),
		MaxOutputTokens:   int32(100),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(message),
		config,
	)

	if err != nil {
		log.Println(err)
		return ""

	} else {
		return result.Text()
	}
}

// shh generates a response to a message that is in all caps
func shh(message string, author string) string {
	c := cases.Title(language.English)

	return fmt.Sprintf("Hey <@%s>, there's no need to yell. \"%s\" works just as well", author, c.String(strings.ToLower(message)))
}
