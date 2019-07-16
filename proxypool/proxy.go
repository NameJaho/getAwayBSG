package proxypool

// 代理实现层
import (
	"getAwayBSG/configs"
	"github.com/gocolly/colly"
	"math/rand"
	"net/http"
	"net/url"
)

type proxyPool struct {
	proxyURLs []*url.URL
}

func (r *proxyPool) GetProxy(pr *http.Request) (*url.URL, error) {
	// 从配置文件读取代理，可以修改返回，从其他地方获取代理，比如代理池

	return r.proxyURLs[rand.Intn(len(r.proxyURLs))], nil
}

func GetProxyPool() (colly.ProxyFunc, error) {
	configInfo := configs.Config()
	var proxyURLs []*url.URL
	if configInfo["proxyList"] != nil && len(configInfo["proxyList"].([]interface{})) > 0 {
		for _, v := range configInfo["proxyList"].([]interface{}) {
			urlLink, err := url.Parse(v.(string))
			if err == nil {
				proxyURLs = append(proxyURLs, urlLink)
			}
		}
	}
	return (&proxyPool{proxyURLs: proxyURLs}).GetProxy, nil
}
