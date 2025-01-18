package auth

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestVerifyAuthToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "Error (no rows)",
			args: args{
				ctx:   context.Background(),
				token: uuid.MustParse("8a43e239-83d1-4d25-86fb-b00e20569246"), // Random UUID
			},
			wantErr: true,
		},
		{
			name: "Valid token",
			args: args{
				ctx:   context.Background(),
				token: uuid.MustParse("fd0d33c9-05b2-4fae-9b19-9e41414cee27"),
			},
			wantErr: false,
		},
		{
			name: "Expired token",
			args: args{
				ctx:   context.Background(),
				token: uuid.MustParse("5d82b60a-7bc9-4fa8-be79-d7bb70435dc5"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, e := VerifyAuthToken(tt.args.ctx, tt.args.token)
			if (e != nil) != tt.wantErr {
				t.Errorf("VerifyAuthToken() error = %v, wantErr %v", e, tt.wantErr)
			}
		})
	}
}
