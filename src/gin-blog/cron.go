package main

import (
	"gin_example/src/gin-blog/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func startCron() {
	log.Println("Starting...")

	c := cron.New()
	var err error
	err = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	if err != nil {
		log.Fatal("Run models.CleanAllTag 定时任务开启失败")
	}
	err = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	if err != nil {
		log.Fatal("Run models.CleanAllArticle 定时任务开启失败")
	}
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
