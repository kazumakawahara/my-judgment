package testutil

import "time"

func ToStringPtr(s string) *string {
	return &s
}

func ToIntPtr(i int) *int {
	return &i
}

func ToTimePtr(t time.Time) *time.Time {
	return &t
}
