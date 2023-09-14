// Code generated by go generate; DO NOT EDIT.
package db

import (
	stdsql "database/sql"
	"time"

	"github.com/sorintlab/errors"
	sq "github.com/huandu/go-sqlbuilder"

	"agola.io/agola/internal/sqlg/sql"

	types "agola.io/agola/services/notification/types"
)
var (
	runWebhookInsertPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inPayload []byte) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhook").Cols("id", "revision", "creation_time", "update_time", "payload").Values(inId, inRevision, inCreationTime, inUpdateTime, inPayload)
	}
	runWebhookUpdatePostgres = func(curRevision uint64, inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inPayload []byte) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runwebhook").Set(ub.Assign("id", inId), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("payload", inPayload)).Where(ub.E("id", inId), ub.E("revision", curRevision))
	}

	runWebhookInsertRawPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inPayload []byte) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhook").Cols("id", "revision", "creation_time", "update_time", "payload").SQL("OVERRIDING SYSTEM VALUE").Values(inId, inRevision, inCreationTime, inUpdateTime, inPayload)
	}
)

func (d *DB) insertRunWebhookPostgres(tx *sql.Tx, runwebhook *types.RunWebhook) error {
	q := runWebhookInsertPostgres(runwebhook.ID, runwebhook.Revision, runwebhook.CreationTime, runwebhook.UpdateTime, runwebhook.Payload)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhook")
	}

	return nil
}

func (d *DB) updateRunWebhookPostgres(tx *sql.Tx, curRevision uint64, runwebhook *types.RunWebhook) (stdsql.Result, error) {
	q := runWebhookUpdatePostgres(curRevision, runwebhook.ID, runwebhook.Revision, runwebhook.CreationTime, runwebhook.UpdateTime, runwebhook.Payload)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runWebhook")
	}

	return res, nil
}

func (d *DB) insertRawRunWebhookPostgres(tx *sql.Tx, runwebhook *types.RunWebhook) error {
	q := runWebhookInsertRawPostgres(runwebhook.ID, runwebhook.Revision, runwebhook.CreationTime, runwebhook.UpdateTime, runwebhook.Payload)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhook")
	}

	return nil
}
var (
	runWebhookDeliveryInsertPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inRunWebhookID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time, inStatusCode int) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhookdelivery").Cols("id", "revision", "creation_time", "update_time", "run_webhook_id", "delivery_status", "delivered_at", "status_code").Values(inId, inRevision, inCreationTime, inUpdateTime, inRunWebhookID, inDeliveryStatus, inDeliveredAt, inStatusCode)
	}
	runWebhookDeliveryUpdatePostgres = func(curRevision uint64, inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inRunWebhookID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time, inStatusCode int) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runwebhookdelivery").Set(ub.Assign("id", inId), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("run_webhook_id", inRunWebhookID), ub.Assign("delivery_status", inDeliveryStatus), ub.Assign("delivered_at", inDeliveredAt), ub.Assign("status_code", inStatusCode)).Where(ub.E("id", inId), ub.E("revision", curRevision))
	}

	runWebhookDeliveryInsertRawPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inRunWebhookID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time, inStatusCode int) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runwebhookdelivery").Cols("id", "revision", "creation_time", "update_time", "sequence", "run_webhook_id", "delivery_status", "delivered_at", "status_code").SQL("OVERRIDING SYSTEM VALUE").Values(inId, inRevision, inCreationTime, inUpdateTime, inSequence, inRunWebhookID, inDeliveryStatus, inDeliveredAt, inStatusCode)
	}
)

func (d *DB) insertRunWebhookDeliveryPostgres(tx *sql.Tx, runwebhookdelivery *types.RunWebhookDelivery) error {
	q := runWebhookDeliveryInsertPostgres(runwebhookdelivery.ID, runwebhookdelivery.Revision, runwebhookdelivery.CreationTime, runwebhookdelivery.UpdateTime, runwebhookdelivery.RunWebhookID, runwebhookdelivery.DeliveryStatus, runwebhookdelivery.DeliveredAt, runwebhookdelivery.StatusCode)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhookDelivery")
	}

	return nil
}

func (d *DB) updateRunWebhookDeliveryPostgres(tx *sql.Tx, curRevision uint64, runwebhookdelivery *types.RunWebhookDelivery) (stdsql.Result, error) {
	q := runWebhookDeliveryUpdatePostgres(curRevision, runwebhookdelivery.ID, runwebhookdelivery.Revision, runwebhookdelivery.CreationTime, runwebhookdelivery.UpdateTime, runwebhookdelivery.RunWebhookID, runwebhookdelivery.DeliveryStatus, runwebhookdelivery.DeliveredAt, runwebhookdelivery.StatusCode)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runWebhookDelivery")
	}

	return res, nil
}

func (d *DB) insertRawRunWebhookDeliveryPostgres(tx *sql.Tx, runwebhookdelivery *types.RunWebhookDelivery) error {
	q := runWebhookDeliveryInsertRawPostgres(runwebhookdelivery.ID, runwebhookdelivery.Revision, runwebhookdelivery.CreationTime, runwebhookdelivery.UpdateTime, runwebhookdelivery.Sequence, runwebhookdelivery.RunWebhookID, runwebhookdelivery.DeliveryStatus, runwebhookdelivery.DeliveredAt, runwebhookdelivery.StatusCode)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runWebhookDelivery")
	}

	return nil
}
var (
	lastRunEventSequenceInsertPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inValue uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("lastruneventsequence").Cols("id", "revision", "creation_time", "update_time", "value").Values(inId, inRevision, inCreationTime, inUpdateTime, inValue)
	}
	lastRunEventSequenceUpdatePostgres = func(curRevision uint64, inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inValue uint64) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("lastruneventsequence").Set(ub.Assign("id", inId), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("value", inValue)).Where(ub.E("id", inId), ub.E("revision", curRevision))
	}

	lastRunEventSequenceInsertRawPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inValue uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("lastruneventsequence").Cols("id", "revision", "creation_time", "update_time", "value").SQL("OVERRIDING SYSTEM VALUE").Values(inId, inRevision, inCreationTime, inUpdateTime, inValue)
	}
)

func (d *DB) insertLastRunEventSequencePostgres(tx *sql.Tx, lastruneventsequence *types.LastRunEventSequence) error {
	q := lastRunEventSequenceInsertPostgres(lastruneventsequence.ID, lastruneventsequence.Revision, lastruneventsequence.CreationTime, lastruneventsequence.UpdateTime, lastruneventsequence.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert lastRunEventSequence")
	}

	return nil
}

func (d *DB) updateLastRunEventSequencePostgres(tx *sql.Tx, curRevision uint64, lastruneventsequence *types.LastRunEventSequence) (stdsql.Result, error) {
	q := lastRunEventSequenceUpdatePostgres(curRevision, lastruneventsequence.ID, lastruneventsequence.Revision, lastruneventsequence.CreationTime, lastruneventsequence.UpdateTime, lastruneventsequence.Value)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update lastRunEventSequence")
	}

	return res, nil
}

func (d *DB) insertRawLastRunEventSequencePostgres(tx *sql.Tx, lastruneventsequence *types.LastRunEventSequence) error {
	q := lastRunEventSequenceInsertRawPostgres(lastruneventsequence.ID, lastruneventsequence.Revision, lastruneventsequence.CreationTime, lastruneventsequence.UpdateTime, lastruneventsequence.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert lastRunEventSequence")
	}

	return nil
}
var (
	commitStatusInsertPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inProjectID string, inState types.CommitState, inCommitSHA string, inRunCounter uint64, inDescription string, inContext string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatus").Cols("id", "revision", "creation_time", "update_time", "project_id", "state", "commit_sha", "run_counter", "description", "context").Values(inId, inRevision, inCreationTime, inUpdateTime, inProjectID, inState, inCommitSHA, inRunCounter, inDescription, inContext)
	}
	commitStatusUpdatePostgres = func(curRevision uint64, inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inProjectID string, inState types.CommitState, inCommitSHA string, inRunCounter uint64, inDescription string, inContext string) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("commitstatus").Set(ub.Assign("id", inId), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("project_id", inProjectID), ub.Assign("state", inState), ub.Assign("commit_sha", inCommitSHA), ub.Assign("run_counter", inRunCounter), ub.Assign("description", inDescription), ub.Assign("context", inContext)).Where(ub.E("id", inId), ub.E("revision", curRevision))
	}

	commitStatusInsertRawPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inProjectID string, inState types.CommitState, inCommitSHA string, inRunCounter uint64, inDescription string, inContext string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatus").Cols("id", "revision", "creation_time", "update_time", "project_id", "state", "commit_sha", "run_counter", "description", "context").SQL("OVERRIDING SYSTEM VALUE").Values(inId, inRevision, inCreationTime, inUpdateTime, inProjectID, inState, inCommitSHA, inRunCounter, inDescription, inContext)
	}
)

func (d *DB) insertCommitStatusPostgres(tx *sql.Tx, commitstatus *types.CommitStatus) error {
	q := commitStatusInsertPostgres(commitstatus.ID, commitstatus.Revision, commitstatus.CreationTime, commitstatus.UpdateTime, commitstatus.ProjectID, commitstatus.State, commitstatus.CommitSHA, commitstatus.RunCounter, commitstatus.Description, commitstatus.Context)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatus")
	}

	return nil
}

func (d *DB) updateCommitStatusPostgres(tx *sql.Tx, curRevision uint64, commitstatus *types.CommitStatus) (stdsql.Result, error) {
	q := commitStatusUpdatePostgres(curRevision, commitstatus.ID, commitstatus.Revision, commitstatus.CreationTime, commitstatus.UpdateTime, commitstatus.ProjectID, commitstatus.State, commitstatus.CommitSHA, commitstatus.RunCounter, commitstatus.Description, commitstatus.Context)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update commitStatus")
	}

	return res, nil
}

func (d *DB) insertRawCommitStatusPostgres(tx *sql.Tx, commitstatus *types.CommitStatus) error {
	q := commitStatusInsertRawPostgres(commitstatus.ID, commitstatus.Revision, commitstatus.CreationTime, commitstatus.UpdateTime, commitstatus.ProjectID, commitstatus.State, commitstatus.CommitSHA, commitstatus.RunCounter, commitstatus.Description, commitstatus.Context)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatus")
	}

	return nil
}
var (
	commitStatusDeliveryInsertPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inCommitStatusID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatusdelivery").Cols("id", "revision", "creation_time", "update_time", "commit_status_id", "delivery_status", "delivered_at").Values(inId, inRevision, inCreationTime, inUpdateTime, inCommitStatusID, inDeliveryStatus, inDeliveredAt)
	}
	commitStatusDeliveryUpdatePostgres = func(curRevision uint64, inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inCommitStatusID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("commitstatusdelivery").Set(ub.Assign("id", inId), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("commit_status_id", inCommitStatusID), ub.Assign("delivery_status", inDeliveryStatus), ub.Assign("delivered_at", inDeliveredAt)).Where(ub.E("id", inId), ub.E("revision", curRevision))
	}

	commitStatusDeliveryInsertRawPostgres = func(inId string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inCommitStatusID string, inDeliveryStatus types.DeliveryStatus, inDeliveredAt *time.Time) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("commitstatusdelivery").Cols("id", "revision", "creation_time", "update_time", "sequence", "commit_status_id", "delivery_status", "delivered_at").SQL("OVERRIDING SYSTEM VALUE").Values(inId, inRevision, inCreationTime, inUpdateTime, inSequence, inCommitStatusID, inDeliveryStatus, inDeliveredAt)
	}
)

func (d *DB) insertCommitStatusDeliveryPostgres(tx *sql.Tx, commitstatusdelivery *types.CommitStatusDelivery) error {
	q := commitStatusDeliveryInsertPostgres(commitstatusdelivery.ID, commitstatusdelivery.Revision, commitstatusdelivery.CreationTime, commitstatusdelivery.UpdateTime, commitstatusdelivery.CommitStatusID, commitstatusdelivery.DeliveryStatus, commitstatusdelivery.DeliveredAt)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatusDelivery")
	}

	return nil
}

func (d *DB) updateCommitStatusDeliveryPostgres(tx *sql.Tx, curRevision uint64, commitstatusdelivery *types.CommitStatusDelivery) (stdsql.Result, error) {
	q := commitStatusDeliveryUpdatePostgres(curRevision, commitstatusdelivery.ID, commitstatusdelivery.Revision, commitstatusdelivery.CreationTime, commitstatusdelivery.UpdateTime, commitstatusdelivery.CommitStatusID, commitstatusdelivery.DeliveryStatus, commitstatusdelivery.DeliveredAt)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update commitStatusDelivery")
	}

	return res, nil
}

func (d *DB) insertRawCommitStatusDeliveryPostgres(tx *sql.Tx, commitstatusdelivery *types.CommitStatusDelivery) error {
	q := commitStatusDeliveryInsertRawPostgres(commitstatusdelivery.ID, commitstatusdelivery.Revision, commitstatusdelivery.CreationTime, commitstatusdelivery.UpdateTime, commitstatusdelivery.Sequence, commitstatusdelivery.CommitStatusID, commitstatusdelivery.DeliveryStatus, commitstatusdelivery.DeliveredAt)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert commitStatusDelivery")
	}

	return nil
}
