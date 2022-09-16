package custom

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(b uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		bytes, _ := b.MarshalBinary()
		w.Write(bytes)
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case uuid.UUID:
		return v, nil
	default:
		return uuid.Nil, fmt.Errorf("%T is not a UUID", v)
	}
}
