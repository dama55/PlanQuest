package models

import (
	"time"
	"github.com/google/uuid"
)

//タスクインターフェース
type TaskInterface interface {
	GetId() uuid.UUID
	GetStatus() Status
	SetName(newname string)
	SetPlDate(newplandate time.Time)
	ExecuteInterval()
	HoldInterval()
	RemoveInterval()
}

//タスク
type Task struct {
	ID uuid.UUID
	Name string
	StatusVal Status
	PareModule ModuleInterface
	PlannedDate time.Time
	TimeIntervals []*Timeset
	
}

func InitTask(paremodule *Module) TaskInterface {
	// 実施予定日，実施履歴などはまだ設定しない
	return &Task{
		ID: uuid.New(),
		Name: "タスク名",
		StatusVal: NotStarted,
		PareModule: paremodule,
	}
}

func (task *Task) SetName(newname string){
	task.Name = newname
}

func (task *Task) SetPlDate(newplandate time.Time){
	task.PlannedDate = newplandate
}

func (task Task) GetId() uuid.UUID {
	return task.ID
}

func (task Task) GetStatus() Status {
	return task.StatusVal
}

func (task *Task) ExecuteInterval(){
	if task.StatusVal == NotStarted || task.StatusVal == OnHold || task.StatusVal == Completed{
		task.TimeIntervals = append(task.TimeIntervals, InitTimeset())
		task.StatusVal = InProgress
	}else if task.StatusVal == InProgress{
		task.TimeIntervals[len(task.TimeIntervals) - 1].SetEndTime()
		task.StatusVal = Completed
	}
	//タスクの状態が変化したので親モジュールに伝える
	task.PareModule.CheckStatus()
}

// タスクを一時停止
func (task *Task) HoldInterval(){
	if task.StatusVal == InProgress{
		task.TimeIntervals[len(task.TimeIntervals) - 1].SetEndTime()
		task.StatusVal = OnHold
	}
	//タスクの状態が変化したので親モジュールに伝える
	task.PareModule.CheckStatus()
}

// タスクを削除
func (task *Task) RemoveInterval(){
	// インターバルの個数を保持
	l := len(task.TimeIntervals)
	// 要素が何もない場合は何もしない
	if l == 0{
		return
	}
	
	if task.StatusVal == NotStarted || task.StatusVal == OnHold || task.StatusVal == Completed{
		// タスクが開始前の時
		task.TimeIntervals = task.TimeIntervals[:l - 1]
	}else if task.StatusVal == InProgress {
		// タスクが開始前の時
		task.TimeIntervals = task.TimeIntervals[:l - 1]
		if l - 1 == 0{
			// インターバルがなくなった場合
			task.StatusVal = NotStarted
		}else{
			// インターバルがなくなっていない場合
			task.StatusVal = OnHold
		}
	}
	//タスクの状態が変化したので親モジュールに伝える
	task.PareModule.CheckStatus()
}





