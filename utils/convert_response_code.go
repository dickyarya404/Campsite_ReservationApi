package utils

import (
	"campsite_reservation/constant"
	"net/http"
)

func ConvertResponseCode(err error) int {
	var badRequestErrors = []error{
		constant.ErrInvalidUsernameOrPassword,
		constant.ErrAllFieldsMustBeFilled,
		constant.ErrEmailAlreadyExist,
		constant.ErrUsernameAlreadyExist,
		// constant.ErrMaxFileSize,
		// constant.ErrMaxFileUpload,
		constant.ErrLimitAndPageMustBeFilled,
		constant.ErrIDMustBeFilled,
		constant.ErrReportNotFound,
		constant.ErrUserNotFound,
		constant.ErrActionNotFound,
		constant.ErrReportSolutionProcessNotFound,
		constant.ErrAdminNotFound,
		constant.ErrSuperAdminCannotBeDeleted,
	}

	if contains(badRequestErrors, err) {
		return http.StatusBadRequest
	} else if err == constant.ErrUnauthorized {
		return http.StatusUnauthorized
	} else {
		return http.StatusInternalServerError
	}
}

func contains(slice []error, item error) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
