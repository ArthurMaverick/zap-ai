package domain_test

// import (
// 	"testing"

// 	"github.com/ArthurMaverick/zap-ai/core/domain"
// 	"github.com/stretchr/testify/require"
// )

// func TestGPTQuestion_IsValid_EmptyText(t *testing.T) {
// 	question := domain.Question{
// 		MessageID: "123",
// 		Text:      "",
// 		Sender:    "Arthur",
// 		Algorithm: domain.Algorithm{
// 			AlgorithmType: "gpt-3",
// 			TokensInput:   100,
// 			TokensOutput:  100,
// 		},
// 	}

// 	isValid, err := question.IsValid()
// 	require.NotNil(t, err)
// 	require.False(t, isValid)
// }

// func TestGPTQuestion_IsValid_EmptyMessageID(t *testing.T) {
// 	question := domain.Question{
// 		MessageID: "",
// 		Text:      "What is the meaning of life?",
// 		Sender:    "Arthur",
// 		Algorithm: domain.Algorithm{
// 			AlgorithmType: "gpt-3",
// 			TokensInput:   100,
// 			TokensOutput:  100,
// 		},
// 	}

// 	isValid, err := question.IsValid()
// 	require.NotNil(t, err)
// 	require.False(t, isValid)
// }

// func TestGPTQuestion_IsValid_EmptySender(t *testing.T) {
// 	question := domain.Question{
// 		MessageID: "123",
// 		Text:      "What is the meaning of life?",
// 		Sender:    "",
// 		Algorithm: domain.Algorithm{
// 			AlgorithmType: "gpt-3",
// 			TokensInput:   100,
// 			TokensOutput:  100,
// 		},
// 	}

// 	isValid, err := question.IsValid()
// 	require.NotNil(t, err)
// 	require.False(t, isValid)
// }

// func TestGPTQuestion_IsValid_EmptyAlgorithmType(t *testing.T) {
// 	question := domain.Question{
// 		MessageID: "123",
// 		Text:      "What is the meaning of life?",
// 		Sender:    "Arthur",
// 		Algorithm: domain.Algorithm{
// 			AlgorithmType: "",
// 			TokensInput:   100,
// 			TokensOutput:  100,
// 		},
// 	}

// 	isValid, err := question.IsValid()
// 	require.NotNil(t, err)
// 	require.False(t, isValid)
// }

// func TestGPTQuestion_IsValid_ZeroTokensInput(t *testing.T) {
// 	t.Run("TokensInput is minor zero", func(t *testing.T) {
// 		question := domain.Question{
// 			MessageID: "123",
// 			Text:      "What is the meaning of life?",
// 			Sender:    "Arthur",
// 			Algorithm: domain.Algorithm{
// 				AlgorithmType: "gpt-3",
// 				TokensInput:   -1,
// 				TokensOutput:  100,
// 			},
// 		}

// 		isValid, err := question.IsValid()
// 		require.NotNil(t, err)
// 		require.False(t, isValid)
// 	})
// }

// func TestGPTQuestion_IsValid_TokensOutput(t *testing.T) {
// 	t.Run("TokensOutput is greater than 150", func(t *testing.T) {
// 		question := domain.Question{
// 			MessageID: "123",
// 			Text:      "What is the meaning of life?",
// 			Sender:    "Arthur",
// 			Algorithm: domain.Algorithm{
// 				AlgorithmType: "gpt-3",
// 				TokensInput:   100,
// 				TokensOutput:  160,
// 			},
// 		}

// 		isValid, err := question.IsValid()
// 		require.NotNil(t, err)
// 		require.False(t, isValid)
// 	})

// }
