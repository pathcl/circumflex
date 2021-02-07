package core

import (
	"gitlab.com/tslocum/cview"
)

type ScreenController struct {
	Application      *cview.Application
	Articles         *cview.List
	MainView         *MainView
	Submissions      []*Submissions
	ApplicationState *ApplicationState
}

type Submission struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	Author        string `json:"user"`
	Time          string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	URL           string `json:"url"`
	Domain        string `json:"domain"`
	Type          string `json:"type"`
}

type Submissions struct {
	MappedSubmissions  int
	MappedPages        int
	StoriesListed      int
	PageToFetchFromAPI int
	MaxPages           int
	Entries            []*Submission
}

type ApplicationState struct {
	SubmissionsToShow                     int
	SubmissionsCategory                   int
	SettingsCategory                      int
	HelpScreenCategory                    int
	ScreenHeight                          int
	ScreenWidth                           int
	CurrentPage                           int
	VimNumberRegister                     string
	IsOffline                             bool
	IsReturningFromSuspension             bool
	IsOnHelpScreen                        bool
	IsOnConfigCreationConfirmationMessage bool
}

type MainView struct {
	Grid        *cview.Grid
	Header      *cview.TextView
	LeftMargin  *cview.TextView
	Panels      *cview.Panels
	StatusBar   *cview.TextView
	PageCounter *cview.TextView
	Settings    *cview.TextView
}

type Config struct {
	CommentWidth        int  `mapstructure:"CLX_COMMENT_WIDTH"`
	IndentSize          int  `mapstructure:"CLX_INDENT_SIZE"`
	PreserveRightMargin bool `mapstructure:"CLX_PRESERVE_RIGHT_MARGIN"`
	HighlightHeadlines  int  `mapstructure:"CLX_HIGHLIGHT_HEADLINES"`
	RelativeNumbering   bool `mapstructure:"CLX_RELATIVE_NUMBERING"`
	HideYCJobs          bool `mapstructure:"CLX_HIDE_YC_JOBS"`
}
