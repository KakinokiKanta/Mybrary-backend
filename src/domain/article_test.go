package domain

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/oklog/ulid/v2"
)

func TestNewArticle(t *testing.T) {
	userID := UserID(ulid.Make().String())
	type args struct {
		userID UserID
		url string
		title string
		description string
		tags []ArticleTag
	}
	tests := []struct {
		testName string
		args args
		expected *Article
		expectedErr bool
	}{
		{
			testName: "正常系",
			args: args{
				userID: userID,
				url: "https://example.dev",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"}},
			},
			expected: &Article{
				userID: userID,
				url: "https://example.dev",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"}},
			},
			expectedErr: false,
		},
		{
			testName: "異常系: ユーザIDが不正",
			args: args{
				userID: "userID",
				url: "https://example.dev",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"}},
			},
			expected: nil,
			expectedErr: true,
		},
		{
			testName: "異常系: URLなし",
			args: args{
				userID: userID,
				url: "",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"}},
			},
			expected: nil,
			expectedErr: true,
		},
		{
			testName: "異常系: 記事タイトルなし",
			args: args{
				userID: userID,
				url: "https://example.dev",
				title: "",
				description: "テスト記事の詳細",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"}},
			},
			expected: nil,
			expectedErr: true,
		},
		{
			testName: "異常系: 記事の詳細なし",
			args: args{
				userID: userID,
				url: "https://example.dev",
				title: "テスト記事",
				description: "",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"}},
			},
			expected: nil,
			expectedErr: true,
		},
		{
			testName: "異常系: 記事タグが6個",
			args: args{
				userID: userID,
				url: "https://example.dev",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{{tagName: "GO"}, {tagName: "Gin"},{tagName: "Echo"}, {tagName: "Beego"},{tagName: "Gorilla"}, {tagName: "Revel"}},
			},
			expected: nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, err := NewArticle(tt.args.userID, tt.args.url, tt.args.title, tt.args.description, tt.args.tags)
			if (err != nil) != tt.expectedErr {
				t.Errorf("NewArticle() error = %v, ExpectedErr %v", err, tt.expectedErr)
				return
			}
			diff := cmp.Diff(
				result, tt.expected,
				cmp.AllowUnexported(Article{}, ArticleTag{}),
				cmpopts.IgnoreFields(Article{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewArticle() = %v, expected %v, error is %s, diff is %v", result, tt.expected, err, diff)
			}
		})
	}
}
