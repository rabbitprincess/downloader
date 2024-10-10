package downloader

import (
	"context"
	"net/url"
	"os"

	_ "github.com/rclone/rclone/backend/all"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configfile"
	"github.com/rclone/rclone/fs/sync"
	_ "github.com/rclone/rclone/lib/plugin"
)

func init() {
	configfile.Install()
}

func Download(ctx context.Context, src, dst string) error {
	var err error
	fsInfo, err := fs.NewFs(ctx, src)
	if err != nil {
		return err
	}
	if isLocalPath(dst) {
		if err = os.MkdirAll(dst, os.ModePerm); err != nil {
			return err
		}
	}

	destinationFs, err := fs.NewFs(ctx, dst)
	if err != nil {
		return err
	}

	err = sync.CopyDir(ctx, destinationFs, fsInfo, false)
	if err != nil {
		return err
	}
	return nil
}

func isLocalPath(path string) bool {
	parsedURL, err := url.Parse(path)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Scheme == "file" {
		return true
	}

	return false
}

// func copyFile(ctx context.Context, url string, dst string) error {
// 	fsrc, err := fs.NewFs(ctx, url)
// 	if err != nil {
// 		return err
// 	}

// 	fdst, err := fs.NewFs(ctx, dst)
// 	if err != nil {
// 		return err
// 	}

// 	isDir := false
// 	objs, err := fs.ListDir(ctx, fsrc, "")
// 	if err == nil && len(objs) > 0 {
// 		// 디렉토리일 경우
// 		isDir = true
// 	}

// 	operations.CopyURL()
// 	operations.Copy()

// 	if err != nil {
// 		// 파일이 아닌 디렉토리인 경우 Walk를 사용해 복사
// 		err = operations.CopyDir(ctx, fdst, fsrc, true)
// 		if err != nil {
// 			return fmt.Errorf("failed to copy directory: %w", err)
// 		}
// 	} else {
// 		// 파일인 경우 CopyFile 사용
// 		err = operations.CopyFile(ctx, fdst, fsrc, srcInfo.Remote(), srcInfo.Remote())
// 		if err != nil {
// 			return fmt.Errorf("failed to copy file: %w", err)
// 		}
// 	}

// 	return nil
// }
