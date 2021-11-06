package errors

import (
	"errors"
)

var (
	ErrorFontReadFailed              = errors.New("Failed to read font.")
	ErrorFontParseFailed             = errors.New("Failed to parse TrueType font.")
	ErrorFontSizeNegative            = errors.New("FontSize should be positive number.")
	ErrorInvalidColorFormat          = errors.New("Invalid color format.")
	ErrorBackgroundImageReadFailed   = errors.New("Failed to read background image.")
	ErrorBackgroundImageDecodeFailed = errors.New("Failed to decode background image.")
	ErrorDrawFailed                  = errors.New("Failed to draw image.")
	ErrorOutputFileCreateFailed      = errors.New("Failed to create output file.")
	ErrorOutputFailed                = errors.New("Failed to output.")
)
