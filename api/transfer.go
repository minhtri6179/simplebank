package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhtri67/simplebank/db/sqlc"
	"github.com/minhtri67/simplebank/token"
)

type transferRequest struct {
	FromAccountID int64  `binding:"required,min=1"`
	ToAccountID   int64  `binding:"required,min=1"`
	Amount        int64  `binding:"required,gt=0"`
	Currency      string `binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}
	authuPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authuPayload.Username {
		err := fmt.Errorf("account %d does not belong to user %s", fromAccount.ID, authuPayload.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	_, valid = server.validAccount(ctx, req.ToAccountID, req.Currency)
	if !valid {
		return
	}

	arg := sqlc.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (sqlc.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account %d is not a valid currency: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}
	return account, true
}
