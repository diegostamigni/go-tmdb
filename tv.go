package tmdb

import (
	"fmt"
)

// TV struct
type TV struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int
		Name        string
		CreditID    string `json:"credit_id"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	}
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int
		Name string
	}
	NextEpisodeToAir struct {
		AirDate        string `json:"air_date"`
		EpisodeNumber  int    `json:"episode_number"`
		ID             int
		Name           string
		Overview       string
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int     `json:"vote_count"`
	} `json:"next_episode_to_air"`
	Homepage     string
	ID           int
	InProduction bool `json:"in_production"`
	Languages    []string
	LastAirDate  string `json:"last_air_date"`
	Name         string
	Networks     []struct {
		ID        int
		Name      string
		LogoPath  string `json:"logo_path"`
		Iso3166_1 string `json:"origin_country"`
	}
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string
	Popularity          float32
	PosterPath          string `json:"poster_path"`
	ProductionCompanies []struct {
		ID        int
		Name      string
		LogoPath  string `json:"logo_path"`
		Iso3166_1 string `json:"origin_country"`
	} `json:"production_companies"`
	Seasons []struct {
		Name         string
		Overview     string
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	}
	Status            string
	Type              string
	VoteAverage       float32              `json:"vote_average"`
	VoteCount         uint32               `json:"vote_count"`
	AlternativeTitles *TvAlternativeTitles `json:"alternative_titles,omitempty"`
	Changes           *TvChanges           `json:",omitempty"`
	Credits           *TvCredits           `json:",omitempty"`
	Images            *TvImages            `json:",omitempty"`
	Keywords          *TvKeywords          `json:",omitempty"`
	Similar           *TvPagedResults      `json:",omitempty"`
	Translations      *TvTranslations      `json:",omitempty"`
	Videos            *TvVideos            `json:",omitempty"`
	ExternalIDs       *TvExternalIds       `json:"external_ids,omitempty"`
}

// TvShort struct
type TvShort struct {
	Adult         bool     `json:"adult"`
	BackdropPath  string   `json:"backdrop_path"`
	ID            int      `json:"id"`
	OriginalName  string   `json:"original_name"`
	GenreIDs      []int32  `json:"genre_ids"`
	OriginCountry []string `json:"origin_country"`
	Popularity    float32  `json:"popularity"`
	PosterPath    string   `json:"poster_path"`
	FirstAirDate  string   `json:"first_air_date"`
	Name          string   `json:"name"`
	Overview      string   `json:"overview"`
	Video         bool     `json:"video"`
	VoteAverage   float32  `json:"vote_average"`
	VoteCount     uint32   `json:"vote_count"`
}

// TvPagedResults struct
type TvPagedResults struct {
	ID                int `json:",omitempty"`
	Page              int
	Results           []TvShort
	TotalPages        int                  `json:"total_pages"`
	TotalResults      int                  `json:"total_results"`
	AlternativeTitles *TvAlternativeTitles `json:"alternative_titles,omitempty"`
	Changes           *TvChanges           `json:",omitempty"`
	Credits           *TvCredits           `json:",omitempty"`
	Images            *TvImages            `json:",omitempty"`
	Keywords          *TvKeywords          `json:",omitempty"`
	Similar           *TvPagedResults      `json:",omitempty"`
	Translations      *TvTranslations      `json:",omitempty"`
	Videos            *TvVideos            `json:",omitempty"`
}

// TvAccountState struct
type TvAccountState struct {
	ID        int
	Favorite  bool
	Watchlist bool
	Rated     bool
}

// TvAlternativeTitles struct
type TvAlternativeTitles struct {
	ID      int
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string
	}
}

// TvChanges struct
type TvChanges struct {
	Changes []struct {
		Key   string
		Items []struct {
			ID     string
			Action string
			Time   string
		}
	}
}

// TvCredits struct
type TvCredits struct {
	ID   int
	Cast []struct {
		Character   string
		CreditID    string `json:"credit_id"`
		ID          int
		Name        string
		Gender      int `json:"gender"`
		Order       int
		ProfilePath string `json:"profile_path"`
	}
	AlternativeTitles *TvAlternativeTitles `json:"alternative_titles,omitempty"`
	Changes           *TvChanges           `json:",omitempty"`
	Crew              []struct {
		CreditID    string `json:"credit_id"`
		Department  string
		Gender      int `json:"gender"`
		ID          int
		Name        string
		Job         string
		ProfilePath string `json:"profile_path"`
	}
	Images       *TvImages       `json:",omitempty"`
	Keywords     *TvKeywords     `json:",omitempty"`
	Similar      *TvPagedResults `json:",omitempty"`
	Translations *TvTranslations `json:",omitempty"`
	Videos       *TvVideos       `json:",omitempty"`
}

// TvExternalIds struct
type TvExternalIds struct {
	ID          int
	ImdbID      string `json:"imdb_id"`
	FreebaseID  string `json:"freebase_id"`
	FreebaseMid string `json:"freebase_mid"`
	TvdbID      int    `json:"tvdb_id"`
	TvrageID    int    `json:"tvrage_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
}

// TvImage struct
type TvImage struct {
	FilePath    string `json:"file_path"`
	Width       int
	Height      int
	Iso639_1    string  `json:"iso_639_1"`
	AspectRatio float32 `json:"aspect_ratio"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   uint32  `json:"vote_count"`
}

// TvImages struct
type TvImages struct {
	ID        int
	Backdrops []TvImage
	Posters   []TvImage
}

// TvKeywords struct
type TvKeywords struct {
	ID      int
	Results []struct {
		ID   int
		Name string
	}
	AlternativeTitles *TvAlternativeTitles `json:"alternative_titles,omitempty"`
	Changes           *TvChanges           `json:",omitempty"`
	Credits           *TvCredits           `json:",omitempty"`
	Images            *TvImages            `json:",omitempty"`
	Similar           *TvPagedResults      `json:",omitempty"`
	Translations      *TvTranslations      `json:",omitempty"`
	Videos            *TvVideos            `json:",omitempty"`
}

// TvRecommendations struct for TV show recommendations.
type TvRecommendations struct {
	Page    int `json:"page"`
	Results []struct {
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIDs         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		PosterPath       string   `json:"poster_path"`
		Popularity       float32  `json:"popularity"`
		Name             string   `json:"name"`
		Networks         []struct {
			ID   int `json:"id"`
			Logo struct {
				Path        string  `json:"path"`
				AspectRatio float32 `json:"aspect_ratio"`
			} `json:"logo"`
			Name          string `json:"name"`
			OriginCountry string `json:"origin_country"`
		} `json:"networks"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   uint32  `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// TvTranslations struct
type TvTranslations struct {
	ID           int
	Translations []struct {
		Iso3166_1   string `json:"iso_3166_1"`
		Iso639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Name     string `json:"name,omitempty"`
			Overview string `json:"overview,omitempty"`
			Homepage string `json:"homepage,omitempty"`
		} `json:"data"`
	}
}

// TvVideos struct
type TvVideos struct {
	ID      int
	Results []struct {
		ID       string
		Iso639_1 string `json:"iso_639_1"`
		Key      string
		Name     string
		Site     string
		Size     int
		Type     string
	}
}

// GetTvInfo gets the primary information about a TV series by id
// https://developers.themoviedb.org/3/tv/get-tv-details
func (tmdb *TMDb) GetTvInfo(id int, options map[string]string) (*TV, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var tvInfo TV
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &tvInfo)
	return result.(*TV), err
}

// GetTvAccountStates gets the status of whether or not the TV show has been rated or added to their favourite or watch lists
// https://developers.themoviedb.org/3/tv/get-tv-account-states
func (tmdb *TMDb) GetTvAccountStates(id int, sessionID string) (*TvAccountState, error) {
	var state TvAccountState
	uri := fmt.Sprintf("%s/tv/%v/account_states?api_key=%s&session_id=%s", baseURL, id, tmdb.apiKey, sessionID)
	result, err := getTmdb(uri, &state)
	return result.(*TvAccountState), err
}

// GetTvAiringToday gets the list of TV shows that air today
// https://developers.themoviedb.org/3/tv/get-tv-airing-today
func (tmdb *TMDb) GetTvAiringToday(options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {},
		"timezone": {}}
	var onAir TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/airing_today?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &onAir)
	return result.(*TvPagedResults), err
}

// GetTvAlternativeTitles gets the alternative titles for a specific show id
// https://developers.themoviedb.org/3/tv/get-tv-alternative-titles
func (tmdb *TMDb) GetTvAlternativeTitles(id int) (*TvAlternativeTitles, error) {
	var titles TvAlternativeTitles
	uri := fmt.Sprintf("%s/tv/%v/alternative_titles?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &titles)
	return result.(*TvAlternativeTitles), err
}

// GetTvChanges gets the changes for a specific show id
// https://developers.themoviedb.org/3/tv/get-tv-changes
func (tmdb *TMDb) GetTvChanges(id int, options map[string]string) (*TvChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes TvChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	return result.(*TvChanges), err
}

// GetTvCredits gets the credits for a specific TV show id
// https://developers.themoviedb.org/3/tv/get-tv-credits
func (tmdb *TMDb) GetTvCredits(id int, options map[string]string) (*TvCredits, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var credits TvCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*TvCredits), err
}

// GetTvExternalIds gets the external ids for a TV series
// https://developers.themoviedb.org/3/tv/get-tv-external-ids
func (tmdb *TMDb) GetTvExternalIds(showID int, options map[string]string) (*TvExternalIds, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var ids TvExternalIds
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/external_ids?api_key=%s%s", baseURL, showID, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &ids)
	return result.(*TvExternalIds), err
}

// GetTvImages gets the images for a TV series
// https://developers.themoviedb.org/3/tv/get-tv-images
func (tmdb *TMDb) GetTvImages(id int, options map[string]string) (*TvImages, error) {
	var availableOptions = map[string]struct{}{
		"language":               {},
		"include_image_language": {}}
	var images TvImages
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/images?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &images)
	return result.(*TvImages), err
}

// GetTvKeywords gets the keywords for a specific TV show id
// https://developers.themoviedb.org/3/tv/get-tv-keywords
func (tmdb *TMDb) GetTvKeywords(id int, options map[string]string) (*TvKeywords, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var keywords TvKeywords
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/keywords?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &keywords)
	return result.(*TvKeywords), err
}

// GetTvRecommendations gets the list of TV show recommendations by id
// https://developers.themoviedb.org/3/tv/get-tv-recommendations
func (tmdb *TMDb) GetTvRecommendations(id int, options map[string]string) (*TvRecommendations, error) {
	var availableOptions = map[string]struct{}{
		"language": {},
		"page":     {}}
	var tvRec TvRecommendations
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/recommendations?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &tvRec)
	return result.(*TvRecommendations), err
}

// GetTvLatest gets the latest TV show
// https://developers.themoviedb.org/3/tv/get-latest-tv
func (tmdb *TMDb) GetTvLatest() (*TV, error) {
	var tv TV
	uri := fmt.Sprintf("%s/tv/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &tv)
	return result.(*TV), err
}

// GetTvOnTheAir gets the list of TV shows that are currently on the air
// https://developers.themoviedb.org/3/tv/get-tv-on-the-air
func (tmdb *TMDb) GetTvOnTheAir(options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var onAir TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/on_the_air?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &onAir)
	return result.(*TvPagedResults), err
}

// GetTvPopular gets the list of popular TV shows
// https://developers.themoviedb.org/3/tv/get-popular-tv-shows
func (tmdb *TMDb) GetTvPopular(options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var onAir TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/popular?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &onAir)
	return result.(*TvPagedResults), err
}

// GetTvSimilar gets the similar TV shows for a specific tv show id
// https://developers.themoviedb.org/3/tv/get-similar-tv-shows
func (tmdb *TMDb) GetTvSimilar(id int, options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":               {},
		"language":           {},
		"append_to_response": {}}
	var similar TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/similar?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &similar)
	return result.(*TvPagedResults), err
}

// GetTvTopRated gets the list of top rated TV shows
// https://developers.themoviedb.org/3/tv/get-top-rated-tv
func (tmdb *TMDb) GetTvTopRated(options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var onAir TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/top_rated?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &onAir)
	return result.(*TvPagedResults), err
}

// GetTvTranslations gets the list of translations that exist for a TV series
// https://developers.themoviedb.org/3/tv/get-tv-translations
func (tmdb *TMDb) GetTvTranslations(id int) (*TvTranslations, error) {
	var translations TvTranslations
	uri := fmt.Sprintf("%s/tv/%v/translations?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &translations)
	return result.(*TvTranslations), err
}

// GetTvVideos gets the videos that have been added to a TV series
// https://developers.themoviedb.org/3/tv/get-tv-videos
func (tmdb *TMDb) GetTvVideos(id int, options map[string]string) (*TvVideos, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var videos TvVideos
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/videos?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &videos)
	return result.(*TvVideos), err
}
