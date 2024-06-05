package gotaseries

// String returns a pointer value for the string value passed in.
func String(v string) *string {
	return &v
}

// Int returns a pointer value for the int value passed in.
func Int(v int) *int {
	return &v
}

// Locale returns a pointer value for the locale value passed in.
func Locale(v locale) *locale {
	return &v
}
