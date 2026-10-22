-- +goose Up

-- Оценка задачи в минутах. Одни сутки в интерфейсе — 8 часов, здесь только минуты.
ALTER TABLE kanban_card
    ADD COLUMN estimate_minutes INT NULL
        CHECK (estimate_minutes IS NULL OR estimate_minutes > 0);

-- Каждая запись прибавляется к сумме. tracked_at — выбранная дата, created_at — момент записи.
CREATE TABLE kanban_card_time_wasting (
    id         BIGSERIAL PRIMARY KEY,
    card_id    BIGINT NOT NULL REFERENCES kanban_card(id) ON DELETE CASCADE,
    user_id    BIGINT NOT NULL,
    minutes    INT NOT NULL CHECK (minutes > 0),
    tracked_at TIMESTAMPTZ(0) NOT NULL,
    created_at TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ(0)
);

CREATE INDEX kanban_card_time_wasting_card_id_idx
    ON kanban_card_time_wasting (card_id)
    WHERE deleted_at IS NULL;

-- +goose Down

DROP TABLE IF EXISTS kanban_card_time_wasting;
ALTER TABLE kanban_card DROP COLUMN IF EXISTS estimate_minutes;
