package profile

type UserProfile struct {
	Email             string   `json:"email"`
	Firstname         string   `json:"firstName"`
	Lastname          string   `json:"lastName"`
	Timezone          string   `json:"timeZone"`
	Enabled           bool     `json:"enabled"`
	LastModifiedBy    string   `json:"lastModifiedBy"`
	LastModifiedTs    int64    `json:"lastModifiedTs"`
	LastMoginTs       int64    `json:"lastLoginTs"`
	DisplayName       string   `json:"displayName"`
	AccessKeysAllowed bool     `json:"accessKeysAllowed"`
	DefaultRoleId     string   `json:"defaultRoleId"`
	RoleIds           []string `json:"roleIds"`
}
