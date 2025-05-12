package uniofound_test

import (
	"github.com/stretchr/testify/require"
	uniofound "go_alg/unio_found"
	"os"
	"testing"
)

const (
	N = 10
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUnion1UnionCount(t *testing.T) {
	unio := uniofound.NewMyUF(N)
	require.Equal(t, N, unio.Count(), "count的数值和N不相等")
}

func TestUnion2UnionUnion(t *testing.T) {
	union := uniofound.NewMyUF(N)
	for i := 1; i < N; i++ {
		union.Union(i-1, i)
	}
	require.Equal(t, 1, union.Count(), "全连接的Union应该为1")
}

func TestUnion3UnionConnect(t *testing.T) {
	union := uniofound.NewMyUF(N)
	require.False(t, union.Connected(1, 2), "1和2在同一个连通分量上面")
	for i := 1; i < N; i++ {
		union.Union(i-1, i)
	}
	require.True(t, union.Connected(1, 2), "1和2在同一个连通分量上面")
}

func TestUnion4UnionQuickFind(t *testing.T) {

}

func TestUnion5UnionQuickUnion(t *testing.T) {

}
