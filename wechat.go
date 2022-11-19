package user_agent

import "strings"

// Wechat 微信环境判断 MicroMessenger
// 企业微信 https://developer.work.weixin.qq.com/document/path/90315
// 微信小程序 https://developers.weixin.qq.com/miniprogram/dev/component/web-view.html
// 微信浏览器 https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/WiFi_Hardware_Authentication_Protocol_Interface_Description.html

const (
	WechatUniqueId     = "MicroMessenger"
	WechatBrowserId    = "micromessenger"
	WechatWorkUniqueId = "wxwork"
	WechatNetTypeId    = "NetType"
	WechatMiniId       = "miniProgram"
)

type Wechat struct {
	netType     string
	version     string
	workVersion string
	mini        bool
	browser     bool
	in          bool
}

func (p *UserAgent) detectWechat(sections []section) {
	if strings.Contains(p.ua, WechatUniqueId) || strings.Contains(p.ua, WechatBrowserId) {
		p.wechat.in = true
	}

	if p.wechat.in {
		sec, _, has, _ := arraySectionFilter(sections, WechatWorkUniqueId, false)
		if has {
			p.wechat.workVersion = sec.version
			if len(sec.comment) >= 2 {
				if sec.comment[0] == WechatUniqueId {
					p.wechat.version = sec.comment[1]
				}
			}
		}

		sec, has = listSectionGetName(sections, WechatUniqueId)
		if has {
			p.wechat.version = sec.version
		}
		// 解析 netType
		sec, has = listSectionGetName(sections, WechatNetTypeId)
		if has {
			p.wechat.netType = sec.version
		}

		// 解析小程序标识
		sec, has = listSectionGetName(sections, WechatMiniId)
		if has {
			p.wechat.mini = true
		}

	}
}

func (p *UserAgent) InWechat() bool {
	return p.wechat.in
}

func (p *UserAgent) InWechatWork() bool {
	return len(p.wechat.workVersion) > 0
}

func (p *UserAgent) InWechatAndroid() bool {
	return p.OSInfo().Name == "Android" && p.InWechat()
}

func (p *UserAgent) InWechatIphone() bool {
	return p.platform == "iPhone" && p.InWechat()
}

func (p *UserAgent) InWechatWindow() bool {
	return p.OSInfo().Name == "Windows" && p.InWechat()
}

func (p *UserAgent) InWechatMac() bool {
	return strings.Contains(p.OSInfo().Name, "Mac") && p.InWechat()
}

func (p *UserAgent) InWechatMini() bool {
	return p.wechat.mini
}

func (p *UserAgent) InWechatBrowser() bool {
	return !p.InWechatMini() && p.InWechat()
}

func (p *UserAgent) WechatKey() string {
	var s = make([]string, 0)
	if p.InWechat() {
		if p.InWechatWork() {
			s = append(s, "wechat_work")
		}

		if p.InWechatAndroid() {
			s = append(s, "wechat_android")
		} else if p.InWechatIphone() {
			s = append(s, "wechat_iphone")
		} else if p.InWechatWindow() {
			s = append(s, "wechat_window")
		} else if p.InWechatMac() {
			s = append(s, "wechat_mac")
		}

		if p.InWechatMini() {
			s = append(s, "wechat_mini")
		} else if p.InWechatBrowser() {
			s = append(s, "wechat_browser")
		}

	}
	return strings.Join(s, "_")
}
