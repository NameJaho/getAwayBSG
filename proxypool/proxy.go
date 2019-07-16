package proxypool

// 代理实现层
import (
	"fmt"
	"getAwayBSG/configs"
	"github.com/gocolly/colly"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
)

type proxyPool struct {
	proxyURLs []*url.URL
}

func (r *proxyPool) GetProxy(pr *http.Request) (*url.URL, error) {
	// 从配置文件读取代理，可以修改返回，从其他地方获取代理，比如代理池

	if len(r.proxyURLs) > 0 {
		return r.proxyURLs[rand.Intn(len(r.proxyURLs))], nil
	} else {
		return url.Parse(getOneProxy())
	}

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

func getOneProxy() string {
	resp, _ := http.Get("http://45.78.45.70:5015/get/")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	proxy := "http://" + string(body)
	fmt.Println("使用默认代理：" + proxy)
	return proxy
}
