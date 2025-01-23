// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: amp_invoices.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const fetchAMPSubInvoiceHTLCs = `-- name: FetchAMPSubInvoiceHTLCs :many
SELECT 
    amp.set_id, amp.root_share, amp.child_index, amp.hash, amp.preimage, 
    invoice_htlcs.id, invoice_htlcs.chan_id, invoice_htlcs.htlc_id, invoice_htlcs.amount_msat, invoice_htlcs.total_mpp_msat, invoice_htlcs.accept_height, invoice_htlcs.accept_time, invoice_htlcs.expiry_height, invoice_htlcs.state, invoice_htlcs.resolve_time, invoice_htlcs.invoice_id
FROM amp_sub_invoice_htlcs amp
INNER JOIN invoice_htlcs ON amp.htlc_id = invoice_htlcs.id
WHERE amp.invoice_id = $1
AND (
    set_id = $2 OR 
    $2 IS NULL
)
`

type FetchAMPSubInvoiceHTLCsParams struct {
	InvoiceID int64
	SetID     []byte
}

type FetchAMPSubInvoiceHTLCsRow struct {
	SetID        []byte
	RootShare    []byte
	ChildIndex   int64
	Hash         []byte
	Preimage     []byte
	ID           int64
	ChanID       string
	HtlcID       int64
	AmountMsat   int64
	TotalMppMsat sql.NullInt64
	AcceptHeight int32
	AcceptTime   time.Time
	ExpiryHeight int32
	State        int16
	ResolveTime  sql.NullTime
	InvoiceID    int64
}

func (q *Queries) FetchAMPSubInvoiceHTLCs(ctx context.Context, arg FetchAMPSubInvoiceHTLCsParams) ([]FetchAMPSubInvoiceHTLCsRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchAMPSubInvoiceHTLCs, arg.InvoiceID, arg.SetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchAMPSubInvoiceHTLCsRow
	for rows.Next() {
		var i FetchAMPSubInvoiceHTLCsRow
		if err := rows.Scan(
			&i.SetID,
			&i.RootShare,
			&i.ChildIndex,
			&i.Hash,
			&i.Preimage,
			&i.ID,
			&i.ChanID,
			&i.HtlcID,
			&i.AmountMsat,
			&i.TotalMppMsat,
			&i.AcceptHeight,
			&i.AcceptTime,
			&i.ExpiryHeight,
			&i.State,
			&i.ResolveTime,
			&i.InvoiceID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchAMPSubInvoices = `-- name: FetchAMPSubInvoices :many
SELECT set_id, state, created_at, settled_at, settle_index, invoice_id
FROM amp_sub_invoices
WHERE invoice_id = $1 
AND (
    set_id = $2 OR 
    $2 IS NULL
)
`

type FetchAMPSubInvoicesParams struct {
	InvoiceID int64
	SetID     []byte
}

func (q *Queries) FetchAMPSubInvoices(ctx context.Context, arg FetchAMPSubInvoicesParams) ([]AmpSubInvoice, error) {
	rows, err := q.db.QueryContext(ctx, fetchAMPSubInvoices, arg.InvoiceID, arg.SetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AmpSubInvoice
	for rows.Next() {
		var i AmpSubInvoice
		if err := rows.Scan(
			&i.SetID,
			&i.State,
			&i.CreatedAt,
			&i.SettledAt,
			&i.SettleIndex,
			&i.InvoiceID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchSettledAMPSubInvoices = `-- name: FetchSettledAMPSubInvoices :many
SELECT 
    a.set_id, 
    a.settle_index as amp_settle_index, 
    a.settled_at as amp_settled_at,
    i.id, i.hash, i.preimage, i.settle_index, i.settled_at, i.memo, i.amount_msat, i.cltv_delta, i.expiry, i.payment_addr, i.payment_request, i.payment_request_hash, i.state, i.amount_paid_msat, i.is_amp, i.is_hodl, i.is_keysend, i.created_at
FROM amp_sub_invoices a
INNER JOIN invoices i ON a.invoice_id = i.id
WHERE (
    a.settle_index >= $1 OR
    $1 IS NULL
) AND (
    a.settle_index <= $2 OR
    $2 IS NULL
)
`

type FetchSettledAMPSubInvoicesParams struct {
	SettleIndexGet sql.NullInt64
	SettleIndexLet sql.NullInt64
}

type FetchSettledAMPSubInvoicesRow struct {
	SetID              []byte
	AmpSettleIndex     sql.NullInt64
	AmpSettledAt       sql.NullTime
	ID                 int64
	Hash               []byte
	Preimage           []byte
	SettleIndex        sql.NullInt64
	SettledAt          sql.NullTime
	Memo               sql.NullString
	AmountMsat         int64
	CltvDelta          sql.NullInt32
	Expiry             int32
	PaymentAddr        []byte
	PaymentRequest     sql.NullString
	PaymentRequestHash []byte
	State              int16
	AmountPaidMsat     int64
	IsAmp              bool
	IsHodl             bool
	IsKeysend          bool
	CreatedAt          time.Time
}

func (q *Queries) FetchSettledAMPSubInvoices(ctx context.Context, arg FetchSettledAMPSubInvoicesParams) ([]FetchSettledAMPSubInvoicesRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchSettledAMPSubInvoices, arg.SettleIndexGet, arg.SettleIndexLet)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchSettledAMPSubInvoicesRow
	for rows.Next() {
		var i FetchSettledAMPSubInvoicesRow
		if err := rows.Scan(
			&i.SetID,
			&i.AmpSettleIndex,
			&i.AmpSettledAt,
			&i.ID,
			&i.Hash,
			&i.Preimage,
			&i.SettleIndex,
			&i.SettledAt,
			&i.Memo,
			&i.AmountMsat,
			&i.CltvDelta,
			&i.Expiry,
			&i.PaymentAddr,
			&i.PaymentRequest,
			&i.PaymentRequestHash,
			&i.State,
			&i.AmountPaidMsat,
			&i.IsAmp,
			&i.IsHodl,
			&i.IsKeysend,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAMPInvoiceID = `-- name: GetAMPInvoiceID :one
SELECT invoice_id FROM amp_sub_invoices WHERE set_id = $1
`

func (q *Queries) GetAMPInvoiceID(ctx context.Context, setID []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAMPInvoiceID, setID)
	var invoice_id int64
	err := row.Scan(&invoice_id)
	return invoice_id, err
}

const insertAMPSubInvoice = `-- name: InsertAMPSubInvoice :exec
INSERT INTO amp_sub_invoices (
    set_id, state, created_at, settled_at, settle_index, invoice_id
) VALUES (
    $1, $2, $3, $4, $5, $6
)
`

type InsertAMPSubInvoiceParams struct {
	SetID       []byte
	State       int16
	CreatedAt   time.Time
	SettledAt   sql.NullTime
	SettleIndex sql.NullInt64
	InvoiceID   int64
}

func (q *Queries) InsertAMPSubInvoice(ctx context.Context, arg InsertAMPSubInvoiceParams) error {
	_, err := q.db.ExecContext(ctx, insertAMPSubInvoice,
		arg.SetID,
		arg.State,
		arg.CreatedAt,
		arg.SettledAt,
		arg.SettleIndex,
		arg.InvoiceID,
	)
	return err
}

const insertAMPSubInvoiceHTLC = `-- name: InsertAMPSubInvoiceHTLC :exec
INSERT INTO amp_sub_invoice_htlcs (
    invoice_id, set_id, htlc_id, root_share, child_index, hash, preimage
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
`

type InsertAMPSubInvoiceHTLCParams struct {
	InvoiceID  int64
	SetID      []byte
	HtlcID     int64
	RootShare  []byte
	ChildIndex int64
	Hash       []byte
	Preimage   []byte
}

func (q *Queries) InsertAMPSubInvoiceHTLC(ctx context.Context, arg InsertAMPSubInvoiceHTLCParams) error {
	_, err := q.db.ExecContext(ctx, insertAMPSubInvoiceHTLC,
		arg.InvoiceID,
		arg.SetID,
		arg.HtlcID,
		arg.RootShare,
		arg.ChildIndex,
		arg.Hash,
		arg.Preimage,
	)
	return err
}

const updateAMPSubInvoiceHTLCPreimage = `-- name: UpdateAMPSubInvoiceHTLCPreimage :execresult
UPDATE amp_sub_invoice_htlcs AS a
SET preimage = $5
WHERE a.invoice_id = $1 AND a.set_id = $2 AND a.htlc_id = (
    SELECT id FROM invoice_htlcs AS i WHERE i.chan_id = $3 AND i.htlc_id = $4
)
`

type UpdateAMPSubInvoiceHTLCPreimageParams struct {
	InvoiceID int64
	SetID     []byte
	ChanID    string
	HtlcID    int64
	Preimage  []byte
}

func (q *Queries) UpdateAMPSubInvoiceHTLCPreimage(ctx context.Context, arg UpdateAMPSubInvoiceHTLCPreimageParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateAMPSubInvoiceHTLCPreimage,
		arg.InvoiceID,
		arg.SetID,
		arg.ChanID,
		arg.HtlcID,
		arg.Preimage,
	)
}

const updateAMPSubInvoiceState = `-- name: UpdateAMPSubInvoiceState :exec
UPDATE amp_sub_invoices
SET state = $2, 
    settle_index = COALESCE(settle_index, $3),
    settled_at = COALESCE(settled_at, $4)
WHERE set_id = $1
`

type UpdateAMPSubInvoiceStateParams struct {
	SetID       []byte
	State       int16
	SettleIndex sql.NullInt64
	SettledAt   sql.NullTime
}

func (q *Queries) UpdateAMPSubInvoiceState(ctx context.Context, arg UpdateAMPSubInvoiceStateParams) error {
	_, err := q.db.ExecContext(ctx, updateAMPSubInvoiceState,
		arg.SetID,
		arg.State,
		arg.SettleIndex,
		arg.SettledAt,
	)
	return err
}

const upsertAMPSubInvoice = `-- name: UpsertAMPSubInvoice :execresult
INSERT INTO amp_sub_invoices (
    set_id, state, created_at, invoice_id
) VALUES (
    $1, $2, $3, $4
) ON CONFLICT (set_id, invoice_id) DO NOTHING
`

type UpsertAMPSubInvoiceParams struct {
	SetID     []byte
	State     int16
	CreatedAt time.Time
	InvoiceID int64
}

func (q *Queries) UpsertAMPSubInvoice(ctx context.Context, arg UpsertAMPSubInvoiceParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertAMPSubInvoice,
		arg.SetID,
		arg.State,
		arg.CreatedAt,
		arg.InvoiceID,
	)
}
