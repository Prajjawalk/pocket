package keygenerator

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"sync"

	"github.com/pokt-network/pocket/shared/crypto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var keygenInitOnce sync.Once
var keygen *keyGenerator

type keyGenerator struct {
	privateKeySeed int
}

func GetInstance() *keyGenerator {
	if keygen == nil {
		var privateKeySeed int
		keygenInitOnce.Do(func() {
			privateKeySeed = getPrivateKeySeed()
			keygen = &keyGenerator{
				privateKeySeed: privateKeySeed,
			}
		})
	}
	return keygen
}

func (k *keyGenerator) reset() {
	k.privateKeySeed = getPrivateKeySeed()
}

func (k *keyGenerator) SetSeed(seed int) (teardown func()) {
	k.privateKeySeed = seed
	return func() {
		k.reset()
	}
}

func (k *keyGenerator) Next() (privateKey, publicKey, address string) {
	k.privateKeySeed += 1 // Different on every call but deterministic
	cryptoSeed := make([]byte, crypto.SeedSize)
	binary.LittleEndian.PutUint32(cryptoSeed, uint32(k.privateKeySeed))

	reader := bytes.NewReader(cryptoSeed)
	privateKeyBz, err := crypto.GeneratePrivateKeyWithReader(reader)
	if err != nil {
		panic(err)
	}

	privateKey = privateKeyBz.String()
	publicKey = privateKeyBz.PublicKey().String()
	address = privateKeyBz.PublicKey().Address().String()

	return
}

func getPrivateKeySeed() (privateKeySeed int) {
	rand.Seed(timestamppb.Now().Seconds)
	privateKeySeed = rand.Int()
	return
}