package request

// EncryptRequest 数据加密请求对象
type EncryptRequest struct {
	AppKey    string `json:"app_key"`
	Plaintext string `json:"plaintext"`
}

// DecryptRequest 数据解密请求对象
type DecryptRequest struct {
	AppKey     string `json:"app_key"`
	Ciphertext string `json:"ciphertext"`
}
