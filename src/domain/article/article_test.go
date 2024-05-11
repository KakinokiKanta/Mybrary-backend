package article

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/oklog/ulid/v2"
)

func TestNewArticle(t *testing.T) {
	userID := ulid.Make().String()
	type args struct {
		userID string
		url string
		title string
		description string
		tags []ArticleTag
	}
	tests := []struct {
		testName string
		args args
		expected Article
		expectedErr bool
	}{
		{
			testName: "正常系",
			args: args{
				userID: userID,
				url: "https://example.dev",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{ArticleTag{"Go"}, ArticleTag{"Gin"}},
			},
			expected: Article{
				userID: userID,
				url: "https://example.dev",
				title: "テスト記事",
				description: "テスト記事の詳細",
				tags: []ArticleTag{ArticleTag{"Go"}, ArticleTag{"Gin"}},
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := NewArticle(tt.args.userID, tt.args.url, tt.args.title, tt.args.description, tt.args.tags)
			if (err != nil) != tt.expectedErr {
				t.Errorf("NewArticle() error = %v, ExpectedErr %v", err, tt.expectedErr)
				return
			}
			diff := cmp.Diff(
				got, tt.expected,
				cmp.AllowUnexported(Article{}),
				cmpopts.IgnoreFields(Article{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewArticle() = %v, expected %v. error is %s", got, tt.expected, err)
			}
		})
	}
}