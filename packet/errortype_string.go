// Code generated by "stringer -type=ErrorType"; DO NOT EDIT.

package packet

import "strconv"

const (
	_ErrorType_name_0 = "WOW_SUCCESS"
	_ErrorType_name_1 = "WOW_FAIL_BANNEDWOW_FAIL_UNKNOWN_ACCOUNTWOW_FAIL_INCORRECT_PASSWORDWOW_FAIL_ALREADY_ONLINEWOW_FAIL_NO_TIMEWOW_FAIL_DB_BUSYWOW_FAIL_VERSION_INVALIDWOW_FAIL_VERSION_UPDATEWOW_FAIL_INVALID_SERVERWOW_FAIL_SUSPENDEDWOW_FAIL_FAIL_NOACCESSWOW_SUCCESS_SURVEYWOW_FAIL_PARENTCONTROLWOW_FAIL_LOCKED_ENFORCEDWOW_FAIL_TRIAL_ENDEDWOW_FAIL_USE_BATTLENETWOW_FAIL_ANTI_INDULGENCEWOW_FAIL_EXPIREDWOW_FAIL_NO_GAME_ACCOUNTWOW_FAIL_CHARGEBACKWOW_FAIL_INTERNET_GAME_ROOM_WITHOUT_BNETWOW_FAIL_GAME_ACCOUNT_LOCKEDWOW_FAIL_UNLOCKABLE_LOCK"
	_ErrorType_name_2 = "WOW_FAIL_CONVERSION_REQUIRED"
	_ErrorType_name_3 = "WOW_FAIL_DISCONNECTED"
)

var (
	_ErrorType_index_1 = [...]uint16{0, 15, 39, 66, 89, 105, 121, 145, 168, 191, 209, 231, 249, 271, 295, 315, 337, 361, 377, 401, 420, 460, 488, 512}
)

func (i ErrorType) String() string {
	switch {
	case i == 0:
		return _ErrorType_name_0
	case 3 <= i && i <= 25:
		i -= 3
		return _ErrorType_name_1[_ErrorType_index_1[i]:_ErrorType_index_1[i+1]]
	case i == 32:
		return _ErrorType_name_2
	case i == 255:
		return _ErrorType_name_3
	default:
		return "ErrorType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
