package wyhash

import (
	"testing"

	"github.com/zhangyunhao116/wyhash/internal/wyhashfv1"
)

func TestWyhash(t *testing.T) {
	testSum(t, Sum64, wyhashfv1.Sum64)
	testSumWithSeed(t, Sum64WithSeed, wyhashfv1.Sum64WithSeed)
}
