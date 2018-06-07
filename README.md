微信小程序SDK for Go
===============================

第一个可用版本，随心使用，注意安全 😄

## 目前功能

- 数据解密

    - 解密 wx.getUserInfo
    - 解密 wx.getShareInfo
    - 解密 wx.getWeRunData

- 接口调用

    - todo

## 使用方法

- 数据解密

```golang

import (
    "fmt"
    "github.com/chekun/go-wechat-app/wechat/decrypt"
)

func main() {
    decryptor, err := New("your-app-id", "your-session-key")
    if err != nil {
        fmt.Println(err)//sessionKey不正确
    }

    if profile := decryptor.Profile( "your-encrypted-data", "your-iv"); decryptor.Err != nil {
        fmr.Println("Profile 解密失败 - ", decryptor.Err)
    } else {
        fmt.Println(profile)
    }
	

    if share := decryptor.Share( "your-encrypted-data", "your-iv"); decryptor.Err != nil {
        fmr.Println("Share Ticket 解密失败 - ", decryptor.Err)
    } else {
        fmt.Println(share)
    }

    if run := decryptor.Run( "your-encrypted-data", "your-iv"); decryptor.Err != nil {
        fmr.Println("Run Data 解密失败 - ", decryptor.Err)
    } else {
        fmt.Println(run)
    }
}

```