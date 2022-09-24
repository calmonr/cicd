package greeting

import (
	"context"

	"github.com/calmonr/cicd/pkg/greeting"
	greetingv1 "github.com/calmonr/cicd/pkg/proto/greeting/v1"
)

type Service struct {
	greetingv1.UnimplementedGreetingServiceServer
}

func (s *Service) Hi(ctx context.Context, req *greetingv1.HiRequest) (*greetingv1.HiResponse, error) {
	return &greetingv1.HiResponse{
		Message: greeting.Hi(req.Name),
	}, nil
}
