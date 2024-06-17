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

// Format returns a pointer value for the FormatType value passed in.
func Format(v FormatType) *FormatType {
	return &v
}

// Locale returns a pointer value for the LocaleType value passed in.
func Locale(v LocaleType) *LocaleType {
	return &v
}

// Order returns a pointer value for the OrderType value passed in.
func Order(v OrderType) *OrderType {
	return &v
}

// OrderDate returns a pointer value for the OrderDateType value passed in.
func OrderDate(v OrderDateType) *OrderDateType {
	return &v
}

// Time returns a pointer value for the time.Time value passed in.
func Time(v time.Time) *time.Time {
	return &v
}

func OrderFavorite(v OrderFavoriteType) *OrderFavoriteType {
	return &v
}

func OrderShowMember(v OrderShowMemberType) *OrderShowMemberType {
	return &v
}

func StatusFavorite(v StatusFavoriteType) *StatusFavoriteType {
	return &v
}

func StatusShowMember(v StatusShowMemberType) *StatusShowMemberType {
	return &v
}
