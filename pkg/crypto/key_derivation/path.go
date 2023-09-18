package key_derivation

import (
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/crypto/bls"
)

func MnemonicAndPathToKey(seed []byte, path string) (*bls.PrivateKey, error) {
	sk, err := PrivateKeyFromSeedAndPath(seed, path)
	if err != nil {
		return nil, err
	}
	return sk, nil
}
