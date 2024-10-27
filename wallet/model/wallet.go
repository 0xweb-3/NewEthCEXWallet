package model

type Wallet struct {
	BaseModel
	UserId      uint64 `gorm:"type:int(11);index:idx_user_id"`
	PrivateKey  string `gorm:"type:varchar(200);not null"`
	PublicKey   string `gorm:"type:varchar(200);not null"`
	Address     string `gorm:"type:varchar(20);not null;index:idx_address"`
	AddressType uint8  `gorm:"type:tinyint(4);not null;comment:'0:用户地址；1:热钱包地址(归集地址)；2:冷钱包地址'"`
}
