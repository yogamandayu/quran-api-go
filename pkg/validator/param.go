package validator

// Validate checks that id is correct.

import (
	"math"
	"quran-api-go/internal/domain"
	"strconv"
)

// Returns domain.ErrInvalidIDParam for any other value.
//
// Usage:
//
//	lang, err := validator.ValidateIDParam(c.Query("id"))
//	if err != nil {
//	    response.BadRequest(c, "invalid param")
//	    return
//	}
func ValidateIDParam(id string) (string, error) {
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return "", domain.ErrInvalidIDParam
	}

	if math.IsNaN(float64(convertedId)) {
		return "", domain.ErrInvalidIDParam
	}

	if convertedId < 1 {
		return "", domain.ErrInvalidIDParam
	}

	return id, nil
}
