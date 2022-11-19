package user_agent

import "testing"

func newTestPointBool(input bool) *bool {
	return &input
}

func TestWechat(t *testing.T) {

	var testWechat = []struct {
		name    string
		ua      string
		uaList  []string
		in      *bool
		inWork  *bool
		iphone  *bool
		window  *bool
		android *bool
		mac     *bool
		mini    *bool
		browser *bool
	}{
		{
			name:   "企业微信 iphone",
			ua:     "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89 wxwork/2.2.0 MicroMessenger/6.3.2",
			inWork: newTestPointBool(true),
			iphone: newTestPointBool(true),
		},
		{
			name:   "企业微信 windows",
			ua:     "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 wxwork/2.1.3 (MicroMessenger/6.2) WindowsWechat QBCore/3.43.644.400 QQBrowser/9.0.2524.400",
			inWork: newTestPointBool(true),
			window: newTestPointBool(true),
		},
		{
			name:    "企业微信 android",
			ua:      "Mozilla/5.0 (Linux; Android 7.1.2; g3ds Build/NJH47F; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043508 Safari/537.36 wxwork/2.2.0 MicroMessenger/6.3.22 NetType/WIFI Language/zh",
			inWork:  newTestPointBool(true),
			android: newTestPointBool(true),
		},
		{
			name:   "企业微信 Mac",
			ua:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) wxwork/2.2.0 (MicroMessenger/6.2) WeChat/2.0.4",
			inWork: newTestPointBool(true),
			mac:    newTestPointBool(true),
		},
		{
			name:    "微信小程序",
			ua:      "Mozilla/5.0 (Linux; Android 7.1.1; MI 6 Build/NMF26X; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.132 MQQBrowser/6.2 TBS/043807 Mobile Safari/537.36 MicroMessenger/6.6.1.1220(0x26060135) NetType/4G Language/zh_CN MicroMessenger/6.6.1.1220(0x26060135) NetType/4G Language/zh_CN miniProgram",
			mini:    newTestPointBool(true),
			android: newTestPointBool(true),
		},
		{
			name: "微信浏览器",
			ua:   "Mozilla/5.0 (Linux; U; Android 4.1.2; zh-cn; GT-I9300 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 MicroMessenger/5.2.380",
			uaList: []string{
				"Mozilla/5.0 (Linux; U; Android 5.0.2; zh-cn; NX511J Build/LRX22G) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/8.8 TBS/88888888 Mobile Safari/533.1 MicroMessenger/6.3.8.56_re6b2553.680 NetType/ctlte Language/zh_CN MicroMessenger/6.3.8.56_re6b2553.680 NetType/ctlte Language/zh_CN",
				"Mozilla/5.0 (Linux; Android 11; SM-N9860 Build/RP1A.200720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3171 MMWEBSDK/20211202 Mobile Safari/537.36 MMWEBID/8157 MicroMessenger/8.0.18.2060(0x28001237) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64",
				"Mozilla/5.0 (Linux; Android 10; EVR-AN00 Build/HUAWEIEVR-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3179 MMWEBSDK/20211001 Mobile Safari/537.36 MMWEBID/2013 MicroMessenger/8.0.16.2040(0x2800105F) Process/toolsmp WeChat/arm64 Weixin NetType/5G Language/zh_CN ABI/arm64",
				"Mozilla/5.0 (Linux; Android 11; POCO F2 Pro Build/RKQ1.200826.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36 MMWEBID/1230 MicroMessenger/8.0.17.2040(0x28001133) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64",
				"Mozilla/5.0 (Linux; Android 10; ASUS_I001DA Build/QKQ1.190825.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.72 MQQBrowser/6.2 TBS/045908 Mobile Safari/537.36 MMWEBID/3216 MicroMessenger/8.0.16.2040(0x28001056) Process/tools WeChat/arm64 Weixin NetType/4G Language/zh_CN ABI/arm64",
				"Mozilla/5.0 (Linux; Android 10; Pixel 4 XL Build/QQ2A.200305.004.A1; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3140 MMWEBSDK/20211001 Mobile Safari/537.36 MMWEBID/8391 MicroMessenger/8.0.16.2040(0x2800103A) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64",
				"Mozilla/5.0 (iPhone; CPU iPhone OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.17(0x18001122) NetType/4G Language/zh_CN",
				"Mozilla/5.0 (Linux; Android 11; Redmi K30 Pro Build/RKQ1.200826.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3179 MMWEBSDK/20211202 Mobile Safari/537.36 MMWEBID/7002 MicroMessenger/8.0.18.2060(0x28001237) Process/toolsmp WeChat/arm64 Weixin NetType/5G Language/zh_CN ABI/arm64",
			},
			browser: newTestPointBool(true),
		},
	}

	for _, s := range testWechat {

		testList := s.uaList
		if len(s.ua) > 0 {
			testList = append(testList, s.ua)
		}
		for index, c := range testList {
			inst := New(c)
			if s.in != nil {
				if inst.InWechat() != *s.in {
					t.Errorf("测试 %s : %d - in  - 需求 %t 实际%t", s.name, index, *s.in, inst.InWechat())
				}
			}
			if s.inWork != nil {
				if inst.InWechatWork() != *s.inWork {
					t.Errorf("测试 %s : %d - inwork  - 需求 %t 实际%t", s.name, index, *s.inWork, inst.InWechatWork())
				}
			}
			if s.iphone != nil {
				if inst.InWechatIphone() != *s.iphone {
					t.Errorf("测试 %s : %d - inwork  - 需求 %t 实际%t", s.name, index, *s.iphone, inst.InWechatIphone())
				}
			}

			if s.window != nil {
				if inst.InWechatWindow() != *s.window {
					t.Errorf("测试 %s : %d - inwindow  - 需求 %t 实际%t", s.name, index, *s.window, inst.InWechatWindow())
				}
			}
			if s.android != nil {
				if inst.InWechatAndroid() != *s.android {
					t.Errorf("测试 %s : %d - inandroid  - 需求 %t 实际%t", s.name, index, *s.android, inst.InWechatAndroid())
				}
			}
			if s.mac != nil {
				if inst.InWechatMac() != *s.mac {
					t.Errorf("测试 %s : %d - inmac  - 需求 %t 实际%t", s.name, index, *s.mac, inst.InWechatMac())
				}
			}
			if s.mini != nil {
				if inst.InWechatMini() != *s.mini {
					t.Errorf("测试 %s : %d - inmini  - 需求 %t 实际%t", s.name, index, *s.mini, inst.InWechatMini())
				}
			}
			if s.browser != nil {
				if inst.InWechatBrowser() != *s.browser {
					t.Errorf("测试 %s : %d - inbrowser  - 需求 %t 实际%t", s.name, index, *s.browser, inst.InWechatBrowser())
				}
			}

		}

	}

	// 微信浏览器
}
