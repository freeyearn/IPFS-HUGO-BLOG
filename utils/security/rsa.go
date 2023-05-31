package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strconv"
	"time"
)

type RSA struct {
	PublicKeyPath  string
	PrivateKeyPath string
}

var RSAHelper *RSA

// GenerateRSAKey 生成RSA私钥和公钥，保存到文件中
func (r RSA) GenerateRSAKey(bits int) error {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create(r.PrivateKeyPath)
	if err != nil {
		return err
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	err = pem.Encode(privateFile, &privateBlock)
	if err != nil {
		return err
	}

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(r.PublicKeyPath)
	if err != nil {
		return err
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	err = pem.Encode(publicFile, &publicBlock)
	return nil
}

// Encrypt RSA加密
func (r RSA) Encrypt(plainText []byte) ([]byte, error) {
	//打开文件
	file, err := os.Open(r.PublicKeyPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	//返回密文
	return cipherText, nil
}

// Decrypt RSA解密
func (r RSA) Decrypt(cipherText []byte) ([]byte, error) {
	//打开文件
	file, err := os.Open(r.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//获取文件内容
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	//buf := privateKey
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	return plainText, err
}

// TimeStampDecrypt 含有时间戳的RSA解密
func (r RSA) TimeStampDecrypt(cipherText []byte, expireSecond int) ([]byte, error) {
	//打开文件
	file, err := os.Open(r.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//获取文件内容
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, info.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	//buf := privateKey
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	text := string(plainText)

	signedTime, err := strconv.Atoi(text[len(text)-12:])

	if err != nil {
		return nil, err
	}

	current, err := strconv.Atoi(time.Now().Format("200601021504"))
	if err != nil {
		return nil, err
	}

	// 时间戳验证
	if (current-signedTime)*60 > expireSecond {
		return nil, errors.New("RSA Expired time out")
	}

	return []byte(text[:len(text)-12]), err
}

func (r RSA) TimeStampEncrypt(cipherText string) ([]byte, error) {
	cipherText = cipherText + time.Now().Format("200601021504")
	encrypt, err := r.Encrypt([]byte(cipherText))
	if err != nil {
		return nil, err
	}
	return encrypt, nil
}

func InitRSAHelper(publicKeyPath string, privateKeyPath string) {
	RSAHelper = &RSA{PublicKeyPath: publicKeyPath, PrivateKeyPath: privateKeyPath}
}
