package thumbnail_test

import (
	"testing"

	"github.com/m0a/myGoSandBox/mp4server/thumbnail"

	// thumbnail "github.com/m0a/myGoSandBox/mp4server/thumbnail"
)

func TestTempDirMake(t *testing.T) {
	name := thumbnail.ThumbnailTempDirPath
	t.Logf("name=[%s]\n", name)
}
