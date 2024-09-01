package repository

import (
	"database/sql"
	"testing"

	"github.com/KakinokiKanta/Mybrary-backend/domain"
	"github.com/google/go-cmp/cmp"
)

func TestCreateArticleTagRepository(t *testing.T) {
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
			args: *succeedDomain,
			expected: *succeedDomain,
			expectedErr: false,
		},
		{
			testName: "Failure: create article tag (this article tag exists in DB)",
			db: testDB,
			args: *succeedDomain,
			expected: domain.ArticleTag{},
			expectedErr: true,
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
				cmp.AllowUnexported(domain.ArticleTag{}),
			)
			if diff != "" {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.testName, result, tt.expected)
			}
		})
	}
}

func TestFindArticleTagRepository(t *testing.T) {
	tests := []struct {
		testName string
		db *sql.DB
		args string
		expected domain.ArticleTag
		expectedErr bool
	}{
		{
			testName: "Successfully find article tag",
			db: testDB,
			args: succeedTagName,
			expected: *succeedDomain,
			expectedErr: false,
		},
		{
			testName: "Failure: find article tag (this article tag does not exists in DB)",
			db: testDB,
			args: "Article Tag",
			expected: domain.ArticleTag{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			repo := NewArticleTagRepository(testDB)

			result, err := repo.FindByName(tt.args)
			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.testName, err, tt.expectedErr)
				t.Errorf("[TestCase '%s'] %s | %v | %v", tt.testName, tt.args, *succeedDomain, result)
				return
			}
			diff := cmp.Diff(
				result, tt.expected,
				cmp.AllowUnexported(domain.ArticleTag{}),
			)
			if diff != "" {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.testName, result, tt.expected)
			}
		})
	}
}

func TestUpdateNumArticleTagRepository(t *testing.T) {
	// TODO: 中身書き換えてないで
	tests := []struct {
		testName string
		db *sql.DB
		args string
		expected domain.ArticleTag
		expectedErr bool
	}{
		{
			testName: "Successfully find article tag",
			db: testDB,
			args: succeedTagName,
			expected: *succeedDomain,
			expectedErr: false,
		},
		{
			testName: "Failure: find article tag (this article tag does not exists in DB)",
			db: testDB,
			args: "Article Tag",
			expected: domain.ArticleTag{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			repo := NewArticleTagRepository(testDB)

			result, err := repo.FindByName(tt.args)
			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.testName, err, tt.expectedErr)
				t.Errorf("[TestCase '%s'] %s | %v | %v", tt.testName, tt.args, *succeedDomain, result)
				return
			}
			diff := cmp.Diff(
				result, tt.expected,
				cmp.AllowUnexported(domain.ArticleTag{}),
			)
			if diff != "" {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.testName, result, tt.expected)
			}
		})
	}
}
