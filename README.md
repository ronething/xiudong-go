<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [xiudong-go](#xiudong-go)
    - [usage](#usage)
    - [cli 效果图](#cli-%E6%95%88%E6%9E%9C%E5%9B%BE)
    - [免责声明](#%E5%85%8D%E8%B4%A3%E5%A3%B0%E6%98%8E)
    - [acknowledgement](#acknowledgement)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### xiudong-go

Web 版秀动请求样例学习

github repo: https://github.com/ronething/xiudong-go

#### usage

- 下载对应 [releases](https://github.com/ronething/xiudong-go/releases) 即可

- 如果你需要手动编译 可以考虑执行以下命令

```shell
git clone https://github.com/ronething/xiudong-go.git && cd xiudong-go/cli
make build
```

具体使用可参考 [./cli](./cli) 和 [./docs](./docs)

notice: 此程序更多是学习用途，可以简单分析下自己写 order 命令，祝大家好运~

```
$ ./showstart         
showstart cli sample

Usage:
  showstart [command]

Available Commands:
  address     查询个人地址
  help        Help about any command
  idCard      查询已绑定观演人 id
  tickets     列出指定场次 ticketId 列表
  version     check version

Flags:
      --config string   config file (default is $HOME/.showstart.yaml)
  -h, --help            help for showstart

Use "showstart [command] --help" for more information about a command.

```

#### 免责声明

见 [Disclaimer](./Disclaimer.md)

#### acknowledgement

- wap.showstart.com