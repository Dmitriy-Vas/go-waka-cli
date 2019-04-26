package api

import (
	"fmt"
	"net/http"
	"time"
)

type (
	WakaClient struct {
		client *http.Client
		token  string
	}

	Data struct {
		Day   int
		Month int
		Year  int
	}

	// Список проектов пользователя
	Projects struct {
		// Список в виде массива
		Data []struct {
			// Время создания проекта
			Created time.Time `json:"created_at"`
			// Отформатированное название проекта
			EscapedName string `json:"html_escaped_name"`
			// ID проекта
			ID string `json:"id"`
			// Название проекта
			Name string `json:"name"`
			// Ссылка на репозиторий проекта
			Repository string `json:"repository"`
			// Ссылка на проект для использования с WakApi
			URL string `json:"url"`
		} `json:"data"`
	}

	// Пользовательская информация
	User struct {
		Data struct {
			// ID пользователя
			ID string `json:"id"`
			// Когда был создан аккаунт
			Created time.Time `json:"created_at"`
			// Когда аккаунт был изменён
			ModifiedAt time.Time `json:"modified_at"`
			// Ник пользователя
			UserName string `json:"username"`
			// Отображаемое имя
			DisplayName string `json:"display_name"`
			// Имя пользователя
			Name string `json:"full_name"`
			// Часовой пояс пользователя
			Timezone string `json:"timezone"`
			// Местонахождение пользователя
			Location string `json:"location"`
			// Почта пользователя
			Email string `json:"email"`
			// Публичная ли почта
			EmailPublic bool `json:"email_public"`
			// Подтверждён ли аккаунт
			Confirmed bool `json:"is_email_confirmed"`
			// Имеется ли премиум
			Premium bool `json:"has_premium_features"`
			// Нужно ли указать способ оплаты
			NeedPaymentMethod bool `json:"need_payment_method"`
			// Тариф пользователя
			Plan string `json:"plan"`
			// Ссылка на аватар пользователя
			Photo string `json:"photo"`
			// Публичное ли фото
			PhotoPublic bool `json:"photo_public"`
			// Публичная ли статистика по языкам программирования
			LanguagesPublic bool `json:"languages_used_public"`
			// Публичное ли время входа
			LoggedTimePublic bool `json:"logged_time_public"`
			// Ссылка на сайт пользователя
			Site string `json:"website"`
			// Упрощённая ссылка на сайт пользователя
			ReadableSite string `json:"human_readable_website"`
			// Ищет ли пользователь работу
			Hireable bool `json:"is_hireable"`
			// Последнее обновление статистики
			LastBeat time.Time `json:"last_heartbeat_at"`
			// Техническая информация о последнем использованном плагине
			LastPlugin string `json:"last_plugin"`
			// Название последнего использованного плагина
			LastPluginName string `json:"last_plugin_name"`
			// Название последнего проекта
			LastProject string `json:"last_project"`
		} `json:"data"`
	}

	Commits struct {
		// todo add commits
	}

	// Активность пользователя на выбранный день
	Durations struct {
		// Список веток проектов
		Branches []string  `json:"branches"`
		End      time.Time `json:"end"`
		Start    time.Time `json:"start"`
		TimeZone string    `json:"timezone"`
		// Основная информация о затраченном времени
		// на каждый проект
		Data []struct {
			// Время создания проекта
			Created time.Time `json:"created_at"`
			// Позиция курсора
			Position string `json:"cursorpos"`
			// Время работы над проектом в секундах
			Duration float64 `json:"duration"`
			// ID проекта
			ID string `json:"id"`
			// Номер строки на которую наведён курсор
			Lineno int `json:"lineno"`
			// ID компьютера на котором велась работа
			MachineID string `json:"machine_name_id"`
			// Название проекта
			Project string `json:"project"`
			// Время начала работы над проектом
			Time float64 `json:"time"`
			// ID пользователя
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	Heartbeat struct {
		Entity       string  `json:"entity"`
		Type         string  `json:"type"`
		Category     string  `json:"category"`
		Time         float64 `json:"time"`
		Project      string  `json:"project"`
		Branch       string  `json:"branch"`
		Language     string  `json:"language"`
		Dependencies string  `json:"dependencies"`
		Lines        int     `json:"lines"`
		Lineno       int     `json:"lineno"`
		Cursorpos    int     `json:"cursorpos"`
		IsWrite      bool    `json:"is_write"`
	}

	Heartbeats struct {
		Data     []*Heartbeat `json:"data"`
		Start    int          `json:"start"`
		End      int          `json:"end"`
		Timezone string       `json:"timezone"`
	}

	HeartbeatResponse struct {
		ID     string  `json:"id"`
		Entity string  `json:"entity"`
		Type   string  `json:"type"`
		Time   float64 `json:"time"`
	}
)

func (d *Data) format() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}
