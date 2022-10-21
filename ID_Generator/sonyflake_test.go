package idg

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
)

var sf *sonyflake.Sonyflake

func Testmain(m *testing.M) {
	var st sonyflake.Settings
	st.MachineID = awsutil.AmazonEC2MachineID
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	m.Run()
}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := sf.NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(sonyflake.Decompose(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	w.Write(body)
}

func TestServer(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
