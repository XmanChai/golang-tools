package wechat

//MEMO 建造者模式
type BuildProcess interface {
	SetConfigs() BuildProcess //获取拼装参数
	SetRequest() BuildProcess //发出请求
	SetResult() BuildProcess  //处理返回结果
	GetResult() interface{}   //返回结果
}

//生产指导者
type BuildDirector struct {
	builder BuildProcess
}

func (b *BuildDirector) Constructor() {
	b.builder.SetConfigs().SetRequest().SetResult()
}

func (b *BuildDirector) SetBuilder(bp BuildProcess) {
	b.builder = bp
}
