package app

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHandler(t *testing.T) {
	te := newTestEnv(t)
	type want struct {
		code int
		body *string
	}
	tests := []struct {
		name  string
		value string
		want  want
	}{
		{
			name:  "without parameter",
			value: "",
			want:  want{code: 400},
		},
		{
			name:  "with wrong parameter",
			value: "999",
			want:  want{code: 400},
		},
		{
			name:  "with parameter",
			value: "8080bf6b83e4-916f-4312-84af-ce686046cf0e",
			want:  want{code: 307},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.value)

			te.urlsArchive.Add("8080bf6b83e4-916f-4312-84af-ce686046cf0e", "ya.ru")

			h := te.s.Get()
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.want.code, rec.Code)
			}
		})
	}
}
