package user_agent

import "testing"

func TestReferrer(t *testing.T) {
	// https://servicewechat.com/{appid}/{version}/page-frame.html

	t.Log(NewReferrer("https://servicewechat.com/wxd2833852a6136687/25/page-frame.html").GetWechatMiniInfo())
	t.Log(NewReferrer("https://wxd2833852a6136687.hybrid.alipay-eco.com/wxd2833852a6136687/25/index.html#page-frame").GetAliMiniInfo())
	t.Log(NewReferrer("https://smartapps.cn/wxd2833852a6136687/25/page-frame.html").GetBaiduMiniInfo())

}
