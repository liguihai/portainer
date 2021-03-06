package portainer

type (
	// Pair defines a key/value string pair
	Pair struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	// CLIFlags represents the available flags on the CLI.
	CLIFlags struct {
		Addr      *string
		Assets    *string
		Data      *string
		Endpoint  *string
		Labels    *[]Pair
		Logo      *string
		Swarm     *bool
		Templates *string
		TLSVerify *bool
		TLSCacert *string
		TLSCert   *string
		TLSKey    *string
	}

	// Settings represents Portainer settings.
	Settings struct {
		Swarm        bool   `json:"swarm"`
		HiddenLabels []Pair `json:"hiddenLabels"`
		Logo         string `json:"logo"`
	}

	// User represent a user account.
	User struct {
		Username string `json:"username"`
		Password string `json:"password,omitempty"`
	}

	// TokenData represents the data embedded in a JWT token.
	TokenData struct {
		Username string
	}

	// EndpointConfiguration represents the data required to connect to a Docker API endpoint.
	EndpointConfiguration struct {
		Endpoint      string
		TLS           bool
		TLSCACertPath string
		TLSCertPath   string
		TLSKeyPath    string
	}

	// CLIService represents a service for managing CLI.
	CLIService interface {
		ParseFlags(version string) (*CLIFlags, error)
		ValidateFlags(flags *CLIFlags) error
	}

	// DataStore defines the interface to manage the data.
	DataStore interface {
		Open() error
		Close() error
	}

	// Server defines the interface to serve the data.
	Server interface {
		Start() error
	}

	// UserService represents a service for managing users.
	UserService interface {
		User(username string) (*User, error)
		UpdateUser(user *User) error
	}

	// CryptoService represents a service for encrypting/hashing data.
	CryptoService interface {
		Hash(data string) (string, error)
		CompareHashAndData(hash string, data string) error
	}

	// JWTService represents a service for managing JWT tokens.
	JWTService interface {
		GenerateToken(data *TokenData) (string, error)
		VerifyToken(token string) error
	}
)

const (
	// APIVersion is the version number of portainer API.
	APIVersion = "1.10.2"
)
