package PgToStruct

import (
	"testing"
)

func TestGormTableName(t *testing.T) {
	var testdata = []struct {
		in  string
		out string
	}{
		{"events", "Event"},
		{"post_comments", "PostComment"},
	}

	for _, td := range testdata {
		s := gormTableName(td.in)

		if s != td.out {
			t.Fatalf("Table name does not match. expect: %s, actual: %s", td.out, s)
		}
	}
}

func TestGormColumnName(t *testing.T) {
	var testdata = []struct {
		in  string
		out string
	}{
		{"description", "Description"},
		{"user_id", "UserID"},
		{"facebook_uid", "FacebookUID"},
		{"candidacy", "Candidacy"},
		{"video_id", "VideoID"},
		{"image_url", "ImageURL"},
		{"curl_name", "CurlName"},
	}

	for _, td := range testdata {
		s := gormColumnName(td.in)

		if s != td.out {
			t.Fatalf("Field name does not match. expect: %s, actual: %s", td.out, s)
		}
	}
}
