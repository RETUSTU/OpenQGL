package main

import (
	"context"
	"sync"
)

// App struct
type App struct {
	ctx              context.Context
	downloadProgress DownloadProgress
	downloadList     []DownloadItem
	downloadMutex    sync.Mutex
	isDownloading    bool
	downloadCancel   chan struct{}
	guestUnlocked    bool
	launchLogPath    string // 当前启动日志路径（版本目录下的 QGL\Logs）
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		downloadList:  make([]DownloadItem, 0),
		guestUnlocked: false,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
