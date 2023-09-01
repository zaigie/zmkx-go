# zmkx-go

[zmkx.app](https://github.com/xingrz/zmkx.app) 的 Go 实现

## 开发

```bash
go mod download
```

### Windows 开发注意事项

[github.com/sstallion/go-hid/issues/10](https://github.com/sstallion/go-hid/issues/10)

本仓库使用的 go-hid 库虽然也是基于 libhidapi，但是在运行过程中需要 GCC 环境编译使用 CGO
CGO 在 Windows 上可能会有些麻烦，但是只需要一个兼容的 GCC 工具链。可以参考以下的链接使用 MSYS2 最有帮助：https://github.com/faiface/pixel/wiki/Building-Pixel-on-Windows

## 相关链接

- [zmkx-app](https://github.com/xingrz/zmkx-sdk)
- [zmkx.app](https://github.com/xingrz/zmkx.app)
- [ZMK for HW-75](https://github.com/xingrz/zmk-config_helloword_hw-75)
- [peng-zhihui/HelloWord-Keyboard](https://github.com/peng-zhihui/HelloWord-Keyboard)

## 协议

[MIT License](LICENSE)
