package services

import (
	"github.com/danyaobertan/validcard/internal/api/domain"
	"github.com/danyaobertan/validcard/pkg/errorops"
)

type (
	ValidCard interface {
		IsValidCardInfo(domain.Card) (bool, *errorops.Error)
	}
)
