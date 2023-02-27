<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [命令行使用简单教程](#%E5%91%BD%E4%BB%A4%E8%A1%8C%E4%BD%BF%E7%94%A8%E7%AE%80%E5%8D%95%E6%95%99%E7%A8%8B)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### 命令行使用简单教程

1、登录

通过 localStorage 获取到对应用户凭证

填写相关配置到 yaml 文件中

```yaml
token: dc98893ca481dd39019abaf80c341872
sign: 33f09e0f11d96cd83821a9945f3da0de
st_flpv: vkx9gecd6JcZH2F9k665 # 注意这里是最外层的 st_flpv 而不是 userInfo.data.st_flpv
cityId: 20
userId: 4612349 # userInfo.data.userId
```

2、拿到对应场次 id

在秀动官网通过后缀 url 拿到 activityId

3、列出场次票种列表

```shell
showstart.exe tickets -a 173474 --config cli-sample.yaml
```

```console
演出名称: xx站
场次 ID: 173474
┌──────────────────────────────────┬───────────┬────────┐
│ 票种标识                         │ 票种      │ 售价   │
├──────────────────────────────────┼───────────┼────────┤
│ 061f79744fcdf29a489b88c87ed912dc │ 早鸟票    │ 160    │
│ 282b6d4fbc24baba6b84471a4eec864f │ 预售票    │ 200    │
│ 1214c6750d03dcdcc5e3979f58c82411 │ 全价票    │ 260    │
└──────────────────────────────────┴───────────┴────────┘
```

4、获取身份 id

> 有一些场次需要指定身份 id 实名

```shell
showstart.exe idCard --config cli-sample.yaml
```

```shell
┌─────────────────┬───────────┬───────────┬────────────────────┬────────┐
│ 观演人标识      │ 姓名      │ 类型      │ 号码               │ isSelf │
├─────────────────┼───────────┼───────────┼────────────────────┼────────┤
│ 538xx81         │ 郑xx      │ 身份证    │ 440************1XX │ 1      │
│ 538xx62         │ 廖xx      │ 身份证    │ 440************2XX │ 0      │
└─────────────────┴───────────┴───────────┴────────────────────┴────────┘
```
