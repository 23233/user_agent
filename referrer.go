package user_agent

import (
	"net/url"
	"strings"
)

type Referrer struct {
	raw string
	u   *url.URL
}

func NewReferrer(ref string) *Referrer {
	u, _ := url.Parse(ref)

	return &Referrer{
		raw: ref,
		u:   u,
	}
}

type ReferrerMiniInfo struct {
	Appid   string `json:"appid,omitempty"`
	Version string `json:"version,omitempty"`
}

func (r *Referrer) pathInfoParse() *ReferrerMiniInfo {
	// 解析
	pathList := strings.Split(r.u.Path, "/")
	if len(pathList) < 4 {
		return nil
	}
	var info = new(ReferrerMiniInfo)
	info.Appid = pathList[1]
	info.Version = pathList[2]

	return info
}

func (r *Referrer) queryInfoParse() *ReferrerMiniInfo {
	appId := r.u.Query().Get("appid")
	version := r.u.Query().Get("version")
	if len(appId) < 1 || len(version) < 1 {
		return nil
	}
	var info = new(ReferrerMiniInfo)
	info.Appid = appId
	info.Version = version
	return info
}

func (r *Referrer) IsMini() bool {
	return r.IsWechatMini() || r.IsQqMini() || r.IsBytedanceMini() || r.IsAliMini() || r.IsBaiduMini()
}

func (r *Referrer) MiniKey() (string, string) {
	if r.IsMini() {
		if r.IsWechatMini() {
			return "wechat_mini", r.GetWechatMiniInfo().Appid
		} else if r.IsQqMini() {
			return "qq_mini", r.GetQqMiniInfo().Appid
		} else if r.IsBytedanceMini() {
			return "bytedance_mini", r.GetBytedanceMiniInfo().Appid
		} else if r.IsAliMini() {
			return "ali_mini", r.GetAliMiniInfo().Appid
		} else if r.IsBaiduMini() {
			return "baidu_mini", r.GetBaiduMiniInfo().Appid
		}
	}
	return "", ""
}

// IsWechatMini 是否是微信小程序请求
// https://developers.weixin.qq.com/miniprogram/dev/framework/ability/network.html
// https://servicewechat.com/{appid}/{version}/page-frame.html
func (r *Referrer) IsWechatMini() bool {
	return r.u.Hostname() == "servicewechat.com"
}

func (r *Referrer) GetWechatMiniInfo() *ReferrerMiniInfo {
	if !r.IsWechatMini() {
		return nil
	}
	return r.pathInfoParse()
}

// IsQqMini 是否为qq小程序
// https://q.qq.com/wiki/develop/miniprogram/frame/basic_ability/basic_network.html
// https://appservice.qq.com/{appid}/{version}/page-frame.html
func (r *Referrer) IsQqMini() bool {
	return r.u.Hostname() == "appservice.qq.com"
}

func (r *Referrer) GetQqMiniInfo() *ReferrerMiniInfo {
	if !r.IsQqMini() {
		return nil
	}
	return r.pathInfoParse()
}

// IsBytedanceMini 是否为字节跳动小程序
// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/guide/basic-ability/network/
// https://tmaservice.developer.toutiao.com/?appid={appid}&version={version}
func (r *Referrer) IsBytedanceMini() bool {
	return r.u.Hostname() == "tmaservice.developer.toutiao.com"
}

func (r *Referrer) GetBytedanceMiniInfo() *ReferrerMiniInfo {
	if !r.IsBytedanceMini() {
		return nil
	}
	return r.queryInfoParse()
}

// IsAliMini 是否为阿里小程序
// https://opendocs.alipay.com/mini/api/owycmh
// https://{appid}.hybrid.alipay-eco.com/{appid}/{version}/index.html#{page}
func (r *Referrer) IsAliMini() bool {
	return strings.HasSuffix(r.u.Hostname(), "hybrid.alipay-eco.com")
}

func (r *Referrer) GetAliMiniInfo() *ReferrerMiniInfo {
	if !r.IsAliMini() {
		return nil
	}
	return r.pathInfoParse()
}

// IsBaiduMini 是否为百度小程序
// https://smartprogram.baidu.com/docs/develop/api/net/net_rule/
// https://{域名}/{appKey}/{version}/page-frame.html
// https://smartapp.baidu.com/{appKey}/{version}/page-frame.html
// https://smartapps.cn/{appKey}/{version}/page-frame.html
// 自基础库版本 V3.170.0 起，其中域名由原来的 https://smartapp.baidu.com 更改为 https://smartapps.cn
func (r *Referrer) IsBaiduMini() bool {
	return r.u.Hostname() == "smartapps.cn" || r.u.Hostname() == "smartapp.baidu.com"
}

func (r *Referrer) GetBaiduMiniInfo() *ReferrerMiniInfo {
	if !r.IsBaiduMini() {
		return nil
	}
	return r.pathInfoParse()
}
