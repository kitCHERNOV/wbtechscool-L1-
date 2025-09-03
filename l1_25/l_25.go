package l1_25

import (
	"context"
	"fmt"
	"time"
)

// Реализация своего таймера

func MainFuncTimer(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 2)
	select {
	case <-ctx.Done():
		return
	case <-ticker.C:
		fmt.Println("doing something after ticker")
	case <-time.After(time.Second * 13):
		return
	}
}
