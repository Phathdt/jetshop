package channel_biz

import (
	"context"

	"jetshop/service-context/core"
	"jetshop/services/channel_service/internal/modules/channel/channel_model"
)

type listHermesChannelCredentialRepo interface {
	ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error)
}

type listHermesChannelCredentialBiz struct {
	repo listHermesChannelCredentialRepo
}

func NewListHermesChannelCredentialBiz(repo listHermesChannelCredentialRepo) *listHermesChannelCredentialBiz {
	return &listHermesChannelCredentialBiz{repo: repo}
}

func (b *listHermesChannelCredentialBiz) Response(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error) {
	credentials, err := b.repo.ListChannelCredentials(ctx, cond)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(channel_model.ErrCannotListHermesChannelCredential.Error()).
			WithDebug(err.Error())
	}

	return credentials, nil
}
