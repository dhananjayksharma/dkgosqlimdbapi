package consts

const (
	InvalidCode       = "Invalid code input request "
	InvalidUpdateType = "Invalid update type input request "

	InvalidUpdateData = "Invalid data found for updating request"

	ErrorUpdateType         = "Update failed for merchant code %v"
	ErrorUpdateTypeNotFound = "Update type not found as %v"

	ErrorUpdateMember = "Update failed for members code %v, email %v"

	ErrorDataNotFoundCode = "Merchant data not found for given code %v"

	ErrorUserNotFoundCode = "User not found for given code %v"

	InvalidUserDocId    = "Invalid user code input request "
	InvalidMerchantCode = "Invalid client code input request "
	InvalidAppCode      = "Invalid client app code input request "

	UserLoginSuccess       = "User login successfully"
	UserAddedSuccess       = "User added successfully"
	TokenRegeneatedSuccess = "Token Regenerated successfully"

	MovieAddedSuccess   = "Movie added successfully"
	MovieUpdatedSuccess = "Movie updated successfully for code: %v"
	PageLimitMessage    = "invalid page limit"
	SkipMessage         = "invalid skip limit"

	ActiveStatus     uint8  = 1 // Active
	DeactiveStatus   uint8  = 0 // DeactiveStatus
	ArchiveStatus    uint8  = 9 // ArchiveStatus
	ParseLayoutISO   string = "2006-01-02"
	DateFormatLayout string = "02-01-2006"

	// MySQL
	DuplicateEntry           string = "Duplicate entry"
	ErrUserAlreadyExists     string = "Person already exists for email: %v"
	ErrMerchantAlreadyExists string = "Merchant already exists for merchant code: %v"
	// ErrUserAlreadyExists string = "User already exists for input"
)
