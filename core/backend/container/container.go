package container

type Instance struct {
	Artist   *Artiststruct
	Location *Locationstruct
	Dates    *Datestruct
	Relation *Relationstruct
}

type Artiststruct struct {
	Index []Subartist
}

type Subartist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
}

type Relationstruct struct {
	Index []Subrelation
}

type Subrelation struct {
	Id             int
	DatesLocations map[string][]string
}

type Datestruct struct {
	Index []Subdate
}
type Subdate struct {
	Id    int
	Dates []string
}

type Locationstruct struct {
	Index []Sublocation
}

type Sublocation struct {
	Id        int
	Locations []string
}

// Geolocalization type ----------------------------
type Geolocalization struct {
	Response GeoObjectCollectionstruct
}
type GeoObjectCollectionstruct struct {
	GeoObjectCollection FeatureMemberstruct
}
type FeatureMemberstruct struct {
	FeatureMember []SubFeatureMember
}
type SubFeatureMember struct {
	GeoObject GeoObjectstruct
}
type GeoObjectstruct struct {
	Point Pointstruct
}
type Pointstruct struct {
	Pos string
}

// ---------
type LocationMain struct {
	LocationMainField []LocationName
}

type LocationName struct {
	LocationNameField map[string][]string
}

var (
	Mainobject         Instance     = Instance{}
	MapAddressLocation LocationMain = LocationMain{}
)
