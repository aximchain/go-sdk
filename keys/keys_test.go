package keys

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	ctypes "github.com/aximchain/go-sdk/common/types"
	"github.com/aximchain/go-sdk/types/msg"
	"github.com/aximchain/go-sdk/types/tx"
)

func TestRecoveryFromKeyWordsNoError(t *testing.T) {
	mnemonic := "bottom quick strong ranch section decide pepper broken oven demand coin run jacket curious business achieve mule bamboo remain vote kid rigid bench rubber"
	keyManager, err := NewMnemonicKeyManager(mnemonic)
	assert.NoError(t, err)
	acc := keyManager.GetAddr()
	key := keyManager.GetPrivKey()
	assert.Equal(t, "axc1ddt3ls9fjcd8mh69ujdg3fxc89qle2a7e80akd", acc.String())
	assert.NotNil(t, key)
	customPathKey, err := NewMnemonicPathKeyManager(mnemonic, "1'/1/1")
	assert.NoError(t, err)
	assert.Equal(t, "axc1c67nwp7u5adl7gw0ffn3d47kttcm4crjte900f", customPathKey.GetAddr().String())
}

func TestRecoveryFromKeyBaseNoError(t *testing.T) {
	file := "testkeystore.json"
	planText := []byte("Test msg")
	keyManager, err := NewKeyStoreKeyManager(file, "Zjubfd@123")
	assert.NoError(t, err)
	sigs, err := keyManager.GetPrivKey().Sign(planText)
	assert.NoError(t, err)
	valid := keyManager.GetPrivKey().PubKey().VerifyBytes(planText, sigs)
	assert.True(t, valid)
}

func TestRecoveryPrivateKeyNoError(t *testing.T) {
	planText := []byte("Test msg")
	priv := "9579fff0cab07a4379e845a890105004ba4c8276f8ad9d22082b2acbf02d884b"
	keyManager, err := NewPrivateKeyManager(priv)
	assert.NoError(t, err)
	sigs, err := keyManager.GetPrivKey().Sign(planText)
	assert.NoError(t, err)
	valid := keyManager.GetPrivKey().PubKey().VerifyBytes(planText, sigs)
	assert.True(t, valid)
}

func TestSignTxNoError(t *testing.T) {

	test1Mnemonic := "swift slam quote sail high remain mandate sample now stamp title among fiscal captain joy puppy ghost arrow attract ozone situate install gain mean"
	test2Mnemonic := "bottom quick strong ranch section decide pepper broken oven demand coin run jacket curious business achieve mule bamboo remain vote kid rigid bench rubber"

	test1KeyManager, err := NewMnemonicKeyManager(test1Mnemonic)
	assert.NoError(t, err)
	test2KeyManager, err := NewMnemonicKeyManager(test2Mnemonic)
	assert.NoError(t, err)

	test1Addr := test1KeyManager.GetAddr()
	test2Addr := test2KeyManager.GetAddr()
	testCases := []struct {
		msg         msg.Msg
		keyManager  KeyManager
		accountNUm  int64
		sequence    int64
		expectHexTx string
		errMsg      string
	}{
		{msg.CreateSendMsg(test1Addr, ctypes.Coins{ctypes.Coin{Denom: "AXC", Amount: 100000000000000}}, []msg.Transfer{{test2KeyManager.GetAddr(), ctypes.Coins{ctypes.Coin{Denom: "AXC", Amount: 100000000000000}}}}),
			test1KeyManager,
			0,
			1,
			"c601f0625dee0a522a2c87fa0a250a141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa120d0a03415843108080e983b1de1612250a146b571fc0a9961a7ddf45e49a88a4d83941fcabbe120d0a03415843108080e983b1de16126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c9711240d4d3dc5808e9fdb42dce1a3b9f8c43daa94f5401bfec7bfb5b6e910e9db7d0e3216f97664f401bf5750f9b0b55ebfc6cc1ffbdf8df2ccc6e1ebcb94ff7db69fe2001",
			"send message sign error",
		},
		{
			msg.NewTokenIssueMsg(test2Addr, "Bitcoin", "BTC", 1000000000000000, true),
			test2KeyManager,
			1,
			0,
			"a701f0625dee0a3317efab800a146b571fc0a9961a7ddf45e49a88a4d83941fcabbe1207426974636f696e1a034254432080809aa6eaafe3012801126c0a26eb5ae9872103d8f33449356d58b699f6b16a498bd391aa5e051085415d0fe1873939bc1d2e3a124000c3fae4a450bbbe81ec32e88fdc5bdd84a9fe4a537be2404a74f53e58b39bc60cc6769bf686649a7bb4acf4a0382461fa4a0859eab02f78c599af218d2a96f61801",
			"issue message sign error",
		},
		{msg.NewMsgSubmitProposal("list BTC/AXC", "{\"base_asset_symbol\":\"BTC-86A\",\"quote_asset_symbol\":\"AXC\",\"init_price\":100000000,\"description\":\"list BTC/AXC\",\"expire_time\":\"2018-12-24T00:46:05+08:00\"}", msg.ProposalTypeListTradingPair, test1Addr, ctypes.Coins{ctypes.Coin{Denom: "AXC", Amount: 200000000000}}, time.Second),
			test1KeyManager,
			0,
			2,
			"ce02f0625dee0ad901b42d614e0a0c6c697374204254432f4158431298017b22626173655f61737365745f73796d626f6c223a224254432d383641222c2271756f74655f61737365745f73796d626f6c223a22415843222c22696e69745f7072696365223a3130303030303030302c226465736372697074696f6e223a226c697374204254432f415843222c226578706972655f74696d65223a22323031382d31322d32345430303a34363a30352b30383a3030227d180422141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa2a0c0a034158431080a0b787e905308094ebdc03126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c9711240964f552fa46006a72bf68f582a3b4cd323c460be4cd102b3f5711f5a351420156ce1649b6db29ec9a9744366f0d2fa6d8d78c2d77c9accd8158ee825a578afc72002",
			"submit proposal sign error",
		},
		{
			msg.NewMsgVote(test1Addr, 1, msg.OptionYes),
			test1KeyManager,
			0,
			3,
			"9201f0625dee0a1ea1cadd36080112141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa1801126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c971124049402a20543a637060651128caaf678cb9a7a4807b1f616b47b261172982688215b6c9303c87a3f0669d16827cb31ceceb668e05b825cc76dcf60fe647a7125a2003",
			"vote proposal sign error",
		},
		{
			msg.NewDexListMsg(test2Addr, 1, "BTC-86A", "AXC", 100000000),
			test2KeyManager,
			1,
			2,
			"a501f0625dee0a2fb41de13f0a146b571fc0a9961a7ddf45e49a88a4d83941fcabbe10011a074254432d38364122034158432880c2d72f126e0a26eb5ae9872103d8f33449356d58b699f6b16a498bd391aa5e051085415d0fe1873939bc1d2e3a1240ce097f0694b443e3bb1cfbc7fc7a82d9dcf6f329219f7f7c2f01affd674c2a734631db18b57376e7cbb28ae27e96d02a56f89792e2457dd36dda6f7d62e4566a18012002",
			"List tradimg sign error",
		},
		{msg.NewCreateOrderMsg(test1Addr, "1D0E3086E8E4E0A53C38A90D55BD58B34D57D2FA-5", 1, "BTC-86A_AXC", 100000000, 1000000000),
			test1KeyManager,
			0,
			4,
			"d801f0625dee0a64ce6dc0430a141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa122a314430453330383645384534453041353343333841393044353542443538423334443537443246412d351a0b4254432d3836415f415843200228013080c2d72f388094ebdc034001126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c9711240a2fa47684b1372936b498ea8782f0c54df2484bf9cc5289187febb67a05acdd8088113bddb6c6457ea7024715fda2e75534816f77ea618dd9a3d8e538392565b2004",
			"Create order sign error",
		},
		{
			msg.NewCancelOrderMsg(test1Addr, "BTC-86A_AXC", "1D0E3086E8E4E0A53C38A90D55BD58B34D57D2FA-5"),
			test1KeyManager,
			0,
			5,
			"c701f0625dee0a53166e681b0a141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa120b4254432d3836415f4158431a2a314430453330383645384534453041353343333841393044353542443538423334443537443246412d35126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c9711240bc872c9de6599d447ecac895cd7884ccfb5935878b9e7cc368a4c5e7dbf3342b281113255478c0ec5251102bae23a1c6f774560a3c20c99d6b30573da44ca4322005",
			"Cancel order sign error",
		},
		{
			msg.NewFreezeMsg(test1Addr, "AXC", 100000000),
			test1KeyManager,
			0,
			10,
			"9801f0625dee0a24e774b32d0a141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa12034158431880c2d72f126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c9711240bbacf9cab0baeec5b8cec10aab2087edd6e74440f663e79e7d3bb8f5486fde512ac2c05ff3a395f3a8477987e9932d1f8ecd01ddd5834066729d6dcfa55d0597200a",
			"Freeze token sign error",
		},
		{
			msg.NewUnfreezeMsg(test1Addr, "AXC", 100000000),
			test1KeyManager,
			0,
			11,
			"9801f0625dee0a246515ff0d0a141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa12034158431880c2d72f126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c9711240ed4dec8afadadced4e1ff58b27a43bce1b6277c16287d6384bb834c68bc7511b789a3ee00e8bc74f9b407646be285c4b48b18baa32313d941328668710f774a8200b",
			"Unfreeze token sign error",
		},
		{
			msg.NewTokenBurnMsg(test1Addr, "AXC", 100000000),
			test1KeyManager,
			0,
			12,
			"9801f0625dee0a247ed2d2a00a141d0e3086e8e4e0a53c38a90d55bd58b34d57d2fa12034158431880c2d72f126c0a26eb5ae98721027e69d96640300433654e016d218a8d7ffed751023d8efe81e55dedbd6754c971124094554b248c6f005fcd8ee5b74b99c044b005d2ad429c1961ed4fdd977a0ef09b30db193ec9588fcb0727f02f3c37aef14df9b82e9b034fa8f1b42e486e7e24c0200c",
			"Burn token sign error",
		},
		{
			msg.NewMintMsg(test2Addr, "BTC-86A", 100000000),
			test2KeyManager,
			1,
			5,
			"9e01f0625dee0a28467e08290a146b571fc0a9961a7ddf45e49a88a4d83941fcabbe12074254432d3836411880c2d72f126e0a26eb5ae9872103d8f33449356d58b699f6b16a498bd391aa5e051085415d0fe1873939bc1d2e3a1240566e6e0b73444d792b524a228d3024b02e6cf4326e2782d71bc7d4205fbe66555ef011d224d686f0092959d666ff05c2fe3256c2424dc7adb8024f4cc77cefe018012005",
			"Mint token sign error",
		},
	}
	for _, c := range testCases {
		signMsg := tx.StdSignMsg{
			ChainID:       "axcchain-1000",
			AccountNumber: c.accountNUm,
			Sequence:      c.sequence,
			Memo:          "",
			Msgs:          []msg.Msg{c.msg},
			Source:        0,
		}

		rawSignResult, err := c.keyManager.Sign(signMsg)
		signResult := []byte(hex.EncodeToString(rawSignResult))
		assert.NoError(t, err)
		expectHexTx := c.expectHexTx
		assert.True(t, bytes.Equal(signResult, []byte(expectHexTx)), c.errMsg)
	}
}

func TestExportAsKeyStoreNoError(t *testing.T) {
	defer os.Remove("TestGenerateKeyStoreNoError.json")
	km, err := NewKeyManager()
	assert.NoError(t, err)
	encryPlain1, err := km.GetPrivKey().Sign([]byte("test plain"))
	assert.NoError(t, err)
	keyJSONV1, err := km.ExportAsKeyStore("testpassword")
	assert.NoError(t, err)
	bz, err := json.Marshal(keyJSONV1)
	assert.NoError(t, err)
	err = ioutil.WriteFile("TestGenerateKeyStoreNoError.json", bz, 0660)
	assert.NoError(t, err)
	newkm, err := NewKeyStoreKeyManager("TestGenerateKeyStoreNoError.json", "testpassword")
	assert.NoError(t, err)
	encryPlain2, err := newkm.GetPrivKey().Sign([]byte("test plain"))
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(encryPlain1, encryPlain2))
}

func TestExportAsMnemonicNoError(t *testing.T) {
	km, err := NewKeyManager()
	assert.NoError(t, err)
	encryPlain1, err := km.GetPrivKey().Sign([]byte("test plain"))
	assert.NoError(t, err)
	mnemonic, err := km.ExportAsMnemonic()
	assert.NoError(t, err)
	newkm, err := NewMnemonicKeyManager(mnemonic)
	assert.NoError(t, err)
	encryPlain2, err := newkm.GetPrivKey().Sign([]byte("test plain"))
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(encryPlain1, encryPlain2))
	_, err = newkm.ExportAsMnemonic()
	assert.NoError(t, err)
}

func TestExportAsPrivateKeyNoError(t *testing.T) {
	km, err := NewKeyManager()
	assert.NoError(t, err)
	encryPlain1, err := km.GetPrivKey().Sign([]byte("test plain"))
	assert.NoError(t, err)
	pk, err := km.ExportAsPrivateKey()
	assert.NoError(t, err)
	newkm, err := NewPrivateKeyManager(pk)
	assert.NoError(t, err)
	encryPlain2, err := newkm.GetPrivKey().Sign([]byte("test plain"))
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(encryPlain1, encryPlain2))
}

func TestExportAsMnemonicyError(t *testing.T) {
	km, err := NewPrivateKeyManager("9579fff0cab07a4379e845a890105004ba4c8276f8ad9d22082b2acbf02d884b")
	assert.NoError(t, err)
	_, err = km.ExportAsMnemonic()
	assert.Error(t, err)
	file := "testkeystore.json"
	km, err = NewKeyStoreKeyManager(file, "Zjubfd@123")
	assert.NoError(t, err)
	_, err = km.ExportAsMnemonic()
	assert.Error(t, err)
}
