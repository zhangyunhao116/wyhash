package wyhash

import (
	"testing"

	"github.com/zhangyunhao116/wyhash/internal/wyhashfv3"
)

func TestWyhashV3(t *testing.T) {
	testSum(t, Sum64V3, wyhashfv3.Sum64)
	testSumWithSeed(t, Sum64WithSeedV3, wyhashfv3.Sum64WithSeed)
}
