package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"runtime"
	"testing"
)

const publicKeyPrefix = "datahome.io"

func TestGenRSAKey(t *testing.T) {
	var prvKey []byte
	var pubKey []byte
	//////////生成私钥/////////
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err == nil {
		x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
		prvKey =
			pem.EncodeToMemory(&pem.Block{
				Type:  eccPrivateKeyPrefix,
				Bytes: x509PrivateKey,
			})
		fmt.Println(`生成的私钥`, string(prvKey))

	}
	////////////生成公钥/////////
	publicKey := privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err == nil {
		pubKey =
			pem.EncodeToMemory(&pem.Block{
				Type:  publicKeyPrefix,
				Bytes: x509PublicKey,
			})
		fmt.Println(`生成的公钥`, string(pubKey))
	}

	////////////公钥加密/////////
	block, _ := pem.Decode(pubKey)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println("error:", err)
	}
	pk := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pk, []byte("杭州点存区块链"))
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("密文", hex.EncodeToString(cipherText))
	////////////公钥解密/////////

	block, _ = pem.Decode(prvKey)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("error:", err)
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, prikey, cipherText)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("明文：", string(plainText))

}

func TestGenECCKey(t *testing.T) {
	var prvKey []byte
	var pubKey []byte
	//////////生成私钥/////////
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err == nil {
		prv, err := x509.MarshalECPrivateKey(privateKey)
		if err == nil {
			prvKey =
				pem.EncodeToMemory(&pem.Block{
					Type:  eccPrivateKeyPrefix,
					Bytes: prv,
				})
			fmt.Println(`生成的私钥`, string(prvKey))
		}
	}
	////////////生成公钥/////////
	publicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err == nil {
		pubKey =
			pem.EncodeToMemory(&pem.Block{
				Type:  publicKeyPrefix,
				Bytes: publicKey,
			})
		fmt.Println(`生成的公钥`, string(pubKey))
	}

	////////////公钥加密/////////
	block, _ := pem.Decode(pubKey)

	//防止用户传的密钥不正确导致panic,这里恢复程序并打印错误
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "请检查密钥是否正确")
			default:
				log.Println("error:", err)
			}
		}
	}()
	//2. block中的Bytes是x509编码的内容, x509解码
	tempPublicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//解码得到ecdsa包中的私钥
	//类型断言
	//转换为以太坊包中的ecies包中的私钥
	ukey := ImportECDSAPublic(tempPublicKey.(*ecdsa.PublicKey))

	crypttext, err := Encrypt(rand.Reader, ukey, []byte("杭州点存区块链"), nil, nil)
	fmt.Println("密文", hex.EncodeToString(crypttext))
	////////////公钥解密/////////

	block, _ = pem.Decode(prvKey)

	//防止用户传的密钥不正确导致panic,这里恢复程序并打印错误
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "请检查密钥是否正确")
			default:
				log.Println("error:", err)
			}
		}
	}()
	//2. block中的Bytes是x509编码的内容, x509解码
	tempPrivateKey, _ := x509.ParseECPrivateKey(block.Bytes)
	//解码得到ecdsa包中的私钥
	//转换为以太坊包中的ecies包中的私钥
	ikey := ImportECDSA(tempPrivateKey)

	//用私钥来解密密文
	plainText, err := ikey.Decrypt(crypttext, nil, nil)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("明文：", string(plainText))
}
