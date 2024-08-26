package models

// タスクのステータスを表す型
type Status int

const {
	StatusNotStarted Status = iota     //0
	StatusInProgress                   // 1
    StatusCompleted                    // 2
    StatusOnHold                       // 3
}

// タスクステータスを文字列で表示するための関数
func (s TaskStatus) String() string {
    return [...]string{"NotStarted", "InProgress", "Completed", "OnHold"}[s]
}