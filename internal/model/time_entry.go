package model

import "time"

// TimeEntry — одно списание на задачу. Сумма минут по карточке — это «списано».
type TimeEntry struct {
	ID      int64
	CardID  int64
	UserID  int64
	Minutes   int
	TrackedAt time.Time
}
