package test

import (
	ethkeystore "github.com/ethereum/go-ethereum/accounts/keystore"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-core/core"
	"github.com/iden3/go-iden3-core/db"
	babykeystore "github.com/iden3/go-iden3-core/keystore"
	"github.com/iden3/go-iden3-core/merkletree"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

type Global struct {
	KsEthPath string
	Pass      string
	dbStorage db.Storage
}

var global *Global

func init() {
	// global.identities = make(map[core.ID]*identity)
}

type Identity struct {
	ksBaby *babykeystore.KeyStore
	ksEth  *ethkeystore.KeyStore
	kOp    *babyjub.PublicKeyComp
	// kDis               *common.Address
	// kReen              *common.Address
	// kUpdateRoot        *common.Address
	id *core.ID
	mt *merkletree.MerkleTree
	// genesisProofClaims *core.GenesisProofClaims
}

func SetGlobal(glob *Global) {
	global = glob
}

// func SetPass(pass string) {
// 	global.Pass = pass
// }

func InitStorage() {
	global.dbStorage = db.NewMemoryStorage()
}

func NewIdentity() (*Identity, error) {
	storage := babykeystore.MemStorage([]byte{})
	ksBaby, err := babykeystore.NewKeyStore(&storage, babykeystore.LightKeyStoreParams)
	if err != nil {
		return nil, err
	}
	kOp, err := ksBaby.NewKey([]byte(global.Pass))
	if err != nil {
		return nil, err
	}

	ksEth := ethkeystore.NewKeyStore(global.KsEthPath,
		ethkeystore.StandardScryptN, ethkeystore.StandardScryptP)
	// accKDis, err := ksEth.NewAccount(global.Pass)
	// if err != nil {
	// 	return nil, err
	// }
	// accKReen, err := ksEth.NewAccount(global.Pass)
	// if err != nil {
	// 	return nil, err
	// }
	// accKUpdateRoot, err := ksEth.NewAccount(global.Pass)
	// if err != nil {
	// 	return nil, err
	// }
	// kDis := &accKDis.Address
	// kReen := &accKReen.Address
	// kUpdateRoot := &accKUpdateRoot.Address
	_kOp, err := kOp.Decompress()
	if err != nil {
		return nil, err
	}
	claimKOp := core.NewClaimAuthorizeKSignBabyJub(_kOp).Entry()
	id, _, err := core.CalculateIdGenesis(claimKOp, []*merkletree.Entry{})
	if err != nil {
		return nil, err
	}
	dbStorage := global.dbStorage.WithPrefix(append([]byte("mt:"), id[:]...))
	mt, err := merkletree.NewMerkleTree(dbStorage, 140)
	if err != nil {
		return nil, err
	}
	// proofClaimsList := []core.ProofClaim{proofClaims.KOp, proofClaims.KDis,
	// 	proofClaims.KReen, proofClaims.KUpdateRoot}
	// for _, proofClaim := range proofClaimsList {
	// 	err = mt.Add(&merkletree.Entry{Data: *proofClaim.Leaf})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	for _, claim := range []*merkletree.Entry{claimKOp} {
		err = mt.Add(claim)
		if err != nil {
			return nil, err
		}
	}
	return &Identity{
		ksBaby: ksBaby,
		ksEth:  ksEth,
		kOp:    kOp,
		// kDis:               kDis,
		// kReen:              kReen,
		// kUpdateRoot:        kUpdateRoot,
		id: id,
		// genesisProofClaims: proofClaims,
		mt: mt,
	}, nil
}

func (iden *Identity) ID() string {
	return iden.id.String()
}

func (iden *Identity) Unlock() error {
	return iden.ksBaby.UnlockKey(iden.kOp, []byte(global.Pass))
}

func (iden *Identity) SignKOp(msg []byte) ([]byte, error) {
	sig, _, err := iden.ksBaby.Sign(iden.kOp, []byte("test:"), msg)
	if err != nil {
		return nil, err
	}
	return sig[:], nil
}
