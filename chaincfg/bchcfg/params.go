package bchcfg

import (
	"errors"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/checksum0/go-cryptoutils/chainhash"
)

const (
	// DeploymentTestDummy ...
	DeploymentTestDummy = iota

	// DeploymentCSV ...
	DeploymentCSV

	//DefinedDeployments ...
	DefinedDeployments
)

var (
	bigOne = big.NewInt(1)

	mainPowLimit       = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)
	regressionPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)
	testnet3PowLimit   = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)
	simnetPowLimit     = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)
)

var (
	// ErrDuplicateNet ...
	ErrDuplicateNet = errors.New("duplicate Bitcoin network")

	// ErrUnknownHDKeyID ...
	ErrUnknownHDKeyID = errors.New("unknown HD private extended key bytes")
)

var (
	registeredNets      = make(map[BitcoinNet]struct{})
	p2pkhAddrIDs        = make(map[byte]struct{})
	p2shAddrIDs         = make(map[byte]struct{})
	cashAddressPrefixes = make(map[string]struct{})
	hdPrivToPubkeyIDs   = make(map[[4]byte][]byte)
)

// Checkpoint ...
type Checkpoint struct {
	Height int32
	Hash   *chainhash.Hash
}

// DNSSeed ...
type DNSSeed struct {
	Host         string
	HasFiltering bool
}

// ConsensusDeployment ...
type ConsensusDeployment struct {
	BitNumber  uint8
	StartTime  uint64
	ExpireTime uint64
}

// BitcoinNet ...
type BitcoinNet uint32

const (
	// Mainnet ...
	Mainnet BitcoinNet = 0xe8f3e1e3

	// Testnet ...
	Testnet BitcoinNet = 0xfabfb5da

	// Testnet3 ...
	Testnet3 BitcoinNet = 0xf4f3e5f4

	// Simnet ...
	Simnet BitcoinNet = 0x12141c16
)

// Params ...
type Params struct {
	Name         string
	Net          BitcoinNet
	DefaultPort  string
	DNSSeeds     []DNSSeed
	GenesisBlock []byte // TODO: Need block struct...
	GenesisHash  *chainhash.Hash
	PowLimit     *big.Int
	PowLimitBits uint32

	BIP0034Height int32
	BIP0065Height int32
	BIP0066Height int32

	UAHFForkHeight            int32
	DAAForkHeight             int32
	MagneticAnomalyForkHeight int32

	GreatWallActivationTime uint64
	GravitonActivationTime  uint64

	CoinbaseMaturity         uint16
	SubsidyReductionInterval int32

	TargetTimespan             time.Duration
	TargetTimePerBlock         time.Duration
	RetargetAdjustementFactor  int64
	ReduceMinDifficulty        bool
	NoDifficultyAdjustement    bool
	MinDifficultyReductionTime time.Duration
	GenerateSupported          bool

	Checkpoints []Checkpoint

	RuleChangeActivationThreshold uint32
	MinerConfirmationWindow       uint32
	Deployments                   [DefinedDeployments]ConsensusDeployment

	RelayNonSTDTxs bool

	CashAddressPrefix string

	LegacyP2PKHAddrID byte
	LegacyP2SHAddrID  byte

	PrivateKeyID byte

	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte

	HDCoinType uint32
}

// MainnetParams ...
var MainnetParams = Params{
	Name:        "mainnet",
	Net:         Mainnet,
	DefaultPort: "8333",
	DNSSeeds: []DNSSeed{
		{"seed.bchd.cash", true},
		{"seed.bitcoinabc.org", true},
		{"seed-abc.bitcoinforks.org", true},
		{"btccash-seeder.bitcoinunlimited.info", true},
		{"seed.bitprim.org", true},
		{"seed.deadalnix.me", true},
	},

	GenesisBlock:  nil, // TODO
	GenesisHash:   nil, // TODO
	PowLimit:      mainPowLimit,
	PowLimitBits:  0x1d00ffff,
	BIP0034Height: 227931,
	BIP0065Height: 388381,
	BIP0066Height: 363725,

	UAHFForkHeight:            478558,
	DAAForkHeight:             504031,
	MagneticAnomalyForkHeight: 556766,

	GreatWallActivationTime: 1557921600,
	GravitonActivationTime:  1573819200,

	CoinbaseMaturity:           100,
	SubsidyReductionInterval:   210000,
	TargetTimespan:             time.Hour * 24 * 14,
	TargetTimePerBlock:         time.Minute * 10,
	RetargetAdjustementFactor:  4,
	ReduceMinDifficulty:        false,
	NoDifficultyAdjustement:    false,
	MinDifficultyReductionTime: 0,
	GenerateSupported:          false,

	Checkpoints: []Checkpoint{
		{Height: 11111, Hash: newHashFromStr("0000000069e244f73d78e8fd29ba2fd2ed618bd6fa2ee92559f542fdb26e7c1d")},
		{Height: 33333, Hash: newHashFromStr("000000002dd5588a74784eaa7ab0507a18ad16a236e7b1ce69f00d7ddfb5d0a6")},
		{Height: 74000, Hash: newHashFromStr("0000000000573993a3c9e41ce34471c079dcf5f52a0e824a81e7f953b8661a20")},
		{Height: 105000, Hash: newHashFromStr("00000000000291ce28027faea320c8d2b054b2e0fe44a773f3eefb151d6bdc97")},
		{Height: 134444, Hash: newHashFromStr("00000000000005b12ffd4cd315cd34ffd4a594f430ac814c91184a0d42d2b0fe")},
		{Height: 168000, Hash: newHashFromStr("000000000000099e61ea72015e79632f216fe6cb33d7899acb35b75c8303b763")},
		{Height: 193000, Hash: newHashFromStr("000000000000059f452a5f7340de6682a977387c17010ff6e6c3bd83ca8b1317")},
		{Height: 210000, Hash: newHashFromStr("000000000000048b95347e83192f69cf0366076336c639f9b7228e9ba171342e")},
		{Height: 216116, Hash: newHashFromStr("00000000000001b4f4b433e81ee46494af945cf96014816a4e2370f11b23df4e")},
		{Height: 225430, Hash: newHashFromStr("00000000000001c108384350f74090433e7fcf79a606b8e797f065b130575932")},
		{Height: 250000, Hash: newHashFromStr("000000000000003887df1f29024b06fc2200b55f8af8f35453d7be294df2d214")},
		{Height: 267300, Hash: newHashFromStr("000000000000000a83fbd660e918f218bf37edd92b748ad940483c7c116179ac")},
		{Height: 279000, Hash: newHashFromStr("0000000000000001ae8c72a0b0c301f67e3afca10e819efa9041e458e9bd7e40")},
		{Height: 300255, Hash: newHashFromStr("0000000000000000162804527c6e9b9f0563a280525f9d08c12041def0a0f3b2")},
		{Height: 319400, Hash: newHashFromStr("000000000000000021c6052e9becade189495d1c539aa37c58917305fd15f13b")},
		{Height: 343185, Hash: newHashFromStr("0000000000000000072b8bf361d01a6ba7d445dd024203fafc78768ed4368554")},
		{Height: 352940, Hash: newHashFromStr("000000000000000010755df42dba556bb72be6a32f3ce0b6941ce4430152c9ff")},
		{Height: 382320, Hash: newHashFromStr("00000000000000000a8dc6ed5b133d0eb2fd6af56203e4159789b092defd8ab2")},
		{Height: 400000, Hash: newHashFromStr("000000000000000004ec466ce4732fe6f1ed1cddc2ed4b328fff5224276e3f6f")},
		{Height: 430000, Hash: newHashFromStr("000000000000000001868b2bb3a285f3cc6b33ea234eb70facf4dcdf22186b87")},
		{Height: 470000, Hash: newHashFromStr("0000000000000000006c539c722e280a0769abd510af0073430159d71e6d7589")},
		{Height: 510000, Hash: newHashFromStr("00000000000000000367922b6457e21d591ef86b360d78a598b14c2f1f6b0e04")},
		{Height: 552979, Hash: newHashFromStr("0000000000000000015648768ac1b788a83187d706f858919fcc5c096b76fbf2")},
		{Height: 556767, Hash: newHashFromStr("0000000000000000004626ff6e3b936941d341c5932ece4357eeccac44e6d56c")},
	},

	RuleChangeActivationThreshold: 1916,
	MinerConfirmationWindow:       2016,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1199145601, // January 1, 2008 UTC
			ExpireTime: 1230767999, // December 31, 2008 UTC
		},
		DeploymentCSV: {
			BitNumber:  0,
			StartTime:  1462060800, // May 1st, 2016
			ExpireTime: 1493596800, // May 1st, 2017
		},
	},

	RelayNonSTDTxs: false,

	CashAddressPrefix: "bitcoincash",

	LegacyP2PKHAddrID: 0x00,
	LegacyP2SHAddrID:  0x05,
	PrivateKeyID:      0x80,

	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4},
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e},

	HDCoinType: 145,
}

// RegTestnetParams ...
var RegTestnetParams = Params{
	Name:        "regtest",
	Net:         Testnet,
	DefaultPort: "18444",
	DNSSeeds:    []DNSSeed{},

	GenesisBlock:  nil, // TODO
	GenesisHash:   nil, // TODO
	PowLimit:      regressionPowLimit,
	PowLimitBits:  0x207fffff,
	BIP0034Height: 100000000,
	BIP0065Height: 1351,
	BIP0066Height: 1251,

	UAHFForkHeight:            0,
	DAAForkHeight:             0,
	MagneticAnomalyForkHeight: 1000,

	CoinbaseMaturity:           100,
	SubsidyReductionInterval:   150,
	TargetTimespan:             time.Hour * 24 * 14,
	TargetTimePerBlock:         time.Minute * 10,
	RetargetAdjustementFactor:  4,
	ReduceMinDifficulty:        true,
	NoDifficultyAdjustement:    true,
	MinDifficultyReductionTime: time.Minute * 20,
	GenerateSupported:          true,

	Checkpoints: nil,

	RuleChangeActivationThreshold: 108,
	MinerConfirmationWindow:       144,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  0,
			ExpireTime: math.MaxInt64,
		},
		DeploymentCSV: {
			BitNumber:  0,
			StartTime:  0,             // May 1st, 2016
			ExpireTime: math.MaxInt64, // May 1st, 2017
		},
	},

	RelayNonSTDTxs: true,

	CashAddressPrefix: "bchreg",

	LegacyP2PKHAddrID: 0x6f,
	LegacyP2SHAddrID:  0xc4,
	PrivateKeyID:      0xef,

	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94},
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf},

	HDCoinType: 1,
}

// Testnet3Params ...
var Testnet3Params = Params{
	Name:        "testnet3",
	Net:         Testnet3,
	DefaultPort: "18333",
	DNSSeeds: []DNSSeed{
		{"testnet-seed.bitcoinabc.org", true},
		{"testnet-seed-abc.bitcoinforks.org", true},
		{"testnet-seed.bitprim.org", true},
		{"testnet-seed.deadalnix.me", true},
		{"testnet-seeder.criptolayer.net", true},
	},

	GenesisBlock:  nil, // TODO
	GenesisHash:   nil, // TODO
	PowLimit:      testnet3PowLimit,
	PowLimitBits:  0x1d00ffff,
	BIP0034Height: 21111,
	BIP0065Height: 581885,
	BIP0066Height: 330776,

	UAHFForkHeight:            1155875,
	DAAForkHeight:             1188697,
	MagneticAnomalyForkHeight: 1267996,

	GreatWallActivationTime: 1557921600,
	GravitonActivationTime:  1573819200,

	CoinbaseMaturity:           100,
	SubsidyReductionInterval:   210000,
	TargetTimespan:             time.Hour * 24 * 14,
	TargetTimePerBlock:         time.Minute * 10,
	RetargetAdjustementFactor:  4,
	ReduceMinDifficulty:        true,
	NoDifficultyAdjustement:    false,
	MinDifficultyReductionTime: time.Minute * 20,
	GenerateSupported:          false,

	Checkpoints: []Checkpoint{
		{Height: 546, Hash: newHashFromStr("000000002a936ca763904c3c35fce2f3556c559c0214345d31b1bcebf76acb70")},
		{Height: 100000, Hash: newHashFromStr("00000000009e2958c15ff9290d571bf9459e93b19765c6801ddeccadbb160a1e")},
		{Height: 200000, Hash: newHashFromStr("0000000000287bffd321963ef05feab753ebe274e1d78b2fd4e2bfe9ad3aa6f2")},
		{Height: 300001, Hash: newHashFromStr("0000000000004829474748f3d1bc8fcf893c88be255e6d7f571c548aff57abf4")},
		{Height: 400002, Hash: newHashFromStr("0000000005e2c73b8ecb82ae2dbc2e8274614ebad7172b53528aba7501f5a089")},
		{Height: 500011, Hash: newHashFromStr("00000000000929f63977fbac92ff570a9bd9e7715401ee96f2848f7b07750b02")},
		{Height: 600002, Hash: newHashFromStr("000000000001f471389afd6ee94dcace5ccc44adc18e8bff402443f034b07240")},
		{Height: 700000, Hash: newHashFromStr("000000000000406178b12a4dea3b27e13b3c4fe4510994fd667d7c1e6a3f4dc1")},
		{Height: 800010, Hash: newHashFromStr("000000000017ed35296433190b6829db01e657d80631d43f5983fa403bfdb4c1")},
		{Height: 900000, Hash: newHashFromStr("0000000000356f8d8924556e765b7a94aaebc6b5c8685dcfa2b1ee8b41acd89b")},
		{Height: 1000007, Hash: newHashFromStr("00000000001ccb893d8a1f25b70ad173ce955e5f50124261bbbc50379a612ddf")},
	},

	RuleChangeActivationThreshold: 1512,
	MinerConfirmationWindow:       2016,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1199145601, // January 1, 2008 UTC
			ExpireTime: 1230767999, // December 31, 2008 UTC
		},
		DeploymentCSV: {
			BitNumber:  0,
			StartTime:  1462060800, // May 1st, 2016
			ExpireTime: 1493596800, // May 1st, 2017
		},
	},

	RelayNonSTDTxs: true,

	CashAddressPrefix: "bchtest",

	LegacyP2PKHAddrID: 0x6f,
	LegacyP2SHAddrID:  0xc4,
	PrivateKeyID:      0xef,

	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94},
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf},

	HDCoinType: 1,
}

// SimnetParams ...
var SimnetParams = Params{
	Name:        "simnet",
	Net:         Simnet,
	DefaultPort: "18555",
	DNSSeeds:    []DNSSeed{},

	GenesisBlock:  nil, // TODO
	GenesisHash:   nil, // TODO
	PowLimit:      simnetPowLimit,
	PowLimitBits:  0x207fffff,
	BIP0034Height: 0,
	BIP0065Height: 0,
	BIP0066Height: 0,

	UAHFForkHeight:            0,
	DAAForkHeight:             2000,
	MagneticAnomalyForkHeight: 3000,

	CoinbaseMaturity:           100,
	SubsidyReductionInterval:   210000,
	TargetTimespan:             time.Hour * 24 * 14,
	TargetTimePerBlock:         time.Minute * 10,
	RetargetAdjustementFactor:  4,
	ReduceMinDifficulty:        true,
	NoDifficultyAdjustement:    true,
	MinDifficultyReductionTime: time.Minute * 20,
	GenerateSupported:          true,

	Checkpoints: nil,

	RuleChangeActivationThreshold: 75,
	MinerConfirmationWindow:       100,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  0,
			ExpireTime: math.MaxInt64,
		},
		DeploymentCSV: {
			BitNumber:  0,
			StartTime:  0,             // May 1st, 2016
			ExpireTime: math.MaxInt64, // May 1st, 2017
		},
	},

	RelayNonSTDTxs: true,

	CashAddressPrefix: "bchsim",

	LegacyP2PKHAddrID: 0x3f,
	LegacyP2SHAddrID:  0x7b,
	PrivateKeyID:      0x64,

	HDPrivateKeyID: [4]byte{0x04, 0x20, 0xb9, 0x00},
	HDPublicKeyID:  [4]byte{0x04, 0x20, 0xbd, 0x3a},

	HDCoinType: 115,
}

// String ...
func (d DNSSeed) String() string {
	return d.Host
}

// Register ...
func Register(params *Params) error {
	if _, ok := registeredNets[params.Net]; ok {
		return ErrDuplicateNet
	}

	registeredNets[params.Net] = struct{}{}
	p2pkhAddrIDs[params.LegacyP2PKHAddrID] = struct{}{}
	p2shAddrIDs[params.LegacyP2SHAddrID] = struct{}{}
	hdPrivToPubkeyIDs[params.HDPrivateKeyID] = params.HDPublicKeyID[:]

	cashAddressPrefixes[params.CashAddressPrefix+":"] = struct{}{}

	return nil
}

func mustRegister(params *Params) {
	err := Register(params)
	if err != nil {
		panic("failed to register network: " + err.Error())
	}
}

// IsP2PKHAddrID ...
func IsP2PKHAddrID(id byte) bool {
	_, ok := p2pkhAddrIDs[id]
	return ok
}

// IsP2SHAddrID ...
func IsP2SHAddrID(id byte) bool {
	_, ok := p2shAddrIDs[id]
	return ok
}

// IsCashAddressPrefix ...
func IsCashAddressPrefix(prefix string) bool {
	prefix = strings.ToLower(prefix)

	_, ok := cashAddressPrefixes[prefix]
	return ok
}

// HDPrivatekeyToPublickeyID ...
func HDPrivatekeyToPublickeyID(id []byte) ([]byte, error) {
	if len(id) != 4 {
		return nil, ErrUnknownHDKeyID
	}

	var key [4]byte
	copy(key[:], id)

	pubBytes, ok := hdPrivToPubkeyIDs[key]
	if !ok {
		return nil, ErrUnknownHDKeyID
	}

	return pubBytes, nil
}

func newHashFromStr(s string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromString(s)
	if err != nil {
		panic(err)
	}

	return hash
}

func init() {
	mustRegister(&MainnetParams)
	mustRegister(&Testnet3Params)
	mustRegister(&RegTestnetParams)
	mustRegister(&SimnetParams)
}
