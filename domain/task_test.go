package domain

import "testing"

func TestNewTask(t *testing.T) {
	cases := []struct {
		title string
		description string
		wantErr bool
	}{
		{"有効なタイトル", "説明", false},
		{"", "説明", true},
	}

	for _, c := range cases {
		task, err := NewTask(c.title, c.description)
		if c.wantErr {
			if err == nil {
				t.Errorf("タイトル=%q: エラーが発生するはずでしたが、発生しませんでした", c.title)
			}
		} else {
			if err != nil {
				t.Errorf("タイトル=%q: エラーは発生しないはずでしたが、発生しました: %v", c.title, err)
			}
			if task == nil {
				t.Errorf("タイトル=%q: タスクが生成されるはずですが、nilになっています", c.title)
			} else {
				if task.Title != c.title {
					t.Errorf("タイトル=%q: 期待したタイトル=%qですが、実際は%qでした", c.title, c.title, task.Title)
				}
				if task.CreatedAt.IsZero() {
					t.Error("作成日時がゼロ値です。作成日時はセットされているはずです")
				}
			}
		}
	}
}

func TestTaskUpdate(t *testing.T) {
	task, err := NewTask("初期タイトル", "初期説明")
	if err != nil {
		t.Fatalf("タスク作成に失敗しました: %v", err)
	}

	err = task.Update("", "新しい説明(タイトル空)")
	if err == nil {
		t.Error("タイトルを空文字で更新しようとしましたが、エラーが発生しませんでした")
	}

	err = task.Update("更新後のタイトル", "更新後の説明")
	if err != nil {
		t.Errorf("タスク更新でエラーが発生しました: %v", err)
	}
	if task.Title != "更新後のタイトル" {
		t.Errorf("タイトル更新が期待どおりではありません。期待=%q, 実際=%q", "更新後のタイトル", task.Title)
	}
	if task.UpdatedAt.Before(task.CreatedAt) {
		t.Error("更新日時が作成日時より前になっています。更新日時は作成日時以降であるはずです")
	}
}