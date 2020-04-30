package pattern

import "fmt"

//MEMO 不变化的是这个播放器接口
type Player interface {
	PlayMusic()
}

//MEMO 实现方法规则方
type MusicBox struct {
	src string
}

//MEMO 被适配器调用的其他方法，同结构
//TODO 异构的情况是否匹配 异构方法完全可以，参数也可以不一样
type SoundBox struct {
	src    string
	author string
}

type Adapter struct {
	//sounder SoundBox  完整的书写方法
	*SoundBox //只放入指针
}

func (m MusicBox) PlayMusic() {
	fmt.Printf("I am musicbox method playmusic %s\n", m.src)
}

func (s SoundBox) PlaySound() {
	fmt.Printf("I am soundbox method playsound %s\t author:%s\n", s.src, s.author)
}

//MEMO 适配器就是实现了接口的固定方法，只是在结构中放入了需要适配的结构指针
func (a Adapter) PlayMusic() {
	//a.sounder.PlaySound()
	a.PlaySound()
}

func Do() {
	var musicPlayer Player
	musicPlayer = MusicBox{src: "music.mp3"}
	musicPlayer.PlayMusic()
	var soundPlayer Player
	soundPlayer = Adapter{&SoundBox{
		src:    "mid",
		author: "xman",
	}}
	soundPlayer.PlayMusic()
}
