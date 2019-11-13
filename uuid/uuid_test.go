package uuid

import (
	"testing"
)

func Test_uuid(t *testing.T) {
	InitUUID()
	for i := 0; i < 10000; i++ {
		UUID()
	}
}

func BenchmarkUUID(b *testing.B) {
	InitUUID()

	b.StartTimer()
	// for i := 0; i < 10000; i++ {
	UUID()
	// }
	b.StopTimer()
}
