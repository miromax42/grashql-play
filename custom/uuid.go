package custom

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalCustomUUID(b uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		toWrite := fmt.Sprintf("%q", b.String())
		w.Write([]byte(toWrite))
	})
}

func UnmarshalCustomUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.MustParse(v), nil
	default:
		return uuid.Nil, fmt.Errorf("%T is not a UUID", v)
	}
}
