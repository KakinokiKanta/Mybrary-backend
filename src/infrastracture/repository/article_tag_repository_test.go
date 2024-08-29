package repository

import (
	"database/sql"
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	"github.com/google/go-cmp/cmp"
)

func TestCreateArticleTagRepository(t *testing.T) {
	articleTagDomain, _ := domain.NewArticleTag("abcdefg1234AABBCCDD", "New Article Tag")

	tests := []struct {
		testName string
		db *sql.DB
		args domain.ArticleTag
		expected domain.ArticleTag
		expectedErr bool
	}{
		{
			testName: "Successfully create article tag",
			db: testDB,
			args: *articleTagDomain,
			expected: *articleTagDomain,
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			repo := NewArticleTagRepository(testDB)

			result, err := repo.Create(tt.args)
			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.testName, err, tt.expectedErr)
				return
			}
			diff := cmp.Diff(
				result, tt.expected,
				cmp.AllowUnexported(domain.User{}),
			)
			if diff != "" {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.testName, result, tt.expected)
			}
		})
	}
}
