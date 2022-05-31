package netease

/**
 * 功能描述：
 *
 * Created with GoLand
 * @Author: david.tao
 * @CreateTime 2022-05-31 10:42
 *
 * 修改历史：(修改人，修改时间，修改原因/内容)
 *
 **/
type MyError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (this *MyError) Error() string {
	return this.Msg
}
