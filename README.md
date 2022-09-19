# Web指纹识别工具

<a href="https://github.com/yhy0/FuckFingerprint"><img alt="Release" src="https://img.shields.io/badge/go-1.18-blue"></a>
<a href="https://github.com/yhy0/FuckFingerprint"><img alt="Release" src="https://img.shields.io/badge/release-v1.0.0-brightgreen"></a>
<a href="https://github.com/yhy0/FuckFingerprint"><img alt="Release" src="https://img.shields.io/badge/LICENSE-GPL-important"></a>


## 独乐乐不如众乐乐
Web指纹采取在线获取的方式，这样就可以不用更新程序即可使用最新的指纹，可以很方便的集成进扫描器中。

默认会使用在线的指纹，当在线获取失败时才会使用内置的指纹。

在线指纹依托于 raw.githubusercontent.com ,对于国内不友好可以通过指定代理获取(注:指定代理后，指纹识别也会走代理)



在线指纹地址可以在 `pkg/config/config.go`中修改为原项目中的地址
``` go
const EHoleFingerDataOnline = "https://raw.githubusercontent.com/veo/vscan/main/pkg/fingerprint/dicts/eHoleFinger.json"
const LocalFingerDataOnline = "https://raw.githubusercontent.com/veo/vscan/main/pkg/fingerprint/dicts/localFinger.json"
const AfrogFingerDataOnline = "https://raw.githubusercontent.com/zan8in/afrog/main/pkg/fingerprint/web_fingerprint_v3.json"
```
默认使用本项目地址
``` go
const EHoleFingerDataOnline = "https://raw.githubusercontent.com/yhy0/FuckFingerprint/main/fingerPrints/eHoleFinger.json"
const LocalFingerDataOnline = "https://raw.githubusercontent.com/yhy0/FuckFingerprint/main/fingerPrints/localFinger.json"
const AfrogFingerDataOnline = "https://raw.githubusercontent.com/yhy0/FuckFingerprint/main/fingerPrints/web_fingerprint_v3.json"
```


## 参考
指纹直接提取缝合了以下两个扫描器：

- [vscan](https://github.com/veo/vscan/)
- [afrog](https://github.com/zan8in/afrog)



感谢大佬无私奉献

## 📄 免责声明

本工具仅面向合法授权的企业安全建设行为，在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，作者将不承担任何法律及连带责任。

在使用本工具前，请您务必审慎阅读、充分理解各条款内容，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要使用本工具。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。

## 🌟 Star

[![Stargazers over time](https://starchart.cc/yhy0/FuckFingerprint.svg)](https://starchart.cc/yhy0/FuckFingerprint)
      
