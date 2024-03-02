package services

import "github.com/lai0xn/hackiwna-backend/internal/gemini"

type AiService struct{}

func (s *AiService) Prompt(input string) string {
	reply := gemini.Prompt(input)
	return reply
}
