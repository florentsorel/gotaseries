package gotaseries

import "time"

// Bool returns a pointer value for the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

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

// Order returns a pointer value for the order value passed in.
func Order(v order) *order {
	return &v
}

// Time returns a pointer value for the time.Time value passed in.
func Time(v time.Time) *time.Time {
	return &v
}
