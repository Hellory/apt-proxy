package httpcache_test

import (
	"testing"

	. "github.com/soulteary/apt-proxy/pkg/httpcache"
	"github.com/stretchr/testify/require"
)

func TestParsingCacheControl(t *testing.T) {
	table := []struct {
		ccString string
		ccStruct CacheControl
	}{
		{`public, private="set-cookie", max-age=100`, CacheControl{
			"public":  []string{},
			"private": []string{"set-cookie"},
			"max-age": []string{"100"},
		}},
		{` foo="max-age=8, space",  public`, CacheControl{
			"public": []string{},
			"foo":    []string{"max-age=8, space"},
		}},
		{`s-maxage=86400`, CacheControl{
			"s-maxage": []string{"86400"},
		}},
		{`max-stale`, CacheControl{
			"max-stale": []string{},
		}},
		{`max-stale=60`, CacheControl{
			"max-stale": []string{"60"},
		}},
		{`" max-age=8,max-age=8 "=blah`, CacheControl{
			" max-age=8,max-age=8 ": []string{"blah"},
		}},
	}

	for _, expect := range table {
		cc, err := ParseCacheControl(expect.ccString)
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, cc, expect.ccStruct)
		require.NotEmpty(t, cc.String())
	}
}

func BenchmarkCacheControlParsing(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ParseCacheControl(`public, private="set-cookie", max-age=100`)
		if err != nil {
			b.Fatal(err)
		}
	}
}