<p><a href="https://go-kratos.dev/" target="_blank"><img src="https://github.com/go-kratos/kratos/blob/main/docs/images/kratos-large.png?raw=true" alt="Kratos"></a></p>


Translations: [简体中文](README.md)

# Kratos-layout

**Kratos-layout** 是基于 [Kratos](https://github.com/go-kratos/kratos)
以及模板项目 [Kartos-layout](https://github.com/go-kratos/kratos-layout) 结合公司实际应用场景封装的 Golang 微服务模板项目

> 名字来源于:《战神》游戏以希腊神话为背景，讲述奎托斯（Kratos）由凡人成为战神并展开弑神屠杀的冒险经历。

## 目标

致力于提供健全的微服务研发体验，整合相关框架及工具，降低微服务学习成本，提高研发效率

## 快速开始

### 基础环境

| 类型          | 依赖                                                              |
|-------------|-----------------------------------------------------------------|
| 开发语言        | [Go1.22](https://golang.org/dl/)                                |
| Protobuf    | [Protoc](https://github.com/protocolbuffers/protobuf)           |
| Protobuf 生成 | [Protoc-gen-go](https://github.com/protocolbuffers/protobuf-go) |
| 基础框架        | [Kratos](https://go-kratos.dev/docs/getting-started/start)      |
| 代码规范检测      | [GolangCi-lint](https://golangci-lint.run/welcome/install/)     |

### 基础环境安装

#### Golang:

```shell
# 下载安装包，根据提示安装
# Windows
https://go.dev/dl/go1.22.2.windows-amd64.msi
# Mac OS
https://go.dev/dl/go1.22.2.darwin-arm64.pkg
# Linux
https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
# Source
https://go.dev/dl/go1.22.2.src.tar.gz

# 安装完成后添加环境变量，下面演示 Mac&Linux 环境下增加 Golang 环境变量方法
# Windows 环境请自行 Google，本文不再概述
vim ~/.zshrc # Mac OS
vim ~/.bashrc # Linux

# vim 在输入法为英文的状态下，按i进入编辑模式，将下边内容添加到文件中
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

# 然后 按 : 输入wq ，保存退出。 重载一下 环境变量文件
source ~/.zshrc # Mac OS 
source ~/.bashrc # Linux 
```
#### Protobuf:

##### Protoc 安装
```shell
# Mac OS
brew install protobuf
# Windows and Source 打开以下链接下载对应系统安装包安装即可
https://github.com/protocolbuffers/protobuf/releases
```
#### Protoc-gen-go 安装
```shell
go install github.com/golang/protobuf/protoc-gen-go@latest
```

#### Kratos:
```shell
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade
```

#### GolangCI-lint:
```shell
# MacOS 
brew install golangci-lint
brew upgrade golangci-lint
# windows
choco install golangci-lint
# 源码安装
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2
```

### 创建模板项目
```
# 直接使用该模板源创建项目
kratos new helloworld -r https://git.corp.doulaoban.com/backend/kratos-layout.git

# 也可以通过环境变量指定源后创建项目
KRATOS_LAYOUT_REPO=https://git.corp.doulaoban.com/backend/kratos-layout.git
kratos new helloworld
```
## 代码提交规范

提交信息的结构应该如下所示:

```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

提交信息应按照下面的格式:

- fix: 简单描述已解决的问题
- feat(log): 新功能的简单描述
- deps(examples): 简单描述依赖关系的变化
- break(http): 简单描述重大变更的原因

## License

Kratos is MIT licensed. See the [LICENSE](./LICENSE) file for details.