// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package usecase

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	_fund "github.com/travelliu/fund/services/fund"
	_fundMod "github.com/travelliu/fund/services/fund/models"
	"github.com/travelliu/fund/services/fund/utils"
	"github.com/travelliu/fund/utils/trace"
	"sync"
)

var startOnce sync.Once

type cronLogger struct {
}

// NewCronLogger New Cron Logger
func NewCronLogger() cron.Logger {
	return &cronLogger{}
}

// Info Info
func (l *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	l.log(nil, msg, keysAndValues...)
}

// Error Error
func (l *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.log(err, msg, keysAndValues...)
}
func (l *cronLogger) log(err error, msg string, keysAndValues ...interface{}) {
	loggerFields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i = i + 2 {
		key := keysAndValues[i].(string)
		loggerFields[key] = keysAndValues[i+1]
	}
	
	logger.WithFields(
		loggerFields).Info(msg)
}
func NewCron(ctx context.Context, fundUc _fund.UseCase, spec string, workerNum int) {
	if spec == "" {
		spec = "0 0/1 * * * ?"
	}
	if workerNum == 0 {
		workerNum = 5
	}
	var options []cron.Option
	options = append(options, cron.WithSeconds())
	options = append(options, cron.WithLogger(NewCronLogger()))
	c := cron.New(options...)
	
	startOnce.Do(func() {
		_, err := c.AddFunc(spec, func() {
			fundUc.FundSync(ctx, workerNum)
		})
		if err != nil {
			logger.Fatalf("add FundSync job failed %s", err)
		}
		c.Start()
		
	})
}

func (f *fund) FundSync(ctx context.Context, parallel int) error {
	funds, err := f.fundRepo.QueryAllFund(ctx)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Errorf("the QueryAllFund error %s", err)
		return err
	}
	reqChan := make(chan *_fundMod.Fund, parallel)
	go func() {
		for i := range funds {
			reqChan <- funds[i]
		}
		close(reqChan)
	}()
	wg := sync.WaitGroup{}
	
	wg.Add(parallel)
	for i := 0; i < parallel; i++ {
		go func() {
			f.fundSync(ctx, reqChan)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
func (f *fund) fundSync(ctx context.Context, reqChan chan *_fundMod.Fund) {
	for {
		select {
		case fund, ok := <-reqChan:
			if !ok {
				return
			}
			if err := f.doFundSync(ctx, fund); err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
func (f *fund) doFundSync(ctx context.Context, fund *_fundMod.Fund) error {
	logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Debugf("the begin sync %s", fund.Code)
	newFund, err := utils.GetFundInfo(fund.Code)
	if err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Debugf("the will sync %s GetFundInfo error %s", fund.Code, err)
		return err
	}
	fund.FundBase = newFund.FundBase
	if err := f.fundRepo.UpdateFund(ctx, fund); err != nil {
		logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Debugf("the UpdateFund %s error %s", fund.Code, err)
		return err
	}
	logger.WithField(string(trace.ContextKeyReqID), trace.GetReqID(ctx)).Debugf("the end sync %s", fund.Code)
	return nil
}
