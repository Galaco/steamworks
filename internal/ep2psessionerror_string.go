// Code generated by "stringer -type EP2PSessionError -trimprefix EP2PSessionError_"; DO NOT EDIT.

package internal

import "strconv"

const _EP2PSessionError_name = "NoneNotRunningAppNoRightsToAppDestinationNotLoggedInTimeoutMax"

var _EP2PSessionError_index = [...]uint8{0, 4, 17, 30, 52, 59, 62}

func (i EP2PSessionError) String() string {
	if i < 0 || i >= EP2PSessionError(len(_EP2PSessionError_index)-1) {
		return "EP2PSessionError(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EP2PSessionError_name[_EP2PSessionError_index[i]:_EP2PSessionError_index[i+1]]
}