package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a MongoDB user schema in Go
type User struct {
	ID                  primitive.ObjectID  `bson:"_id,omitempty"`
	Name                string              `bson:"sName,omitempty"`
	Username            string              `bson:"sUsername" validate:"required"`
	Email               string              `bson:"sEmail,omitempty"`
	IsEmailVerified     bool                `bson:"bIsEmailVerified,omitempty"`
	MobileNumber        string              `bson:"sMobNum" validate:"required"`
	IsMobileVerified    bool                `bson:"bIsMobVerified,omitempty"`
	ProfilePic          string              `bson:"sProPic,omitempty"`
	UserType            string              `bson:"eType,omitempty"` // Enum: U = User, B = Bot
	Gender              string              `bson:"eGender,omitempty"`
	PushTokens          []string            `bson:"aPushTokens,omitempty"`
	State               string              `bson:"sState,omitempty"`
	City                string              `bson:"sCity,omitempty"`
	Address             string              `bson:"sAddress,omitempty"`
	PinCode             int                 `bson:"nPinCode,omitempty"`
	DeviceTokens        []string            `bson:"aDeviceToken,omitempty"`
	Status              string              `bson:"eStatus,omitempty"` // Enum: Y (Active), N (Inactive)
	ReferredBy          *primitive.ObjectID `bson:"iReferredBy,omitempty"`
	ReferCode           string              `bson:"sReferCode,omitempty"`
	ReferLink           string              `bson:"sReferLink,omitempty"`
	LoginAt             *time.Time          `bson:"dLoginAt,omitempty"`
	DeletedAt           *time.Time          `bson:"dDeletedAt,omitempty"`
	VerificationToken   string              `bson:"sVerificationToken,omitempty"`
	Reason              string              `bson:"sReason,omitempty"`
	ReferrerRewardsOn   string              `bson:"sReferrerRewardsOn,omitempty"`
	ReferAmount         float64             `bson:"nReferAmount,omitempty"`
	ReferrerAmount      float64             `bson:"nReferrerAmount,omitempty"`
	ReferStatus         string              `bson:"eReferStatus,omitempty"`        // Enum: P (Pending)
	Platform            string              `bson:"ePlatform" validate:"required"` // Enum: A (Android), I (iOS), W (Web), O (Other), AD (Admin)
	PolicyID            *primitive.ObjectID `bson:"iPolicyId,omitempty"`
	LoginCount          int                 `bson:"nLogin,omitempty"`
	EligibleForBenefits bool                `bson:"bEligibleForBenifits,omitempty"`
	ProfileLevelID      *primitive.ObjectID `bson:"iProfileLevelId,omitempty"`
	IsEmailUnsubscribed bool                `bson:"bIsEmailUnSubscribe,omitempty"`
	IsUsernameChanged   bool                `bson:"bIsUsernameChanged,omitempty"`
	IDToken             string              `bson:"idToken,omitempty"`
	CountryCode         string              `bson:"sCountryCode,omitempty"`
	KYCStatus           string              `bson:"eKYCStatus,omitempty"` // Enum: p (Pending), s (Started), c (Completed)
	KYCRejectReason     string              `bson:"sKYCRejectReason,omitempty"`
	KYCStartedAt        *time.Time          `bson:"dKYCStartedAt,omitempty"`
	CreatedAt           time.Time           `bson:"dCreatedAt,omitempty"`
	UpdatedAt           time.Time           `bson:"dUpdatedAt,omitempty"`
}
