package wxpay

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"
)

type WxPayConfig struct {
	AppId          string
	MchId          string
	AppKey         string
	TradeType      string
	SpbillCreateIp string
	NotifyUrl      string
	TlsConfig      *tls.Config
	Timeout        time.Duration
}

// 支付配置
func NewWxPayConfig(appId, appKey, mchId, tradeType, notifyUrl, spbillCreateIp string, tlsConfig *tls.Config) (config WxPayConfig) {
	config = WxPayConfig{}
	config.AppId = appId
	config.AppKey = appKey
	config.MchId = mchId
	config.TradeType = tradeType
	config.NotifyUrl = notifyUrl
	config.SpbillCreateIp = spbillCreateIp
	config.Timeout = 6
	// 安全证书
	if tlsConfig != nil {
		config.TlsConfig = tlsConfig
	}
	return
}

// 安全证书 导入顺序 cert、key、rootca
func NewWxPayTlsConfig(paths ...string) (tlsConfig *tls.Config, err error) {
	tlsConfig = new(tls.Config)

	var cert tls.Certificate
	cert, err = tls.LoadX509KeyPair(paths[0], paths[1])
	if err != nil {
		return
	}
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)

	if len(paths) >= 3 {
		var pemCerts []byte
		pemCerts, err = ioutil.ReadFile(paths[2])
		if err != nil {
			return
		}

		tlsConfig.RootCAs = x509.NewCertPool()
		tlsConfig.RootCAs.AppendCertsFromPEM(pemCerts)
	}
	return
}
