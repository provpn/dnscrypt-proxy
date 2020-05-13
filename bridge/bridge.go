package bridge

import dnscryptproxy "github.com/provpn/dnscrypt-proxy/dnscrypt-proxy"

type BridgeClient struct {
}

func InitBridgeClient(list, listAll, json, check, child, showCerts bool) error {
	if err := dnscryptproxy.InitApp(list, listAll, json, check, child, showCerts); err != nil {
		return err
	}

	return nil
}
