package uuid

import (
	"testing"
	"time"
)

func Test_uuid(t *testing.T) {
	InitUUID(time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local))
	t.Log(UUID())
}

func BenchmarkUUID(b *testing.B) {
	InitUUID()

	b.StartTimer()
	// for i := 0; i < 10000; i++ {
	UUID()
	// }
	b.StopTimer()
}
