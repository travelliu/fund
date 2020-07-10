// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"github.com/sirupsen/logrus"
	_fund "github.com/travelliu/fund/services/fund"
	"github.com/travelliu/fund/utils/logs"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logs.NewLogger()
}

type fund struct {
	fundRepo _fund.Repository
}

// NewFundUc New Fund UseCase
func NewFundUc(fundRepo _fund.Repository) _fund.UseCase {
	return &fund{
		fundRepo: fundRepo,
	}
}
