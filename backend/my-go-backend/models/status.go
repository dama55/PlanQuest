package models

// タスクのステータスを表す型
type Status int

const (
	NotStarted Status = iota     //0
	InProgress                   // 1
    Completed                    // 2
    OnHold                       // 3
)

// タスクステータスを文字列で表示するための関数
func (s Status) String() string {
    return [...]string{"NotStarted", "InProgress", "Completed", "OnHold"}[s]
}

// ステータスが有効かどうかを確認する関数
func IsState(val int) bool {
    if val >= 0 && val < 3{
        return true
    }else{
        return false
    }
}