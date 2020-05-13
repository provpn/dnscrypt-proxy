package bridge

import dnscryptproxy "github.com/provpn/dnscrypt-proxy/dnscrypt-proxy"

type BridgeClient struct {
}

// list - print the list of available resolvers for the enabled filters
// listAll - print the complete list of available resolvers, ignoring filters
// json - output list as JSON
// check - check the configuration file and exit
// child - Invokes program as a child process
// showCerts - print DoH certificate chain hashes
// configFilePath - path to dnscrypt-proxy.toml
func InitBridgeClient(list, listAll, json, check, child, showCerts bool, configFilePath string) error {
	if err := dnscryptproxy.InitApp(list, listAll, json, check, child, showCerts, configFilePath); err != nil {
		return err
	}

	return nil
}
