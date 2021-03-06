package baseabci

import (
	ctx "github.com/QOSGroup/qbase/context"
	"github.com/QOSGroup/qbase/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// initialize application state at genesis
type InitChainHandler func(ctx ctx.Context, req abci.RequestInitChain) abci.ResponseInitChain

// run code before the transactions in a block
type BeginBlockHandler func(ctx ctx.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock

// run code after the transactions in a block and return updates to the validator set
type EndBlockHandler func(ctx ctx.Context, req abci.RequestEndBlock) abci.ResponseEndBlock

// -------------------------------------------------------------------------------------------------------------------------

//TxQcpResultHandler qcpTx result 回调函数，在TxQcpResult.Exec中调用
//Important!: txQcpResult 类型为 *txs.QcpTxResult
type TxQcpResultHandler func(ctx ctx.Context, txQcpResult interface{}) types.Result
