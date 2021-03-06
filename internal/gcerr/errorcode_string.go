package gcerr

import "strconv"

const _ErrorCode_name = "OKUnknownNotFoundAlreadyExistsInvalidArgumentInternalUnimplementedFailedPreconditionPermissionDeniedResourceExhaustedCanceledDeadlineExceeded"

var _ErrorCode_index = [...]uint8{0, 2, 9, 17, 30, 45, 53, 66, 84, 100, 117, 125, 141}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
