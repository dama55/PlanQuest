package models

import (
	"time"
    "github.com/google/uuid"
)

type Timeset struct {
	ID uuid.UUID
	StartTime time.Time
	EndTime time.Time
	HasEndTime bool // EndTimeが
}

/*コンストラクタ*/
func InitTimeset() *Timeset {
	return &Timeset{
		ID: uuid.New(),
		StartTime: time.Now(), //開始時間の設定
		HasEndTime: false,
	}
}

/*終了時刻を設定するメソッド*/
func (t *Timeset) SetEndTime(){
	t.EndTime = time.Now() //終了時間の設定
	t.HasEndTime = true // EndTimeが設定されたことを記録
}

/*終了時刻が設定されたか判定*/
func (t *Timeset) IsFinished() bool {
	return t.HasEndTime
}

/* 開始時刻と終了時刻の組を返す */
func (t *Timeset) GetTimes() (time.Time, time.Time, bool){
	return t.StartTime, t.EndTime, t.HasEndTime
}