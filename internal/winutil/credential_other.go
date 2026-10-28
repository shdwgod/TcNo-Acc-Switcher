//go:build !windows

package winutil

import "errors"

var errCredentialWindowsOnly = errors.New("Windows Credential Manager is only available on Windows")

func ReadCredential(target string) (Credential, error) {
	return Credential{}, errCredentialWindowsOnly
}

func WriteCredential(credential Credential) error {
	return errCredentialWindowsOnly
}

func DeleteCredential(target string) error {
	return errCredentialWindowsOnly
}

func DPAPIProtect(data []byte) ([]byte, error) {
	return nil, errCredentialWindowsOnly
}

func DPAPIUnprotect(data []byte) ([]byte, error) {
	return nil, errCredentialWindowsOnly
}
