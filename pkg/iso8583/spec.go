package iso8583

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/specs"
)

var customSpec *iso8583.MessageSpec

func Init() *iso8583.MessageSpec {
	customSpec = &iso8583.MessageSpec{
		Name:   "Custom ISO 8583 v1987 ASCII",
		Fields: make(map[int]field.Field),
	}

	// Все стандартные поля из базовой спеки копираем
	for k, v := range specs.Spec87ASCII.Fields {
		customSpec.Fields[k] = v
	}

	// Необходимые кастомизируем
	customSpec.Fields[3] = field.NewString(&field.Spec{
		Length:      6,
		Description: "Processing Code",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
	})
	customSpec.Fields[70] = field.NewString(&field.Spec{
		Length:      3,
		Description: "Network Management Information Code",
		Enc:         encoding.ASCII,
		Pref:        prefix.ASCII.Fixed,
	})

	return customSpec
}
