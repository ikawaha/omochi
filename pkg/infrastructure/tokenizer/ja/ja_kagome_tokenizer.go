package ja

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/service"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var indexable = filter.NewPOSFilter([]filter.POS{
	{"感動詞"},
	{"形容詞"},
	{"動詞"},
	{"名詞"},
	{"副詞"},
}...)

type JaKagomeTokenizer struct {
	t *tokenizer.Tokenizer
}

func NewJaKagomeTokenizer() service.Tokenizer {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	return &JaKagomeTokenizer{t: t}
}

func (tokenizer *JaKagomeTokenizer) Tokenize(ctx context.Context, japaneseContent string) (*[]entities.TermCreate, *errors.Error) {
	tokens := tokenizer.t.Tokenize(japaneseContent)
	var terms []entities.TermCreate
	for _, token := range tokens {
		if indexable.Match(token.POS()) {
			terms = append(terms, *entities.NewTermCreate(token.Surface, nil))
		}
	}
	return &terms, nil
}
