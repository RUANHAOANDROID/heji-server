package config

import (
	"time"
)

type ClientType string

const (
	ClientPublic ClientType = "public"
	ClientShare  ClientType = "share"
	ClientUser   ClientType = "user"
)

// ClientConfig represents HTTP client / Web UI config options.
type ClientConfig struct {
	Mode             string         `json:"mode"`
	Name             string         `json:"name"`
	About            string         `json:"about"`
	Edition          string         `json:"edition"`
	Version          string         `json:"version"`
	Copyright        string         `json:"copyright"`
	Flags            string         `json:"flags"`
	BaseUri          string         `json:"baseUri"`
	StaticUri        string         `json:"staticUri"`
	CssUri           string         `json:"cssUri"`
	JsUri            string         `json:"jsUri"`
	ManifestUri      string         `json:"manifestUri"`
	ApiUri           string         `json:"apiUri"`
	ContentUri       string         `json:"contentUri"`
	VideoUri         string         `json:"videoUri"`
	WallpaperUri     string         `json:"wallpaperUri"`
	SiteUrl          string         `json:"siteUrl"`
	SiteDomain       string         `json:"siteDomain"`
	SiteAuthor       string         `json:"siteAuthor"`
	SiteTitle        string         `json:"siteTitle"`
	SiteCaption      string         `json:"siteCaption"`
	SiteDescription  string         `json:"siteDescription"`
	SitePreview      string         `json:"sitePreview"`
	LegalInfo        string         `json:"legalInfo"`
	LegalUrl         string         `json:"legalUrl"`
	AppName          string         `json:"appName"`
	AppMode          string         `json:"appMode"`
	AppIcon          string         `json:"appIcon"`
	AppColor         string         `json:"appColor"`
	Restart          bool           `json:"restart"`
	Debug            bool           `json:"debug"`
	Trace            bool           `json:"trace"`
	Test             bool           `json:"test"`
	Demo             bool           `json:"demo"`
	Sponsor          bool           `json:"sponsor"`
	ReadOnly         bool           `json:"readonly"`
	UploadNSFW       bool           `json:"uploadNSFW"`
	Public           bool           `json:"public"`
	AuthMode         string         `json:"authMode"`
	UsersPath        string         `json:"usersPath"`
	LoginUri         string         `json:"loginUri"`
	RegisterUri      string         `json:"registerUri"`
	PasswordLength   int            `json:"passwordLength"`
	PasswordResetUri string         `json:"passwordResetUri"`
	Experimental     bool           `json:"experimental"`
	AlbumCategories  []string       `json:"albumCategories"`
	Tier             int            `json:"tier"`
	Membership       string         `json:"membership"`
	Customer         string         `json:"customer"`
	MapKey           string         `json:"mapKey"`
	DownloadToken    string         `json:"downloadToken,omitempty"`
	PreviewToken     string         `json:"previewToken,omitempty"`
	Disable          ClientDisable  `json:"disable"`
	Count            ClientCounts   `json:"count"`
	Pos              ClientPosition `json:"pos"`
	Years            Years          `json:"years"`
}

// Years represents a list of years.
type Years []int

// ClientDisable represents disabled client features a user cannot turn back on.
type ClientDisable struct {
	WebDAV         bool `json:"webdav"`
	Settings       bool `json:"settings"`
	Places         bool `json:"places"`
	Backups        bool `json:"backups"`
	TensorFlow     bool `json:"tensorflow"`
	Faces          bool `json:"faces"`
	Classification bool `json:"classification"`
	Sips           bool `json:"sips"`
	FFmpeg         bool `json:"ffmpeg"`
	ExifTool       bool `json:"exiftool"`
	Darktable      bool `json:"darktable"`
	RawTherapee    bool `json:"rawtherapee"`
	ImageMagick    bool `json:"imagemagick"`
	HeifConvert    bool `json:"heifconvert"`
	Vectors        bool `json:"vectors"`
	JpegXL         bool `json:"jpegxl"`
	Raw            bool `json:"raw"`
}

// ClientCounts represents photo, video and album counts for the client UI.
type ClientCounts struct {
	All            int `json:"all"`
	Photos         int `json:"photos"`
	Live           int `json:"live"`
	Videos         int `json:"videos"`
	Cameras        int `json:"cameras"`
	Lenses         int `json:"lenses"`
	Countries      int `json:"countries"`
	Hidden         int `json:"hidden"`
	Favorites      int `json:"favorites"`
	Review         int `json:"review"`
	Stories        int `json:"stories"`
	Private        int `json:"private"`
	Albums         int `json:"albums"`
	PrivateAlbums  int `json:"private_albums"`
	Moments        int `json:"moments"`
	PrivateMoments int `json:"private_moments"`
	Months         int `json:"months"`
	PrivateMonths  int `json:"private_months"`
	States         int `json:"states"`
	PrivateStates  int `json:"private_states"`
	Folders        int `json:"folders"`
	PrivateFolders int `json:"private_folders"`
	Files          int `json:"files"`
	People         int `json:"people"`
	Places         int `json:"places"`
	Labels         int `json:"labels"`
	LabelMaxPhotos int `json:"labelMaxPhotos"`
}

type CategoryLabels []CategoryLabel

type CategoryLabel struct {
	LabelUID   string `json:"UID"`
	CustomSlug string `json:"Slug"`
	LabelName  string `json:"Name"`
}

type ClientPosition struct {
	PhotoUID string    `json:"uid"`
	CellID   string    `json:"cid"`
	TakenAt  time.Time `json:"utc"`
	PhotoLat float64   `json:"lat"`
	PhotoLng float64   `json:"lng"`
}
