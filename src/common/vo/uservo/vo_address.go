package uservo

import (
	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

type Address string

func NewAddress(address string) (Address, error) {
	switch address {
	case Address00001.Value():
	case Address00002.Value():
	case Address00003.Value():
	case Address00004.Value():
	case Address00005.Value():
	case Address00006.Value():
	case Address00007.Value():
	case Address00008.Value():
	case Address00009.Value():
	case Address00010.Value():
	case Address00011.Value():
	case Address00012.Value():
	case Address00013.Value():
	case Address00014.Value():
	case Address00015.Value():
	case Address00016.Value():
	case Address00017.Value():
	case Address00018.Value():
	case Address00019.Value():
	case Address00020.Value():
	case Address00021.Value():
	case Address00022.Value():
	case Address00023.Value():
	case Address00024.Value():
	case Address00025.Value():
	case Address00026.Value():
	case Address00027.Value():
	case Address00028.Value():
	case Address00029.Value():
	case Address00030.Value():
	case Address00031.Value():
	case Address00032.Value():
	case Address00033.Value():
	case Address00034.Value():
	case Address00035.Value():
	case Address00036.Value():
	case Address00037.Value():
	case Address00038.Value():
	case Address00039.Value():
	case Address00040.Value():
	case Address00041.Value():
	case Address00042.Value():
	case Address00043.Value():
	case Address00044.Value():
	case Address00045.Value():
	case Address00046.Value():
	case Address00047.Value():

	default:
		return "", mjerr.Wrap(
			nil,
			mjerr.WithOriginError(apperr.InvalidParameter),
			mjerr.WithLogDetail(
				map[string]interface{}{
					"assetType": address,
				},
			),
		)
	}

	return Address(address), nil
}

func (address Address) Value() string {
	return string(address)
}

func (address Address) Equals(assetType Address) bool {
	return address.Value() == assetType.Value()
}

const (
	Address00001 Address = "00001" // 北海道
	Address00002 Address = "00002" // 青森県
	Address00003 Address = "00003" // 岩手県
	Address00004 Address = "00004" // 宮城県
	Address00005 Address = "00005" // 秋田県
	Address00006 Address = "00006" // 山形県
	Address00007 Address = "00007" // 福島県
	Address00008 Address = "00008" // 茨城県
	Address00009 Address = "00009" // 栃木県
	Address00010 Address = "00010" // 群馬県
	Address00011 Address = "00011" // 埼玉県
	Address00012 Address = "00012" // 千葉県
	Address00013 Address = "00013" // 東京都
	Address00014 Address = "00014" // 神奈川県
	Address00015 Address = "00015" // 新潟県
	Address00016 Address = "00016" // 富山県
	Address00017 Address = "00017" // 石川県
	Address00018 Address = "00018" // 福井県
	Address00019 Address = "00019" // 山梨県
	Address00020 Address = "00020" // 長野県
	Address00021 Address = "00021" // 岐阜県
	Address00022 Address = "00022" // 静岡県
	Address00023 Address = "00023" // 愛知県
	Address00024 Address = "00024" // 三重県
	Address00025 Address = "00025" // 滋賀県
	Address00026 Address = "00026" // 京都府
	Address00027 Address = "00027" // 大阪府
	Address00028 Address = "00028" // 兵庫県
	Address00029 Address = "00029" // 奈良県
	Address00030 Address = "00030" // 和歌山県
	Address00031 Address = "00031" // 鳥取県
	Address00032 Address = "00032" // 島根県
	Address00033 Address = "00033" // 岡山県
	Address00034 Address = "00034" // 広島県
	Address00035 Address = "00035" // 山口県
	Address00036 Address = "00036" // 徳島県
	Address00037 Address = "00037" // 香川県
	Address00038 Address = "00038" // 愛媛県
	Address00039 Address = "00039" // 高知県
	Address00040 Address = "00040" // 福岡県
	Address00041 Address = "00041" // 佐賀県
	Address00042 Address = "00042" // 長崎県
	Address00043 Address = "00043" // 熊本県
	Address00044 Address = "00044" // 大分県
	Address00045 Address = "00045" // 宮崎県
	Address00046 Address = "00046" // 鹿児島県
	Address00047 Address = "00047" // 沖縄県
)
