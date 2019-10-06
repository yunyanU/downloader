package downloader

import (
	"github.com/valyala/fasthttp"
	"os"
)

// 任务状态
const (
	wait = iota
	downloading
	paused
	errored
)

// 任务事件
type TaskEvent struct {
	Resume chan struct{}
	Pause  chan struct{}
	Cancel chan struct{}
}

type Task struct {
	id         int
	renewal    bool // 是否支持断点续传
	Status     int  //下载状态
	fileLength int
	Url        string
	finalLink  string
	file       *os.File
	FileName   string
	SavePath   string
	bts        map[int]*bt
	Event      *TaskEvent
	client     *fasthttp.HostClient
}

func (task *Task) Id() taskId {
	id := taskId(task.id)
	return id
}
