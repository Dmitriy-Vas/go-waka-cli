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

	Range struct {
		Date     string `json:"date"`
		End      int    `json:"end"`
		Start    int    `json:"start"`
		Text     string `json:"text"`
		Timezone string `json:"timezone"`
	}

	Goals struct {
		Data []struct {
			AverageStatus string `json:"average_status"`
			Chart         struct {
				ActualSeconds     float64 `json:"actual_seconds"`
				ActualSecondsText string  `json:"actual_seconds_text"`
				GoalSeconds       int     `json:"goal_seconds"`
				GoalSecondsText   string  `json:"goal_seconds_text"`
				Range             Range   `json:"range"`
				RangeStatus       string  `json:"range_status"`
				RangeStatusReason string  `json:"range_status_reason"`
			} `json:"char_data"`
			CumulativeStatus string   `json:"cumulative_status"`
			Delta            string   `json:"delta"`
			ID               string   `json:"id"`
			IgnoreDays       []string `json:"ignore_days"`
			ImproveByPercent float64  `json:"improve_by_percent"`
			IsEnabled        bool     `json:"is_enabled"`
			Languages        []string `json:"languages"`
			Projects         []string `json:"projects"`
			RangeText        string   `json:"range_text"`
			Seconds          int      `json:"seconds"`
			Status           string   `json:"status"`
			Subscribers      []struct {
				Email          string `json:"email"`
				EmailFrequency string `json:"email_frequency"`
				FullName       string `json:"full_name"`
				UserID         string `json:"user_id"`
				UserName       string `json:"username"`
			} `json:"subscribers"`
		} `json:"data"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	}

	leadersUser struct {
		Rank         int `json:"rank"`
		RunningTotal struct {
			TotalSeconds              float64 `json:"total_seconds"`
			HumanReadableTotal        string  `json:"human_readable_format"`
			DailyAverage              float64 `json:"daily_average"`
			HumanReadableDailyAverage string  `json:"human_readable_daily_average"`
			Languages                 []struct {
				Name         string  `json:"name"`
				TotalSeconds float64 `json:"total_seconds"`
			}
		} `json:"running_total"`
		User struct {
			ID                string `json:"id"`
			Email             string `json:"email"`
			UserName          string `json:"username"`
			FullName          string `json:"full_name"`
			DisplayName       string `json:"display_name"`
			Site              string `json:"website"`
			HumanReadableSite string `json:"human_readable_website"`
			Location          string `json:"location"`
			EmailPublic       bool   `json:"email_public"`
			PhotoPublic       bool   `json:"photo_public"`
		} `json:"user"`
	}

	Leaders struct {
		CurrentUser leadersUser   `json:"current_user"`
		Data        []leadersUser `json:"data"`
		Page        int           `json:"page"`
		TotalPages  int           `json:"total_pages"`
		Range       struct {
			StartDate string `json:"start_date"`
			StartText string `json:"start_text"`
			EndDate   string `json:"end_date"`
			EndText   string `json:"end_text"`
			Name      string `json:"name"`
			Text      string `json:"text"`
		} `json:"range"`
		Language   string    `json:"language"`
		ModifiedAt time.Time `json:"modified_at"`
		Timeout    int       `json:"timeout"`
		WritesOnly bool      `json:"writes_only"`
	}

	Meta struct {
		Data struct {
			IPs struct {
				API     []string `json:"api"`
				Website []string `json:"website"`
				Worker  []string `json:"worker"`
			} `json:"ips"`
		} `json:"data"`
	}

	OrgDashboards struct {
		Data []struct {
			ID                          string    `json:"id"`
			FullName                    string    `json:"full_name"`
			CreatedBy                   string    `json:"created_by"`
			Timezone                    string    `json:"timezone"`
			HasChangedTImezone          bool      `json:"has_changed_timezone"`
			MembersCount                int       `json:"members_count"`
			MembersCountHumanReadable   string    `json:"members_count_human_readable"`
			IsCurrentUserMember         bool      `json:"is_current_user_member"`
			CanCurrentUserView          bool      `json:"can_current_user_view"`
			CanCurrentUserRequestToView bool      `json:"can_current_user_request_to_view"`
			CanCUrrentUserRequestToJoin bool      `json:"can_current_user_request_to_join"`
			CanCurrentUserAddMembers    bool      `json:"can_current_user_add_members"`
			CanCurrentUserRemoveMember  bool      `json:"can_current_user_remove_members"`
			CanCurrentUserDelete        bool      `json:"can_current_user_delete"`
			CreatedAt                   time.Time `json:"created_at"`
			ModifiedAt                  time.Time `json:"modified_at"`
		}
		NextPage   int `json:"next_page"`
		Page       int `json:"page"`
		PrevPage   int `json:"prev_page"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	}

	Orgs struct {
		Data []struct {
			ID                                       string    `json:"id"`
			Name                                     string    `json:"name"`
			DefaultProjectPrivacy                    string    `json:"default_project_privacy"`
			InvitedPeopleCount                       int       `json:"invited_people_count"`
			InvitedPeopleCountHumanReadable          string    `json:"invited_people_count_human_readable"`
			IsDurationVisible                        bool      `json:"is_duration_visible"`
			PeopleCount                              int       `json:"people_count"`
			PeopleCountHumanReadable                 string    `json:"people_count_human_readable"`
			Timeout                                  int       `json:"timeout"`
			Timezone                                 int       `json:"timezone"`
			WritesOnly                               int       `json:"writes_only"`
			CanCurrentUserListDashboards             bool      `json:"can_current_user_list_dashboards"`
			CanCurrentUserCreateDashboards           bool      `json:"can_current_user_create_dashboards"`
			CanCurrentUserDisplayCodingOnDashboards  bool      `json:"can_current_user_display_coding_on_dashboards"`
			CanCurrentUserViewAllDashboards          bool      `json:"can_current_user_view_all_dashboards"`
			CanCurrentUserAddPeopleToDashboards      bool      `json:"can_current_user_add_people_to_dashboards"`
			CanCurrentUserRemovePeopleFromDashboards bool      `json:"can_current_user_remove_people_from_dashboards"`
			CanCurrentUserEditAndDeleteDashboards    bool      `json:"can_current_user_edit_and_delete_dashboards"`
			CanCurrentUserAddPeopleToOrg             bool      `json:"can_current_user_add_people_to_org"`
			CanCurrentUserRemovePeopleFromOrg        bool      `json:"can_current_user_remove_people_from_org"`
			CanCurrentUserManageGroups               bool      `json:"can_current_user_manage_groups"`
			CanCurrentUserViewAuditLog               bool      `json:"can_current_user_view_audit_log"`
			CanCurrentUserEditOrg                    bool      `json:"can_current_user_edit_org"`
			CanCurrentUserManageBilling              bool      `json:"can_current_user_manage_billing"`
			CanCurrentUserDeleteOrg                  bool      `json:"can_current_user_delete_org"`
			CreatedAt                                time.Time `json:"created_at"`
			ModifiedAt                               time.Time `json:"modified_at"`
		} `json:"data"`
		NextPage   int `json:"next_page"`
		Page       int `json:"page"`
		PrevPage   int `json:"prev_page"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	}

	OrgDashboardMembers struct {
		Data []struct {
			ID         string `json:"id"`
			Email      string `json:"email"`
			FullName   string `json:"full_name"`
			IsViewOnly bool   `json:"is_view_only"`
			Photo      string `json:"photo"`
			Username   string `json:"username"`
		}
		NextPage   int `json:"next_page"`
		Page       int `json:"page"`
		PrevPage   int `json:"prev_page"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	}

	Stat struct {
		Name         string  `json:"name"`
		TotalSeconds float64 `json:"total_seconds"`
		Percent      float64 `json:"percent"`
		Digital      string  `json:"digital"`
		Text         string  `json:"text"`
		Hours        int     `json:"hours"`
		Minutes      int     `json:"minutes"`
		Seconds      int     `json:"seconds"`
	}

	OrgDashboardMemberSummaries struct {
		Data []struct {
			GrandTotal       Stat   `json:"grand_total"`
			Projects         []Stat `json:"projects"`
			Languages        []Stat `json:"languages"`
			Editors          []Stat `json:"editors"`
			OperatingSystems []Stat `json:"operating_systems"`
			Branches         []Stat `json:"branches"`
			Entities         []Stat `json:"entities"`
			Range            Range  `json:"struct"`
			Start            int    `json:"start"`
			End              int    `json:"end"`
		}
	}

	PrivateLeaderboards struct {
		Data []struct {
			CanDelete                 bool      `json:"can_delete"`
			CanEdit                   bool      `json:"can_edit"`
			CreatedAt                 time.Time `json:"created_at"`
			HasAvailableSeat          bool      `json:"has_available_seat"`
			ID                        string    `json:"id"`
			MembersCount              int       `json:"members_count"`
			MembersWithTimezonesCount int       `json:"members_with_timezones_count"`
			ModifiedAt                time.Time `json:"modified_at"`
			Name                      string    `json:"name"`
			TimeRange                 string    `json:"time_range"`
		}
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	}

	PrivateLeaderboardsLeaders struct {
		Data []struct {
			Rank         int `json:"rank"`
			RunningTotal struct {
				TotalSeconds              float64 `json:"total_seconds"`
				HumanReadableTotal        string  `json:"human_readable_total"`
				DailyAverage              float64 `json:"daily_average"`
				HumanReadableDailyAverage string  `json:"human_readable_daily_average"`
				Languages                 []struct {
					Name         string  `json:"name"`
					TotalSeconds float64 `json:"total_seconds"`
				} `json:"languages"`
			} `json:"running_total"`
			User struct {
				ID                   string `json:"id"`
				Email                string `json:"email"`
				Username             string `json:"username"`
				FullName             string `json:"full_name"`
				DisplayName          string `json:"display_name"`
				Website              string `json:"website"`
				HumanReadableWebsite string `json:"human_readable_website"`
				Location             string `json:"location"`
				EmailPublic          bool   `json:"email_public"`
				PhotoPublic          bool   `json:"photo_public"`
			} `json:"user"`
			Language   int       `json:"language"`
			ModifiedAt time.Time `json:"modified_at"`
			Page       int       `json:"page"`
			Range      struct {
				StartDate int    `json:"start_date"`
				StartText string `json:"start_text"`
				EndDate   int    `json:"end_date"`
				EndText   string `json:"end_text"`
				Name      string `json:"name"`
				Text      string `json:"text"`
			} `json:"range"`
			Timeout    int  `json:"timeout"`
			TotalPages int  `json:"total_pages"`
			WritesOnly bool `json:"writes_only"`
		}
	}

	timeRange struct {
		string
	}

	Stats struct {
		Data struct {
			TotalSeconds              float64 `json:"total_seconds"`
			HumanReadableTotal        string  `json:"human_readable_total"`
			DailyAverage              float64 `json:"daily_average"`
			HumanReadableDailyAverage string  `json:"human_readable_daily_average"`
			Projects                  []Stat  `json:"projects"`
			Languages                 []Stat  `json:"languages"`
			Editors                   []Stat  `json:"editors"`
			OperatingSystems          []Stat  `json:"operating_systems"`
			Dependencies              []Stat  `json:"dependencies"`
			Machines                  []struct {
				Machine struct {
					CreatedAt  time.Time `json:"created_at"`
					ID         string    `json:"id"`
					IP         string    `json:"ip"`
					LastSeenAt time.Time `json:"last_seen_at"`
					Name       string    `json:"name"`
					Value      string    `json:"value"`
				} `json:"machine"`
				Stat
			} `json:"machines"`
			BestDay struct {
				Date         string  `json:"date"`
				Text         string  `json:"text"`
				TotalSeconds float64 `json:"total_seconds"`
			} `json:"best_day"`
			Range                   string    `json:"range"`
			Holidays                int       `json:"holidays"`
			DaysIncludingHolidays   int       `json:"days_including_holidays"`
			DaysMinusHolidays       int       `json:"days_minus_holidays"`
			Status                  string    `json:"status"`
			IsAlreadyUpdating       bool      `json:"is_already_updating"`
			IsCodingActivityVisible bool      `json:"is_coding_activity_visible"`
			IsOtherUsageVisible     bool      `json:"is_other_usage_visible"`
			IsStuck                 bool      `json:"is_stuck"`
			IsUpToDate              bool      `json:"is_up_to_date"`
			Start                   time.Time `json:"start"`
			End                     time.Time `json:"end"`
			Timezone                string    `json:"timezone"`
			Timeout                 int       `json:"timeout"`
			WritesOnly              bool      `json:"writes_only"`
			UserID                  string    `json:"user_id"`
			UserName                string    `json:"username"`
			CreatedAt               time.Time `json:"created_at"`
			ModifiedAt              time.Time `json:"modified_at"`
		} `json:"data"`
	}
)

func (d *Data) format() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}
