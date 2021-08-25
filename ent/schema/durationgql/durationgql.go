package durationgql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalDuration(t time.Duration) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(int64(t), 10))
	})
}

func UnmarshalDuration(v interface{}) (t time.Duration, err error) {
	i, ok := v.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid type %T, expect string", v)
	}

	return time.Duration(i), nil
}
