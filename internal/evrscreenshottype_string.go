// Code generated by "stringer -type EVRScreenshotType -trimprefix EVRScreenshotType_"; DO NOT EDIT.

package internal

import "strconv"

const _EVRScreenshotType_name = "NoneMonoStereoMonoCubemapMonoPanoramaStereoPanorama"

var _EVRScreenshotType_index = [...]uint8{0, 4, 8, 14, 25, 37, 51}

func (i EVRScreenshotType) String() string {
	if i < 0 || i >= EVRScreenshotType(len(_EVRScreenshotType_index)-1) {
		return "EVRScreenshotType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EVRScreenshotType_name[_EVRScreenshotType_index[i]:_EVRScreenshotType_index[i+1]]
}