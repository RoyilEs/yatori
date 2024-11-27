package examples

import (
	"fmt"
	"testing"
	"yatori-go-core/global"
	"yatori-go-core/utils"
)

// 测试加密函数
func TestEncrypted(t *testing.T) {
	utils.YatoriCoreInit()
	//测试账号
	setup()
	users := global.Config.Users[4]
	// 调用函数进行加密
	accEncrypted := utils.CqieEncrypt(users.Account)
	passEncrypted := utils.CqieEncrypt(users.Password)
	// 输出加密后的数据
	fmt.Printf("Encrypted data: %x\n", accEncrypted)
	fmt.Printf("Encrypted data: %x\n", passEncrypted)
}
