// +build !linux,!windows

package dnscryptproxy

func ServiceManagerStartNotify() error {
	return nil
}

func ServiceManagerReadyNotify() error {
	return nil
}
