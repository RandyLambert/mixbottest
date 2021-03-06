package main

import (
	"context"
	"log"

	sdk "github.com/fox-one/mixin-sdk"
	"github.com/gofrs/uuid"
)

func doTransaction(ctx context.Context, user *sdk.User, assetID, opponentKey, amount, memo, pin string) {
	//做交易到mixin主网
	snapshot, err := user.Transaction(ctx, &sdk.TransferInput{
		TraceID:     uuid.Must(uuid.NewV4()).String(),
		AssetID:     assetID,
		OpponentKey: opponentKey,
		Amount:      amount,
		Memo:        memo,
	}, pin)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("do transfer", snapshot)
}

func doTransfer(ctx context.Context, user *sdk.User, assetID, opponentID, amount, memo, pin string) *sdk.Snapshot {
	//opponentid是转给谁
	// Transfer transfer to account
	//	asset_id, opponent_id, amount, traceID, memo
	// 把该user的钱转账到该账户返回快照
	snapshot, err := user.Transfer(ctx, &sdk.TransferInput{
		TraceID:    uuid.Must(uuid.NewV4()).String(),
		AssetID:    assetID,
		OpponentID: opponentID,
		Amount:     amount,
		Memo:       memo,
	}, pin)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("do transfer", snapshot)
	return snapshot
}
//提现函数
func doWithdraw(ctx context.Context, user *sdk.User, assetID, publicKey, amount, memo, pin string) *sdk.Snapshot {
	//创建提款地址
	addrID := doCreateAddress(ctx, user, assetID, publicKey, "Test Withdraw", pin)

	snapshot, err := user.Withdraw(ctx, &sdk.TransferInput{
		TraceID:   uuid.Must(uuid.NewV4()).String(),
		AssetID:   assetID,
		AddressID: addrID,
		Amount:    amount,
		Memo:      memo,
	}, pin)
	if err != nil {
		log.Panicln(err)
	}
	printJSON("do withdraw", snapshot)
	//删除提款地址
	doDeleteAddress(ctx, user, addrID, pin)
	return snapshot
}
