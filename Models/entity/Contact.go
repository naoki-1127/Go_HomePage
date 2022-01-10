package entity

import "time"

// DB上のテーブル、カラムと構造体との関連付けが自動的に行われる

type Contact struct {
	Created_at time.Time
	Updated_at time.Time
}
