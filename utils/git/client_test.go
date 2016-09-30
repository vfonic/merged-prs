package git

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testCase = `3d7dae8 Merge pull request #3524 from vsco/gearman-backfill
2ad2ce3 Merge pull request #3525 from vsco/WEB-8321
532fef7 Merge pull request #3522 from vsco/feed-service-for-all
8f48443 Merge pull request #3519 from vsco/klj/google-tag-manager
b04edd4 Merge pull request #3520 from vsco/ljk/x-client-uploader
ba0b52d Merge pull request #3517 from vsco/cgd/pin-put-changes
a542386 Merge pull request #3518 from vsco/ljk/feed-deadline-increase
3a0c214 Merge pull request #3509 from vsco/check-invalid-gridsites
fd48931 Merge pull request #3508 from vsco/ljk/fix-search-button
d7b2810 Merge pull request #3516 from vsco/include-src-tests
e6ce78b Merge pull request #3515 from vsco/ljk/ingest-vsco
40b28ac Merge pull request #3507 from vsco/ljk/search-grpc-base
be3da80 Merge pull request #3514 from vsco/ljk/grpc-vsco
ca24d51 Merge pull request #3512 from vsco/cgd/pin-test2
d15936a Merge pull request #3510 from vsco/cgd/fix-pin-tests
ac31f5c Merge pull request #3505 from vsco/cub/pagination-arrows-darker
65781df Merge pull request #3448 from vsco/bms/careers-order-by-open-jobs
`

func TestProcessMergeCommits(t *testing.T) {
	r := bytes.NewReader([]byte(testCase))
	c := New()
	list := c.processMergeCommits(r)

	assert.Equal(t, "3d7dae8", list[0].Ref)
	assert.Equal(t, "Merge pull request #3524 from vsco/gearman-backfill", list[0].Message)
	assert.Equal(t, "65781df", list[16].Ref)
	assert.Equal(t, "Merge pull request #3448 from vsco/bms/careers-order-by-open-jobs", list[16].Message)
}
