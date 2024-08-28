package encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/zeromicro/go-zero/core/codec"
)

// 密码加密种子
const (
	passwordEncryptSeed = "(beyond)@#$"
	// 手机AES加密密钥
	mobileAesKey = "5A2E746B08D846502F37A6E2D85D583B"
)

// 加密密码
func EncPassword(password string) string {
	// 将密码和种子拼接后进行MD5加密
	return Md5Sum([]byte(strings.TrimSpace(password + passwordEncryptSeed)))
}

// 加密手机号
func EncMobile(mobile string) (string, error) {
	// 使用AES加密手机号
	data, err := codec.EcbEncrypt([]byte(mobileAesKey), []byte(mobile))
	if err != nil {
		return "", err
	}

	// 将加密后的数据转换为Base64编码
	return base64.StdEncoding.EncodeToString(data), nil
}

// 解密手机号
func DecMobile(mobile string) (string, error) {
	// 将Base64编码的数据解码
	originalData, err := base64.StdEncoding.DecodeString(mobile)
	if err != nil {
		return "", err
	}
	// 使用AES解密手机号
	data, err := codec.EcbDecrypt([]byte(mobileAesKey), originalData)
	if err != nil {
		return "", err
	}

	// 返回解密后的手机号
	return string(data), nil
}

// 计算MD5值
func Md5Sum(data []byte) string {
	// 计算MD5值
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}

// 将16字节的数组转换为字节数组
func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	// 返回字节数组
	return tmp[16:]
}
