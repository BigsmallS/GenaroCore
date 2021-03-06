package genaro

import (
	"encoding/json"
	"github.com/GenaroNetwork/GenaroCore/common"
	"github.com/GenaroNetwork/GenaroCore/core/types"
)

// the field "extra" store the json of ExtraData
type ExtraData struct {
	CommitteeRank           []common.Address                    `json:"committeeRank"` // rank of committee
	LastSynBlockNum         uint64                              `json:"lastBlockNum"`
	LastSynBlockHash        common.Hash                         `json:"lastSynBlockHash"`
	Signature               []byte                              `json:"signature"` // the signature of block broadcaster
	Proportion              []uint64                            `json:"ratio"`
	CommitteeAccountBinding map[common.Address][]common.Address `json:"CommitteeAccountBinding"` // 委员会账号的绑定信息
}

func UnmarshalToExtra(header *types.Header) *ExtraData {
	result := new(ExtraData)
	json.Unmarshal(header.Extra, result)
	return result
}

func ResetHeaderSignature(header *types.Header) {
	extraData := UnmarshalToExtra(header)
	extraData.Signature = nil
	extraByte, _ := json.Marshal(extraData)

	header.Extra = make([]byte, len(extraByte))
	copy(header.Extra, extraByte)
}

func SetHeaderSignature(header *types.Header, signature []byte) {
	extraData := UnmarshalToExtra(header)
	extraData.Signature = make([]byte, len(signature))
	copy(extraData.Signature, signature)
	extraByte, _ := json.Marshal(extraData)
	header.Extra = make([]byte, len(extraByte))
	copy(header.Extra, extraByte)
}

func SetHeaderCommitteeRankList(header *types.Header, committeeRank []common.Address, proportion []uint64) error {
	extraData := UnmarshalToExtra(header)
	extraData.CommitteeRank = make([]common.Address, len(committeeRank))
	copy(extraData.CommitteeRank, committeeRank)
	extraData.Proportion = make([]uint64, len(proportion))
	copy(extraData.Proportion, proportion)
	extraByte, err := json.Marshal(extraData)
	if err != nil {
		return err
	}
	header.Extra = make([]byte, len(extraByte))
	copy(header.Extra, extraByte)
	return nil
}

func SetCommitteeAccountBinding(header *types.Header, committeeAccountBinding map[common.Address][]common.Address) error {
	extraData := UnmarshalToExtra(header)
	extraData.CommitteeAccountBinding = committeeAccountBinding
	extraByte, err := json.Marshal(extraData)
	if err != nil {
		return err
	}
	header.Extra = make([]byte, len(extraByte))
	copy(header.Extra, extraByte)
	return nil
}

func GetHeaderCommitteeRankList(header *types.Header) ([]common.Address, []uint64) {
	extraData := UnmarshalToExtra(header)
	return extraData.CommitteeRank, extraData.Proportion
}

func CreateCommitteeRankByte(address []common.Address) []byte {
	extra := new(ExtraData)
	extra.CommitteeRank = make([]common.Address, len(address))
	copy(extra.CommitteeRank, address)
	extraByte, _ := json.Marshal(extra)
	return extraByte

}

func GetCommitteeAccountBinding(header *types.Header) (committeeAccountBinding map[common.Address][]common.Address) {
	extraData := UnmarshalToExtra(header)
	return extraData.CommitteeAccountBinding
}
