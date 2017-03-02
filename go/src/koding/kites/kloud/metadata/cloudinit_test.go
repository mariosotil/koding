package metadata_test

import (
	"flag"
	"testing"
)

var update = flag.Bool("update-golden", false, "Update golden files.")

func TestParseCloudInit(t *testing.T) {

}

func TestCloudInitMergeIn(t *testing.T) {

}

func TestUpdateGolden(t *testing.T) {
	if !*update {
		return
	}
}
