package util

const SubsliceLabelsList = []string{
	"google.com/gke-tpu-sub-slice-2x2",
	"google.com/gke-tpu-sub-slice-2x4",
	"google.com/gke-tpu-sub-slice-4x4",
	"google.com/gke-tpu-sub-slice-4x8",
	"google.com/gke-tpu-sub-slice-8x8",
	"google.com/gke-tpu-sub-slice-8x16",
	"google.com/gke-tpu-sub-slice-16x16",
	"google.com/gke-tpu-sub-slice-2x2x1",
	"google.com/gke-tpu-sub-slice-2x2x2",
	"google.com/gke-tpu-sub-slice-2x2x4",
	"google.com/gke-tpu-sub-slice-2x4x4",
	"google.com/gke-tpu-sub-slice-4x4x4",
}

const (
	ControllerDriverName     = "tpu.google.com" 
	ClusterResourceSliceName = "subslice-tpu-resources" 
)

// set for quick lookup
var SubsliceTPULabelsSet = make(map[string]bool)

func init() {
	for _, label := range SubsliceLabelsList {
		SubsliceTPULabelsSet[label] = true
	}
}