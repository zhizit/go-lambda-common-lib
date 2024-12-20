package utils

import "database/sql"

// sql.Null 型に関連した変換を行います

// *bool を sql.NullBool に変換
func ToNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false} // nil の場合、NULLとして扱う
	}
	return sql.NullBool{
		Bool:  *b,
		Valid: true, // bがnilでない場合は有効な値として設定
	}
}

// *int64 を sql.NullInt64 に変換
func ToNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{
		Int64: *i,
		Valid: true,
	}
}

// *float64 を sql.NullFloat64 に変換
func ToNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{
		Float64: *f,
		Valid:   true,
	}
}

// *string を sql.NullString に変換
func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}
