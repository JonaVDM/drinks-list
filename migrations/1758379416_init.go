package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		// https://github.com/pocketbase/pocketbase/blob/develop/core/settings_model.go
		// most of the sensible settings have been put here. stuff that require secrets (eg: smtp, s3) should still be done
		// via the ui.
		settings.Meta.AppName = "Drankjes lijst"
		settings.Meta.AppURL = "https://drankjes.sv-rt.nl"
		settings.Meta.HideControls = true
		settings.Meta.SenderName = "KanCo"
		settings.Meta.SenderAddress = "kantoor@sv-realtime.nl"

		settings.Backups.Cron = "0 0 * * 0" // once a week on sunday, https://crontab.guru
		settings.Backups.CronMaxKeep = 3

		settings.Logs.MaxDays = 31
		settings.Logs.LogIP = true
		settings.Logs.LogAuthId = false
		settings.Logs.MinLevel = 0 // -4 = debug; 0 = info; 4 = warn; 8 = error

		// Watch out! Dragons!
		// Batch api is still experimental, what's down here are the defaults.
		//
		// settings.Batch.Enabled = false
		// settings.Batch.MaxBodySize = 128 << 20
		// settings.Batch.MaxRequests = 50
		// settings.Batch.Timeout = 3

		return app.Save(settings)
	}, func(app core.App) error {
		return nil
	})
}
