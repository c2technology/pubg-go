package pubg

type GameMode string

type Shard string

const (
	Duo                      = GameMode("duo")
	DuoFPP                   = GameMode("duo-fpp")
	Solo                     = GameMode("solo")
	Squad                    = GameMode("squad")
	SquadFPP                 = GameMode("squad-fpp")
	XBoxAsia                 = Shard("xbox-as")
	XBoxEurope               = Shard("xbox-eu")
	XBoxNorthAmerica         = Shard("xbox-na")
	XBoxOceania              = Shard("xbox-oc")
	PCKoreaJapan             = Shard("pc-krjp")
	PCNorthAmerica           = Shard("pc-na")
	PCEurope                 = Shard("pc-eu")
	PCOceania                = Shard("pc-oc")
	PCKakao                  = Shard("pc-kakao")
	PCSouthEastAsia          = Shard("pc-sea")
	PCSouthAndCentralAmerica = Shard("pc-sa")
	PCAsia                   = Shard("pc-as")
)

type Tags struct {
	Description string `json:"description,omitempty"`
}
type Stats struct {
	Assists         int     `json:"assists,omitempty"`
	Boosts          int     `json:"boosts,omitempty"`
	DamageDealt     float32 `json:"damageDealt,omitempty"`
	DBNOs           int     `json:"DBNOs,omitempty"`
	DeathType       string  `json:"deathType,omitempty"`
	Description     string  `json:"description,omitempty"`
	HeadshotKills   int     `json:"headshotKills,omitempty"`
	Heals           int     `json:"heals,omitempty"`
	KillPlace       int     `json:"killPlace,omitempty"`
	KillPoints      int     `json:"killPoints,omitempty"`
	KillPointsDelta float32 `json:"killPointsDelta,omitempty"`
	KillStreaks     int     `json:"killStreaks,omitempty"`
	Kills           int     `json:"kills,omitempty"`
	LastKillPoints  int     `json:"lastKillPoints,omitempty"`
	LastWinPoints   int     `json:"lastWinPoints,omitempty"`
	LongestKill     int     `json:"longestKill,omitempty"`
	MostDamage      int     `json:"mostDamage,omitempty"`
	Name            string  `json:"name,omitempty"`
	PlayerID        string  `json:"playerId,omitempty"`
	Rank            int     `json:"rank,omitempty"`
	Revives         int     `json:"revives,omitempty"`
	RideDistance    float32 `json:"rideDistance,omitempty"`
	RoadKills       int     `json:"roadKills,omitempty"`
	TeamID          int     `json:"teamId,omitempty"`
	TeamKills       int     `json:"teamKills,omitempty"`
	TimeSurvived    int     `json:"timeSurvived,omitempty"`
	VehicleDestroys int     `json:"vehicleDestroys,omitempty"`
	WalkDistance    float64 `json:"walkDistance,omitempty"`
	WeaponsAcquired int     `json:"weaponsAcquired,omitempty"`
	WinPlace        int     `json:"winPlace,omitempty"`
	WinPoints       int     `json:"winPoints,omitempty"`
	WinPointsDelta  float32 `json:"winPointsDelta,omitempty"`
}

type Rounds struct {
	Description string `json:"description,omitempty"`
}

type Spectators struct {
	Description string `json:"description,omitempty"`
}
type Attributes struct {
	Actor        string `json:"actor,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	Description  string `json:"description,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	GameMode     string `json:"gameMode,omitempty"`
	Name         string `json:"Name,omitempty"`
	Name2        string `json:"name,omitempty"`
	PatchVersion string `json:"patchVersion,omitempty"`
	ReleasedAt   string `json:"releasedAt,omitempty"`
	ShardID      string `json:"shardId,omitempty"`
	Stats        *Stats `json:"stats,omitempty"`
	Tags         *Tags  `json:"tags,omitempty"`
	TitleID      string `json:"titleId,omitempty"`
	URL          string `json:"URL,omitempty"`
	Version      string `json:"version,omitempty"`
	Won          string `json:"won,omitempty"`
}

type Asset struct {
	Data []*Data `json:"data,omitempty"`
}
type Roster struct {
	Data []*Data `json:"data,omitempty"`
}

type Matches struct {
	Data []*Data `json:"data,omitempty"`
}

type Participants struct {
	Data []*Data `json:"data,omitempty"`
}

type Team struct {
	Data []*Data `json:"data,omitempty"`
}

type Relationship struct {
	Assets       *Asset        `json:"assets,omitempty"`
	Description  string        `json:"description,omitempty"`
	Matches      *Matches      `json:"matches,omitempty"`
	Participants *Participants `json:"participants,omitempty"`
	Rosters      *Roster       `json:"rosters,omitempty"`
	Rounds       *Rounds       `json:"rounds,omitempty"`
	Spectators   *Spectators   `json:"spectators,omitempty"`
	Team         *Team         `json:"team,omitempty"`
}

type Link struct {
	First    string `json:"first,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Schema   string `json:"schema,omitempty"`
	Self     string `json:"self,omitempty"`
}

type Data struct {
	Attributes    *Attributes   `json:"attributes,omitempty"`
	Description   string        `json:"description,omitempty"`
	ID            string        `json:"id,omitempty"`
	Links         *Link         `json:"links,omitempty"`
	Relationships *Relationship `json:"relationships,omitempty"`
	Type          string        `json:"type,omitempty"`
}

type Meta struct {
}

//The status endpoint can be called to verify that the API is up and running. It also provides the most recent release date and version of the API service itself.
type Status struct {
	Data *Data `json:"data,omitempty"`
}

//Match contains the results of a completed match such as the game mode played, duration, and which players participated
type Match struct {
	Data *Data `json:"data,omitempty"`
	//Included string `json:"included,omitempty"`
	Links *Link `json:"links,omitempty"`
	Meta  *Meta `json:"meta,omitempty"`
}

//Player contains aggregated lifetime information about each player.
type Player struct {
	Data       *Data       `json:"data,omitempty"`
	Links      *Link       `json:"links,omitempty"`
	Meta       *Meta       `json:"meta,omitempty"`
}

//-------------------------

//Participant represents each player in the context of a match. Participant objects are only meaningful within the context of a match and are not exposed as a standalone resource.
//type Participant struct {
//	description   string         `json:"description,omitempty"`
//	ID            string         `json:"id,omitempty"`
//	Attributes    []Attribute    `json:"attributes,omitempty"`
//	Relationships []Relationship `json:"relationships,omitempty"`
//	Links         []Link         `json:"links,omitempty"`
//}