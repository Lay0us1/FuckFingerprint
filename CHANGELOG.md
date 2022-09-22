# Changelog

## v1.1.0
- [feat] title 获取重构
- [feat] 添加简单蜜罐识别
- [fix] 修复默认`User-agent` 不规范的错误

## v1.0.2
- [feat] 日志默认改到当前路径下的 logs 文件夹下
- [feat] 在线获取默认改为 false ,可通过-g 指定为在线获取 eg: ./FuckFingerprint -url http://127.0.0.1 -g
- [feat] 默认不会进行自动跳转，当 vscan 的指纹识别不到时，增加一次自动跳转的请求

## v1.0.1
- [fix] 修复 http(s) 协议拼接错误

## v1.0.0
- 发布