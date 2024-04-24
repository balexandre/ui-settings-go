package ui_settings

type UiSettingEntry struct {
	Name string `json:"name,omitempty" bson:"name"`
	Value any `json:"value,omitempty" bson:"value"`
	ExpiresAt string `json:"expires_at,omitempty" bson:"expires_at"`
	MerchantCode string `json:"merchant_code,omitempty" bson:"merchant_code"`
	UserId uint64 `json:"user_id,omitempty" bson:"user_id"`
}
