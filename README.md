Based on the code provided for the `tokenizer` repository, here is a comprehensive `README.md`.

I have updated the installation and usage sections to reflect the features found in `prepare.go` (specifically the pipeline construction logic) and `prepare_test.go` (usage with embedded files).

---

# tokenizer

> A configuration-driven, singleton wrapper for the `sugarme/tokenizer` library (BERT/RoBERTa compatible).

`tokenizer` simplifies the usage of HuggingFace-style tokenizers in Go. Instead of manually constructing the tokenization pipeline (Normalizers, PreTokenizers, Decoders) in code, this library hydrates the entire state from a standard `tokenizer.json` configuration file.

It allows you to load tokenizers directly from an `embed.FS` (or any `fs.FS`) and provides a thread-safe singleton pattern to ensure resource-heavy models are loaded only once.

## Features

* **JSON Configuration:** Builds the entire tokenizer pipeline (Model, Normalizer, PreTokenizer, PostProcessor, Decoder, Padding, Truncation) from a single JSON file.
* **Embed-Ready:** Designed to work seamlessly with Go's `embed` package via the `fs.FS` interface.
* **Singleton Pattern:** Includes a thread-safe `GetTokenizer` method that lazily loads the model on first use and caches it for the application's lifetime.
* **Full Pipeline Support:** Automatically configures:
* **Added Tokens:** Special tokens handling.
* **Truncation & Padding:** max length and padding strategies.
* **Normalization:** Unicode normalization, lowercasing, etc.



## Installation

```bash
go get github.com/frogoai/tokenizer

```

## Usage

### 1. Using the Singleton (Recommended)

This method ensures the tokenizer is loaded only once, even if called concurrently. It is ideal for use in API handlers.

```go
package main

import (
	"embed"
	"fmt"
	
	"github.com/frogoai/tokenizer"
)

//go:embed resources/tokenizer.json
var assets embed.FS

func main() {
	// 1. Get the singleton instance.
	// The first call parses the JSON and builds the model.
	// Subsequent calls return the cached instance.
	tk, err := tokenizer.GetTokenizer(assets, "resources/tokenizer.json")
	if err != nil {
		panic(err)
	}

	// 2. Encode text
	// Helper method EncodeSingle simplifies the common case
	en, err := tk.EncodeSingle("Hello, World!")
	if err != nil {
		panic(err)
	}

	// 3. Access results
	fmt.Printf("Tokens: %v\n", en.Tokens)
	fmt.Printf("IDs:    %v\n", en.Ids)
}

```

### 2. Manual Loading

If you need to load multiple different tokenizers (e.g., one for BERT and one for GPT), use `FromFile` directly.

```go
tk, err := tokenizer.FromFile(os.DirFS("./configs"), "bert-base.json")

```

## Configuration Format

The `tokenizer.json` should follow the structure exported by the HuggingFace `tokenizers` library. The wrapper looks for these top-level keys:

* `model`: The underlying model parameters (e.g., WordPiece, BPE).
* `normalizer`: Text cleaning rules.
* `pre_tokenizer`: Splitting rules (e.g., whitespace).
* `post_processor`: Special token insertion (e.g., `[CLS]`, `[SEP]`).
* `decoder`: ID to string conversion rules.
* `added_tokens`: Special vocabulary.
* `truncation`: Max length settings.
* `padding`: Padding strategy.

## Testing

The repository includes tests for validating encoding against expected token counts, useful for ensuring parity with Python implementations.

```bash
go test ./...

```

## License

[MIT](https://www.google.com/search?q=LICENSE)