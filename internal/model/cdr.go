package model

import (
	"time"
)

type CDR struct {
	CallID string `json:"call_id" gorm:"type:uuid"`

	AccountName string `json:"account_name"`
	AccountType string `json:"account_type,omitempty"`

	OriginalCallerNumber string `json:"original_caller_number"`
	OriginalCalleeNumber string `json:"original_callee_number"`

	PreviousCallerNumber string `json:"previous_caller_number"`
	PreviousCalleeNumber string `json:"previous_callee_number"`

	UpdatedCallerAccount string `json:"updated_caller_account"`
	UpdatedCallerNumber  string `json:"updated_caller_number"`
	UpdatedCalleeAccount string `json:"updated_callee_account"`
	UpdatedCalleeNumber  string `json:"updated_callee_number"`

	CallDirection string `json:"call_direction"`
	Duration      int    `json:"duration"`
	BillSec       int    `json:"billsec"`

	StartStamp    string `json:"start_stamp"`
	ProgressStamp string `json:"progress_stamp"`
	AnswerStamp   string `json:"answer_stamp"`
	EndStamp      string `json:"end_stamp"`

	StartEpoch    int64 `json:"start_epoch"`
	ProgressEpoch int64 `json:"progress_epoch"`
	AnswerEpoch   int64 `json:"answer_epoch"`
	EndEpoch      int64 `json:"end_epoch"`

	Status               string `json:"status"`
	TerminatedBy         string `json:"terminated_by"`
	HangupCause          string `json:"hangup_cause"`
	SipHangupDisposition string `json:"sip_hangup_disposition"`

	SipTermStatus          string `json:"sip_term_status"`
	SipTermCause           string `json:"sip_term_cause"`
	SipInviteFailureStatus string `json:"sip_invite_failure_status"`
	SipInviteFailurePhrase string `json:"sip_invite_failure_phrase"`
	TrunkAddress           string `json:"trunk_address"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (CDR) TableName() string { return "cdr" }
