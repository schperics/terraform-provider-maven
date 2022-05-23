package provider

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDownloadFromMavenCentral(t *testing.T) {
	r := NewRepository("", "", "")
	a := NewArtifact("org.apache.commons", "commons-text", "1.9", "", "")
	path, err := DownloadMavenPackage(r, a, "")
	assert.Equal(t, "commons-text-1.9.jar", path)
	assert.Nil(t, err)
	fi, err := os.Stat(path)
	assert.Positive(t, fi.Size())
}

func TestDownloadWithDir(t *testing.T) {
	r := NewRepository("", "", "")
	a := NewArtifact("org.apache.commons", "commons-text", "1.9", "", "")
	path, err := DownloadMavenPackage(r, a, "out")
	assert.Equal(t, "out/commons-text-1.9.jar", path)
	assert.Nil(t, err)
	fi, err := os.Stat(path)
	assert.Positive(t, fi.Size())
}

func TestDownloadNotFound(t *testing.T) {
	r := NewRepository("", "", "")
	a := NewArtifact("invalid.group", "commons-text", "1.9", "", "")
	path, err := DownloadMavenPackage(r, a, "")
	assert.Equal(t, "", path)
	assert.Equal(t, "status code 404 returned. URL: https://repo1.maven.org/maven2/invalid/group/commons-text/1.9/commons-text-1.9.jar", err.Error())
}

func TestDownloadFromPrivateRepository(t *testing.T) {
	r := NewRepository("", "", "")
	a := NewArtifact("invalid.group", "commons-text", "1.9", "", "")
	path, err := DownloadMavenPackage(r, a, "")
	assert.Equal(t, "", path)
	assert.Equal(t, "status code 404 returned. URL: https://repo1.maven.org/maven2/invalid/group/commons-text/1.9/commons-text-1.9.jar", err.Error())
}