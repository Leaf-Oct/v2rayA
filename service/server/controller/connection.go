package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/v2rayA/v2rayA/common"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/util/log"
	"github.com/v2rayA/v2rayA/server/service"
	"os/exec"
	"runtime"
)

func PostConnection(ctx *gin.Context) {
	updatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		updatingMu.Unlock()
		return
	}
	updating = true
	updatingMu.Unlock()
	defer func() {
		updatingMu.Lock()
		updating = false
		updatingMu.Unlock()
	}()

	var which configure.Which
	err := ctx.ShouldBindJSON(&which)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	err = service.Connect(&which)
	if err != nil {
		log.Warn("PostConnection: %v", err)
		common.ResponseError(ctx, logError(fmt.Errorf("failed to connect: %w", err)))
		return
	}
	getTouch(ctx)
}

func DeleteConnection(ctx *gin.Context) {
	updatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		updatingMu.Unlock()
		return
	}
	updating = true
	updatingMu.Unlock()
	defer func() {
		updatingMu.Lock()
		updating = false
		updatingMu.Unlock()
	}()

	var which configure.Which
	err := ctx.ShouldBindJSON(&which)
	if err != nil {
		common.ResponseError(ctx, logError("bad request"))
		return
	}
	err = service.Disconnect(which, false)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	getTouch(ctx)
}

func PostV2ray(ctx *gin.Context) {
	updatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		updatingMu.Unlock()
		return
	}
	updating = true
	updatingMu.Unlock()
	defer func() {
		updatingMu.Lock()
		updating = false
		updatingMu.Unlock()
	}()

	err := service.StartV2ray()
	if err != nil {
		common.ResponseError(ctx, logError(fmt.Errorf("failed to start v2ray-core: %w", err)))
		return
	}
	// 判断当前系统，如果是linux，则通过下面那个指令打开设置里的系统代理(貌似仅限gnome桌面环境，如果不是gnome的话，根本没有这个指令，执行了也无妨)
	if runtime.GOOS == "linux" {
		cmd := exec.Command("gsettings", "set", "org.gnome.system.proxy", "mode", "manual")
		err := cmd.Run()
		if err != nil {
			log.Warn("cmd.Run() failed with %s\n", err)
		}
	}
	getTouch(ctx)
}

func DeleteV2ray(ctx *gin.Context) {
	updatingMu.Lock()
	if updating {
		common.ResponseError(ctx, processingErr)
		updatingMu.Unlock()
		return
	}
	updating = true
	updatingMu.Unlock()
	defer func() {
		updatingMu.Lock()
		updating = false
		updatingMu.Unlock()
	}()

	err := service.StopV2ray()
	if err != nil {
		common.ResponseError(ctx, logError(fmt.Errorf("failed to stop v2ray-core: %w", err)))
		return
	}
	// 判断当前系统，如果是linux，则通过下面那个指令关闭设置里的系统代理(貌似仅限gnome桌面环境，如果不是gnome的话，根本没有这个指令，执行了也无妨)
	if runtime.GOOS == "linux" {
		cmd := exec.Command("gsettings", "set", "org.gnome.system.proxy", "mode", "none")
		err := cmd.Run()
		if err != nil {
			log.Warn("cmd.Run() failed with %s\n", err)
		}
	}
	getTouch(ctx)
}
