package user_agent

import "strings"

// Qq 环境判断
// qq小程序 https://q.qq.com/wiki/develop/miniprogram/component/open-ability/web-view.html miniProgram
// qq内置浏览器
// ---- 安卓 MQQBrowser QQ
// ---- ios 空格 QQ
// qq浏览器 MQQBrowser

const (
	QqBrowserId = "MQQBrowser"
	QqUniqueId  = "QQ"
	QqMiniId    = "miniProgram"
)

type Qq struct {
	version string
	mini    bool
	in      bool
}

func (p *UserAgent) detectQq(sections []section) {
	if strings.Contains(p.ua, QqBrowserId) || strings.Contains(p.ua, QqUniqueId) {
		p.qq.in = true
	}
	if p.qq.in {
		// 解析小程序标识
		sec, has := listSectionGetName(sections, QqMiniId)
		if has {
			p.qq.mini = true
			p.qq.version = sec.version
		}

	}
}
func (p *UserAgent) InQq() bool {
	return p.qq.in
}
func (p *UserAgent) InQqMini() bool {
	return p.qq.mini && p.InQq()
}
func (p *UserAgent) InQqBrowser() bool {
	return (p.InQqBrowserAndroid() || p.InQqBrowserIos()) && p.InQq()
}
func (p *UserAgent) InQqBrowserAndroid() bool {
	return !p.InQqMini() && strings.Contains(p.ua, QqBrowserId) && strings.Contains(p.ua, QqUniqueId) && p.InQq()
}
func (p *UserAgent) InQqBrowserIos() bool {
	return !p.InQqMini() && (p.platform == "iPhone" || p.platform == "iPad") && strings.Contains(p.ua, QqUniqueId) && p.InQq()
}

func (p *UserAgent) QqKey() string {
	var s = make([]string, 0)
	if p.InQq() {
		if p.InQqMini() {
			s = append(s, "qq_mini")
		} else if p.InQqBrowser() {
			s = append(s, "qq_browser")
			if p.InQqBrowserAndroid() {
				s = append(s, "qq_android")
			} else if p.InQqBrowserIos() {
				s = append(s, "qq_ios")
			}
		}
	}
	return strings.Join(s, "_")
}
