package uguis

// Tweet represents a tweet.
type Tweet struct {
	Contributors interface{} `json:"contributors"`
	Coordinates  interface{} `json:"coordinates"`
	CreatedAt    string      `json:"created_at"`
	Entities     struct {
		Hashtags []struct {
			Indices []float64 `json:"indices"`
			Text    string    `json:"text"`
		} `json:"hashtags"`
		Symbols []interface{} `json:"symbols"`
		Urls    []struct {
			DisplayURL  string    `json:"display_url"`
			ExpandedURL string    `json:"expanded_url"`
			Indices     []float64 `json:"indices"`
			URL         string    `json:"url"`
		} `json:"urls"`
		UserMentions []interface{} `json:"user_mentions"`
	} `json:"entities"`
	FavoriteCount        float64     `json:"favorite_count"`
	Favorited            bool        `json:"favorited"`
	Geo                  interface{} `json:"geo"`
	ID                   float64     `json:"id"`
	IDStr                string      `json:"id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	Lang                 string      `json:"lang"`
	Place                interface{} `json:"place"`
	PossiblySensitive    bool        `json:"possibly_sensitive"`
	RetweetCount         float64     `json:"retweet_count"`
	Retweeted            bool        `json:"retweeted"`
	Source               string      `json:"source"`
	Text                 string      `json:"text"`
	Truncated            bool        `json:"truncated"`
	User                 struct {
		ContributorsEnabled bool   `json:"contributors_enabled"`
		CreatedAt           string `json:"created_at"`
		DefaultProfile      bool   `json:"default_profile"`
		DefaultProfileImage bool   `json:"default_profile_image"`
		Description         string `json:"description"`
		Entities            struct {
			Description struct {
				Urls []struct {
					DisplayURL  string    `json:"display_url"`
					ExpandedURL string    `json:"expanded_url"`
					Indices     []float64 `json:"indices"`
					URL         string    `json:"url"`
				} `json:"urls"`
			} `json:"description"`
			URL struct {
				Urls []struct {
					DisplayURL  string    `json:"display_url"`
					ExpandedURL string    `json:"expanded_url"`
					Indices     []float64 `json:"indices"`
					URL         string    `json:"url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount                float64 `json:"favourites_count"`
		FollowRequestSent              bool    `json:"follow_request_sent"`
		FollowersCount                 float64 `json:"followers_count"`
		Following                      bool    `json:"following"`
		FriendsCount                   float64 `json:"friends_count"`
		GeoEnabled                     bool    `json:"geo_enabled"`
		ID                             float64 `json:"id"`
		IDStr                          string  `json:"id_str"`
		IsTranslationEnabled           bool    `json:"is_translation_enabled"`
		IsTranslator                   bool    `json:"is_translator"`
		Lang                           string  `json:"lang"`
		ListedCount                    float64 `json:"listed_count"`
		Location                       string  `json:"location"`
		Name                           string  `json:"name"`
		Notifications                  bool    `json:"notifications"`
		ProfileBackgroundColor         string  `json:"profile_background_color"`
		ProfileBackgroundImageURL      string  `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHttps string  `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool    `json:"profile_background_tile"`
		ProfileImageURL                string  `json:"profile_image_url"`
		ProfileImageURLHTTPS           string  `json:"profile_image_url_https"`
		ProfileLinkColor               string  `json:"profile_link_color"`
		ProfileSidebarBorderColor      string  `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string  `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string  `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool    `json:"profile_use_background_image"`
		Protected                      bool    `json:"protected"`
		ScreenName                     string  `json:"screen_name"`
		StatusesCount                  float64 `json:"statuses_count"`
		TimeZone                       string  `json:"time_zone"`
		URL                            string  `json:"url"`
		UtcOffset                      float64 `json:"utc_offset"`
		Verified                       bool    `json:"verified"`
	} `json:"user"`
}
