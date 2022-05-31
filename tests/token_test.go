package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/libra82/netease-im"
)

var client *netease.ImClient

func init() {
	os.Setenv("GOCACHE", "off")
	client = netease.CreateImClient("730dfa83ac3d4fd9172aaaaaaaaaa", "aaaaaaaaaaae", "")
}

func TestToken(t *testing.T) {
	user := &netease.ImUser{ID: "1", Name: "david.tao", Gender: 1}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}

func TestToken2(t *testing.T) {
	user := &netease.ImUser{ID: "2", Name: "joan", Gender: 2}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}

func TestRefreshToken(t *testing.T) {
	tk, err := client.RefreshToken("7")
	if err != nil {
		t.Error(err)
	}
	b, err1 := json.Marshal(tk)
	t.Log(string(b), err1)
}

func Benchmark_SyncMap(b *testing.B) {
	netease.CreateImClient("", "", "")
}
