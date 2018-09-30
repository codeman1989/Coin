// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte  
	Vout      int     
	Signature []byte
	PubKey    []byte
}
// TXOutput represents a transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

type Transaction struct {
	ID   []byte        //交易唯一ID
	Vin  []TXInput     //交易输入序列
	Vout []TXOutput    //交易输出序列
}