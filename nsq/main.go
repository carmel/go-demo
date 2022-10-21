package nsq

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

const (
	secret = "ssss"
)

type Authorization struct {
	Topic       string   `json:"topic"`
	Channels    []string `json:"channels"`
	Permissions []string `json:"permissions"`
}

type State struct {
	TTL            int             `json:"ttl"` // 每隔几秒查询一次授权服务器
	Authorizations []Authorization `json:"authorizations"`
	Identity       string          `json:"identity"`
	IdentityURL    string          `json:"identity_url"`
	Expires        time.Time
}

func TestFunc(t *testing.T) {

	go func() {
		time.Sleep(6 * time.Second)
		go startConsumer()
		startProducer()
	}()

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		r.ParseForm()
		if sct := r.Form.Get("secret"); sct == secret {
			data, _ := json.Marshal(&State{
				TTL:      5,
				Identity: "1",
				// IdentityURL: "127.0.0.1:8080/auth",
				Authorizations: []Authorization{
					Authorization{Topic: "test", Channels: []string{".*"}, Permissions: []string{"subscribe", "publish"}},
				},
			})
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	})
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}

var url string

func init() {
	//具体ip,端口根据实际情况传入或者修改默认配置
	flag.StringVar(&url, "url", "127.0.0.1:4150", "nsqd")
	flag.Parse()
}

// 生产者
func startProducer() {

	cfg := nsq.NewConfig()
	// cfg.AuthSecret = "dfsfdsf"
	cfg.Set("auth_secret", secret)
	producer, err := nsq.NewProducer(url, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 发布消息
	var i int
	for {
		if err := producer.Publish("test", []byte(fmt.Sprintf("test message %d", i))); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		i++
		time.Sleep(1 * time.Second)
	}
}

// 消费者
func startConsumer() {
	cfg := nsq.NewConfig()
	cfg.Set("auth_secret", secret)
	// cfg.AuthSecret = "dfsfdsf"
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	// 设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body), "from Producer")
		return nil
	}))
	// 连接到单例nsqd
	if err := consumer.ConnectToNSQD(url); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
