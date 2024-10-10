package downloader

import (
	"context"
	"testing"

	_ "github.com/rclone/rclone/backend/all"
	"github.com/stretchr/testify/require"
)

func TestDownloaderLocal(t *testing.T) {
	src := "./testfile"
	dst := "./testfileb"

	err := Download(context.Background(), src, dst)
	require.NoError(t, err)
}

func TestDownloaderHTTP(t *testing.T) {
	src := ":http,url=http://speedtest.tele2.net/1MB.zip"
	dst := "./testfilec"

	err := Download(context.Background(), src, dst)
	require.NoError(t, err)
}

func TestDownloaderFTP(t *testing.T) {

}
