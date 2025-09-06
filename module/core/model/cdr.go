package model

import (
	"time"

	"gorm.io/gorm"
)

type CDR struct {
	gorm.Model
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
	Duration      string `json:"duration"`
	BillSec       string `json:"billsec"`

	StartStamp  string `json:"start_stamp"`
	AnswerStamp string `json:"answer_stamp"`
	EndStamp    string `json:"end_stamp"`
	HangupCause string `json:"hangup_cause"`

	TrunkAddress string `json:"trunk_address"`

	ProtoSpecificHangupCause string `json:"proto_specific_hangup_cause"`
	SipHangupDisposition     string `json:"sip_hangup_disposition"`
	SipInviteFailurePhrase   string `json:"sip_invite_failure_phrase"`
	SipInviteFailureStatus   string `json:"sip_invite_failure_status"`
	SipTermCause             string `json:"sip_term_cause"`
	SipTermStatus            string `json:"sip_term_status"`

	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// func (CDR) TableName() string { return "cdr" }
