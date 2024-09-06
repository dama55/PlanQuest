package models

import (
	"testing"
)

// テスト：タスクをモジュールに追加し、ステータスを確認する
func TestAddTask(t *testing.T) {
	// モジュールを初期化
	module := InitModule().(*Module)

	// タスクを追加
	module.AddTask()

	// 追加されたタスクの確認
	if len(module.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(module.Tasks))
	}

	// モジュールのステータスがNotStartedであることを確認
	if module.StatusVal != NotStarted {
		t.Errorf("Expected module status to be NotStarted, got %v", module.StatusVal)
	}
}

// テスト：タスクを削除し、ステータスを確認する
func TestRemoveTask(t *testing.T) {
	// モジュールを初期化
	module := InitModule().(*Module)

	// タスクを追加
	module.AddTask()
	taskId := module.Tasks[0].GetId()

	// タスクを削除
	err := module.RemoveTask(taskId)
	if err != nil {
		t.Fatalf("Failed to remove task: %v", err)
	}

	// タスクが削除されたことを確認
	if len(module.Tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(module.Tasks))
	}
}

// テスト：タスクのステータスに応じてモジュールのステータスが正しく変更されるか確認
func TestCheckStatus(t *testing.T) {
	// モジュールを初期化
	module := InitModule().(*Module)

	// タスクを追加
	module.AddTask()
	module.AddTask()

	// 全タスクがNotStartedの状態を確認
	module.CheckStatus()
	if module.StatusVal != NotStarted {
		t.Errorf("Expected module status to be NotStarted, got %v", module.StatusVal)
	}

	// 最初のタスクのステータスをInProgressに直接変更
	task := module.Tasks[0].(*Task)
	task.StatusVal = InProgress
	module.CheckStatus()
	if module.StatusVal != InProgress {
		t.Errorf("Expected module status to be InProgress, got %v", module.StatusVal)
	}

	// 全タスクをCompletedに変更
	for _, t := range module.Tasks {
		task := t.(*Task)
		task.StatusVal = Completed
	}
	module.CheckStatus()
	if module.StatusVal != Completed {
		t.Errorf("Expected module status to be Completed, got %v", module.StatusVal)
	}
}