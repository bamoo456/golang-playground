package lib

import "testing"

func BenchmarkToJSONArray(b *testing.B) {
	test := struct {
		data interface{}
	}{
		[]map[string]interface{}{
			map[string]interface{}{
				"foo":  "bar",
				"ping": 1,
				"pong": 0,
			},
			map[string]interface{}{
				"header": "1",
			},
		},
	}
	for i := 0; i < b.N; i++ {
		ToJSONArray(test.data)
	}
}
