package tokenizer

import (
	"github.com/sugarme/tokenizer/normalizer"
)

func Normalize(str string) (string, error) {
	n := normalizer.NewBertNormalizer(true, true, true, true)

	res, err := n.Normalize(normalizer.NewNormalizedFrom(CommonString(SanitizeEmail(str))))
	if err != nil {
		return "", err
	}

	return res.GetNormalized(), nil
}
