package user_agent

import "strings"

type Env struct {
	ua       string
	referrer string
	uaInst   *UserAgent
	refInst  *Referrer
}

func NewEnv(ua, referrer string) *Env {
	inst := &Env{
		ua:       ua,
		referrer: referrer,
	}
	inst.uaInst = New(ua)
	inst.refInst = NewReferrer(referrer)
	return inst
}

func GetEnvKey(ua, referrer string) string {
	if len(ua) < 1 && len(referrer) < 1 {
		return ""
	}
	return NewEnv(ua, referrer).GetKey()
}

func (c *Env) GetKey() string {

	var s = make([]string, 0)
	// 对于小程序环境来说 通过referrer 判断相对比较准确
	if c.refInst.IsMini() {
		k, appid := c.refInst.MiniKey()
		s = append(s, k)
		s = append(s, appid)
	} else {
		// 若不是小程序层面则通过 ua 来切分
		if c.uaInst.InWechat() {
			s = append(s, c.uaInst.WechatKey())
		} else if c.uaInst.InQq() {
			s = append(s, c.uaInst.QqKey())
		} else {
			// 先排除bot
			if c.uaInst.Bot() {
				name, version := c.uaInst.Browser()
				s = append(s, "bot")
				s = append(s, name)
				s = append(s, version)
			} else {
				// 不是在常用app里的话则获取运行环境
				s = append(s, c.uaInst.os)
				name, version := c.uaInst.Engine()
				s = append(s, name)    // AppleWebKit
				s = append(s, version) // 533.1
			}
		}
	}

	// 去掉为空的值
	var n = make([]string, 0, len(s))
	for _, ss := range s {
		if len(ss) > 0 {
			n = append(n, ss)
		}
	}

	return strings.Join(n, "_")
}
