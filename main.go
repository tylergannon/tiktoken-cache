package main

import (
	"log"
	"os"

	"github.com/pkoukk/tiktoken-go"
)

const (
	GPT432K0613           = "gpt-4-32k-0613"
	GPT432K0314           = "gpt-4-32k-0314"
	GPT432K               = "gpt-4-32k"
	GPT40613              = "gpt-4-0613"
	GPT40314              = "gpt-4-0314"
	GPT4o                 = "gpt-4o"
	GPT4o20240513         = "gpt-4o-2024-05-13"
	GPT4Turbo             = "gpt-4-turbo"
	GPT4Turbo20240409     = "gpt-4-turbo-2024-04-09"
	GPT4Turbo0125         = "gpt-4-0125-preview"
	GPT4Turbo1106         = "gpt-4-1106-preview"
	GPT4TurboPreview      = "gpt-4-turbo-preview"
	GPT4VisionPreview     = "gpt-4-vision-preview"
	GPT4                  = "gpt-4"
	GPT3Dot5Turbo0125     = "gpt-3.5-turbo-0125"
	GPT3Dot5Turbo1106     = "gpt-3.5-turbo-1106"
	GPT3Dot5Turbo0613     = "gpt-3.5-turbo-0613"
	GPT3Dot5Turbo0301     = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo16K      = "gpt-3.5-turbo-16k"
	GPT3Dot5Turbo16K0613  = "gpt-3.5-turbo-16k-0613"
	GPT3Dot5Turbo         = "gpt-3.5-turbo"
	GPT3Dot5TurboInstruct = "gpt-3.5-turbo-instruct"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci003 = "text-davinci-003"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci002 = "text-davinci-002"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextCurie001 = "text-curie-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextBabbage001 = "text-babbage-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextAda001 = "text-ada-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci001 = "text-davinci-001"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3DavinciInstructBeta = "davinci-instruct-beta"
	// Deprecated: Model is shutdown. Use davinci-002 instead.
	GPT3Davinci    = "davinci"
	GPT3Davinci002 = "davinci-002"
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3CurieInstructBeta = "curie-instruct-beta"
	GPT3Curie             = "curie"
	GPT3Curie002          = "curie-002"
	// Deprecated: Model is shutdown. Use babbage-002 instead.
	GPT3Ada    = "ada"
	GPT3Ada002 = "ada-002"
	// Deprecated: Model is shutdown. Use babbage-002 instead.
	GPT3Babbage    = "babbage"
	GPT3Babbage002 = "babbage-002"
)

func main() {
	log.Printf("Downloading tiktoken data files to '%s'", os.Getenv("TIKTOKEN_CACHE_DIR"))
	for _, model := range []string{
		GPT4o,
		GPT4,
		GPT4Turbo,
		GPT3Dot5Turbo,
	} {
		if _, err := tiktoken.EncodingForModel(model); err != nil {
			log.Fatalf("Failed getting encodings for model %s: %v", model, err)
		}
	}
}
