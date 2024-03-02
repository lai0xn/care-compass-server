package gemini

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Content struct {
	Parts []string `json:Parts`
	Role  string   `json:Role`
}
type Candidates struct {
	Content *Content `json:Content`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:Candidates`
}

func Prompt(input string) string {
	// Access your API key as an environment variable (see "Set up your API key" above)
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(
		ctx,
		option.WithAPIKey("AIzaSyBjcYGhJBcYz_4ENvH9wJIgU-9NAgm8xvg"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	res, err := model.GenerateContent(
		ctx,
		genai.Text(input+"don't answer my question if it is not related to the medical field"),
	)
	if err != nil {
		return "i don't know about this"
	}
	marshalResponse, _ := json.MarshalIndent(res, "", "  ")

	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		return "i don't know about this"
	}
	var reply string = ""
	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				reply = reply + part
			}
		}
	}
	return strings.ReplaceAll(reply, "\n", " ")
}
