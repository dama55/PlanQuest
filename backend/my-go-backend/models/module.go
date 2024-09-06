package models

import (
	"errors"
	"github.com/google/uuid"
)

// モジュールインターフェース

type ModuleInterface interface {
	SetName(newname string)
	SetDescription(newdesc string)
	SetPriority(newpriority int) error
	AddTask()
	FindTask(id uuid.UUID) (int, error)
	CheckStatus()
	RemoveTask(id uuid.UUID) error
}

// モジュール
type Module struct {
	ID uuid.UUID
	Name string
	Description string
	StatusVal Status
	PriorityVal Priority
	Tasks []TaskInterface
	PareProject *Project
}

// モジュール初期化
func InitModule() ModuleInterface {
	return &Module{
		ID: uuid.New(),
		Name: "モジュール名",
		Description: "モジュールの説明",
		StatusVal: NotStarted,
		PriorityVal: NewPriority(3),
	}
}

func (m *Module) SetName(newname string){
	m.Name = newname
}

func (m *Module) SetDescription(newdesc string){
	m.Description = newdesc
}

func (m *Module) SetPriority(newpriority int) error {
	// エラーを返す
	return m.PriorityVal.Set(newpriority)
}

func (m *Module) AddTask(){
	m.Tasks = append(m.Tasks, InitTask(m))
}

func (m *Module) FindTask(id uuid.UUID) (int, error){
	// 該当するタスクを検索
	for i, task := range m.Tasks {
		if task.GetId() == id {
			// タスクが見つかった場合返す
			return i, nil
		}
	}
	return -1, errors.New("タスクが見つかりません")
}
func (m *Module) CheckStatus(){
	allNotStarted := true
	allCompleted := true 
	for _, task := range m.Tasks{
		st := task.GetStatus()
		if st == InProgress || st == OnHold{
			m.StatusVal = InProgress
			return
		}else if st == NotStarted{
			// 全ての要素がCompletedという可能性は消えた
			allCompleted = false
		}else{
			// 残っているのはst == Completedの場合
			// 全ての要素がNotStartedという可能性は消えた
			allNotStarted = false
		}
	}
	if allNotStarted{
		//全てがNotStartedの場合
		m.StatusVal = NotStarted
	}else if allCompleted{
		//全てがCompletedの場合
		m.StatusVal = Completed
	}else{
		m.StatusVal = InProgress
	}
}

func (m *Module) RemoveTask(id uuid.UUID) error {
	idx, e := m.FindTask(id)
	if e != nil{
		return e
	}
	// 要素の削除
	m.Tasks = append(m.Tasks[:idx], m.Tasks[idx+1:]...)
	m.CheckStatus()
	return nil
}
