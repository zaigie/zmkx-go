# zmkx-go

[zmkx.app](https://github.com/xingrz/zmkx.app) 的 Go 实现

## 使用

提供了一个简单的 CLI 操作及作为调用 demo

### 编译

```bash
rm -rf dist/ && mkdir -p dist/ && go build -o ./dist/zmkx-cli ./cmd/zmkx-cli/main.go
```

### 运行

```bash
# 获取版本信息
./zmkx-cli version
# 获取旋钮信息
./zmkx-cli knob
# 获取电机状态
./zmkx-cli motor
# 获取RGB状态
./zmkx-cli rgb
# 设置墨水屏
./zmkx-cli eink -f {*.jpg|png}
# 设置墨水屏翻转阈值（1-65535，默认为32768）
./zmkx-cli eink -f {*.jpg|png} -t 65535
```

## 开发

### 安装依赖

```bash
go mod download
```

### Windows 开发注意事项

本仓库使用的 go-hid 库虽然也是基于 libhidapi，但是在运行过程中需要 GCC 环境编译

在 Windows 上需要一个兼容的 GCC 工具链，使用 WSL 等方案都可行，最为快速简单的还是使用 MSYS2：

1. 到 [MSYS2 官网](http://www.msys2.org/) 下载安装包
2. 在 MSYS2 环境中安装 GCC 工具链，默认回车 all 全安装即可
   - `pacman -S --needed base-devel mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain`
3. 添加 mingw64 到环境变量
   - `echo 'export PATH=/mingw64/bin:$PATH' >> ~/.bashrc`
4. 将 Go 开发环境安装的路径添加到环境变量(这里只是举例)
   - `echo 'export PATH="/d/Program Files/Go/bin:$PATH"' >> ~/.bashrc`

## 相关链接

- [zmkx-sdk](https://github.com/xingrz/zmkx-sdk)
- [zmkx.app](https://github.com/xingrz/zmkx.app)
- [ZMK for HW-75](https://github.com/xingrz/zmk-config_helloword_hw-75)
- [peng-zhihui/HelloWord-Keyboard](https://github.com/peng-zhihui/HelloWord-Keyboard)

## 协议

[MIT License](LICENSE)
