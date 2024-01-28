package errs

type ErrorCode = string // Use as naming

const (
	DuplicateBrandName = "duplicate_brand_name"

	DuplicateTagName = "duplicate_tag_name"

	OnlyOneCompanyAllowed = "only_one_company_allowed"

	DuplicateAccessGroupName = "duplicate_access_group_name"
)
