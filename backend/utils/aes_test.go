package utils

import "testing"

func TestAESKey(t *testing.T) {
	key, _ := RandAesKey()
	t.Error(key)
}

func TestAESGCMEncode(t *testing.T) {
	key, _ := RandAesKey()
	ciphertext, _ := AESGCMEncode(key, []byte("i love network security"))
	t.Log(key, ciphertext)
}

func TestAESGCMDecode(t *testing.T) {
	key := "8d2582a1ce5637109263b433aa8b0ae7ae2e67ab0aed6f6a3f96a64e8b7c69e6"
	text := "aiWonkMHnHPNzF5wOOVBCTOfkEVgSxvY2tye0+SHxqaARNoSl4nEp8aiONqpe2jfxPz9Y/nakzyJ58QcksRA28g4PUKDLTMGlb8zdmgKBbfZ1ynif1DSWWyhpB4"
	plaintext, err := AESGCMDecode(key, text)
	t.Error(err, string(plaintext))
}

func TestAESGCMDecode2(t *testing.T) {
	key := "8bdd7c8fc6b8fcc83eae2c415845e0e60dbf3711b0f9517517bbdbf075d9884a"
	text := "/JwAOappsPlS4c+/8Q+GSkaZMtkqHbZKFp0DhUjt8bpZFIyq435RIyYnf2aomnR8A6Wx"
	plaintext, err := AESGCMDecode(key, text)
	t.Log(err, string(plaintext))
}

func TestAESGCMDecode3(t *testing.T) {
	key := "feb25db5dc122a255ea7042e38218c3ccd9569050ffb6e4c7dc780dfd2753123"
	text := "/JwAOappsPlS4c+/8Q+GSkaZMtkqHbZKFp0DhUjt8bpZFIyq435RIyYnf2aomnR8A6Wx6Wx"
	plaintext, err := AESGCMDecode(key, text)
	t.Log(err, string(plaintext))
}
