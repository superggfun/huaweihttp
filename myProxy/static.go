package myProxy

const (
	html = `<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0"/>
	<meta name="robots" content="noarchive">
	<title>华为云函数-IP代理</title>
	<link rel="shortcut icon" href="http://p0.meituan.net/csc/90273a80c903e1482783451820ab928a6190.png" type="image/x-icon">
	<style>
	h1{color:#FF00FF; text-align:center; font-family: Microsoft Jhenghei;}p{color:#FF00FF; font-size: 1.2rem;text-align:center;font-family: Microsoft Jhenghei;}
	#master {text-align: center;width: 100%;}
	</style>
	</head>
	<body style="background: #FFFFFF url(https://tenapi.cn/acg) no-repeat fixed center;">
	<table width="100%" height="100%" align="center"><tbody><tr>
	<td align="center">
	<h1><b>华为云函数-IP代理<br><h1>您好像没有输入链接地址哦，请在链接后面加上?url=https://xxx</h1><p>支持GET、POST、Headers、no-referers<br>不以盈利为目的使用本接口造成的任何后果概不负责</p></b></h1>
	</td></tr>
	<td id="master" >
		<p><font size="2"><a href="https://github.com/superggfun/huaweihttp">华为云函数-IP代理开源地址</a><br>内容仅用于学习研究，禁止用于任何商业或非法用途，违反者作者概不负责。</font></p>
	</td>
	</tbody>
	</table>

	</body>
	</html>`
)

func (m *APIG) noUrl() (string, map[string]string, int, error) {
	h := make(map[string]string, 1)
	h["Content-Type"] = "text/html; charset=UTF-8"
	return html, h, 201, nil
}
