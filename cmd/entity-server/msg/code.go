package msg

const (
	SUCCESS                        = 200
	ERROR                          = 500
	INVALID_PARAMS                 = 400
	ERROR_OKTA_CHECK_TOKEN_TIMEOUT = 415
	ERROR_EXIST                    = 1001
	ERROR_EXIST_FAIL               = 1002
	ERROR_NOT_EXIST                = 1003
	ERROR_GET_FAIL                 = 1004
	ERROR_COUNT_FAIL               = 1005
	ERROR_ADD_FAIL                 = 1006
	ERROR_EDIT_FAIL                = 1007
	ERROR_DELETE_FAIL              = 1008
	ERROR_EXPORT_FAIL              = 1009
	ERROR_IMPORT_FAIL              = 1010

	ERROR_AUTH_FAIL                = 2001
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 2002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 2003
	ERROR_AUTH_TOKEN               = 2004
	ERROR_AUTH                     = 2005
)
