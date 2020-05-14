package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"pet-paradise/utils"
)

var (
	ROOT        = Agent.Root
	DATE_FORMAT = "2006-01-02 15:04:05"
)

func Get(ctx *gin.Context) {
	path := ctx.Query("path")

	if exists := pathExists(ROOT + path); !exists {
		utils.Success(ctx, "no such directory", nil)
		return
	}

	ctx.File(ROOT + path)
}

func Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	path := ctx.PostForm("path")

	if err := ctx.SaveUploadedFile(file, ROOT+path+file.Filename); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func List(ctx *gin.Context) {
	dir := ctx.Query("dir")
	if exists := pathExists(ROOT + dir); !exists {
		utils.Success(ctx, "no such directory", nil)
		return
	}

	fileInfo, err := getAllFile(ROOT + dir)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", fileInfo)
}

func Delete(ctx *gin.Context) {
	path := ctx.Query("path")
	if exists := pathExists(ROOT + path); !exists {
		utils.Success(ctx, "no such file", nil)
		return
	}

	if err := os.Remove(path); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

type FileInfo struct {
	Name       string `json:"name"`
	IsDir      bool   `json:"is_dir"`
	Size       int64  `json:"size"`
	ModifiedAt string `json:"modified_at"`
}

func getAllFile(dir string) ([]FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fileInfo []FileInfo
	for _, fi := range files {
		fileInfo = append(fileInfo, FileInfo{
			Name:       fi.Name(),
			IsDir:      fi.IsDir(),
			Size:       fi.Size(),
			ModifiedAt: fi.ModTime().Format(DATE_FORMAT),
		})
	}
	return fileInfo, nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
