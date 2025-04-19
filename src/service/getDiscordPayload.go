package service

import (
	"go-html-monitor/src/model"
	"go-html-monitor/src/util"
	"log/slog"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func getDiscordPayload(link, web string) (res model.DiscordPayload) {
	// Jalankan browser headless
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	// Buka halaman LinkedIn job dengan query filter
	page := browser.MustPage(link).MustWaitLoad()

	// Tunggu beberapa detik secara random biar ga keliatan bot banget
	util.SleepRandom(2, 4)

	// Tunggu konten siap
	slog.Info("Waiting selector to visible")

	switch web {
	case "linkedInJob":
		// Tunggu sampai job card pertama muncul
		page.MustElement(".base-search-card__title").MustWaitVisible()

		// Ambil elemen job card
		titles := page.MustElements(".base-search-card__title")
		companies := page.MustElements(".base-search-card__subtitle")
		links := page.MustElements("a.base-card__full-link")
		times := page.MustElements(".job-search-card__listdate, .job-search-card__listdate--new")

		listJob := []model.Job{}
		for i := range titles {
			title := strings.TrimSpace(titles[i].MustText())
			company := strings.TrimSpace(companies[i].MustText())
			link, _ := links[i].Attribute("href")
			if link == nil {
				slog.Warn("Link empty", "title", title)
				continue
			}
			if i >= len(times) {
				slog.Warn("Skip job: no post time found", "index", i)
				continue
			}
			postTime := strings.ToLower(strings.TrimSpace(times[i].MustText()))

			// Skip kalau bukan postingan baru (< 24 jam)
			if !isRecent(postTime) {
				continue
			}

			job := model.Job{
				Title:     title,
				Company:   company,
				Link:      *link,
				PostedAgo: postTime,
			}
			jobId := job.ExtractID(*link)
			if util.CheckIdExists(web, jobId) {
				slog.Warn("ID job duplicate (already processed).")
				continue
			} else {
				util.AddCache(web, jobId)
			}

			slog.Debug("Job Parsed", "title", job.Title, "company", job.Company)
			listJob = append(listJob, job)
			util.SleepRandom(1, 2)
		}

		if len(listJob) == 0 {
			slog.Info("😪 No jobs found.")
			return
		}

		slog.Info("✅ Selesai parsing", "total jobs", len(titles), "valid jobs", len(listJob), "skipped", len(titles)-len(listJob))
		res = model.Job_BuildDiscordPayload(listJob)
	}
	return
}
