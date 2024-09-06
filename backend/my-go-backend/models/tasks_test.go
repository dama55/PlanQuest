package models

import (
	"testing"
	"time"
)

// タスクの初期化と状態変更のテスト
func TestTaskLifecycle(t *testing.T) {
	// モックモジュールの初期化（仮）
	module := InitModule()  // ここで ModuleInterface 型を返す
	mod, ok := module.(*Module)  // 必要ならばポインタ型にキャスト
	if !ok {
		t.Fatalf("Expected *Module, got %T", module)
	}

	// タスクの初期化
	taskInterface := InitTask(mod)
	task, ok := taskInterface.(*Task)
	if !ok {
		t.Fatalf("Expected *Task, got %T", taskInterface)
	}

	// 名前の設定
	task.SetName("新しいタスク")
	if task.Name != "新しいタスク" {
		t.Errorf("Expected task name to be 新しいタスク, but got %s", task.Name)
	}

	// プラン日付の設定
	plannedDate := time.Now().Add(24 * time.Hour)
	task.SetPlDate(plannedDate)
	if !task.PlannedDate.Equal(plannedDate) {
		t.Errorf("Expected planned date to be %v, but got %v", plannedDate, task.PlannedDate)
	}

	// インターバルの実行
	task.ExecuteInterval()
	if task.GetStatus() != InProgress {
		t.Errorf("Expected task status to be InProgress, but got %v", task.GetStatus())
	}

	// インターバルの削除
	task.RemoveInterval()
	if len(task.TimeIntervals) != 0 {
		t.Errorf("Expected TimeIntervals length to be 0, but got %d", len(task.TimeIntervals))
	}

	// インターバルの再作成
	task.ExecuteInterval()
	if task.GetStatus() != InProgress {
		t.Errorf("Expected task status to be InProgress, but got %v", task.GetStatus())
	}

	// インターバールの一時停止
	task.HoldInterval()
	if task.GetStatus() != OnHold {
		t.Errorf("Expected task status to be OnHold, but got %v", task.GetStatus())
	}

	// インターバルの再開
	task.ExecuteInterval()
	if task.GetStatus() != InProgress {
		t.Errorf("Expected task status to be InProgress, but got %v", task.GetStatus())
	}

	// インターバルの再開
	task.ExecuteInterval()
	if task.GetStatus() != Completed {
		t.Errorf("Expected task status to be Completed, but got %v", task.GetStatus())
	}

	// インターバルの削除
	task.RemoveInterval()
	if len(task.TimeIntervals) != 1 {
		t.Errorf("Expected TimeIntervals length to be 0, but got %d", len(task.TimeIntervals))
	}
}