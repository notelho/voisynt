package error

type StatusCode = int
type StatusMessage = string

type StatusError struct {
	Code    StatusCode
	Message StatusMessage
}

type Error struct {
	TempDirFailed      *StatusError
	ExecRunFailed      *StatusError
	ArgsMessageMissing *StatusError
	ArgsOutputMissing  *StatusError
	PathNotFound       *StatusError
}

var Exceptions = &Error{
	TempDirFailed: &StatusError{
		Message: "Failed to create a temporary directory",
		Code:    1,
	},
	ExecRunFailed: &StatusError{
		Message: "Failed to run a CLI exec command",
		Code:    1,
	},
	ArgsMessageMissing: &StatusError{
		Message: "The --message attribute is required",
		Code:    1,
	},
	ArgsOutputMissing: &StatusError{
		Message: "The --output attribute is required",
		Code:    1,
	},
	PathNotFound: &StatusError{
		Message: "Cant resolve provided path",
		Code:    1,
	},
}
