package xdrbuilder_test

import (
	"testing"

	xdrHelper "github.com/kafakk/xdr-builder"
	"github.com/stretchr/testify/assert"
)

func TestSetAsset(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		code := "ABC"
		issuerPublicKey := "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR"
		resp, err := xdrHelper.SetAsset(code, issuerPublicKey)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong issuerPublicKey", func(t *testing.T) {
		code := "ABC"
		issuerPublicKey := "asfewAS"
		_, err := xdrHelper.SetAsset(code, issuerPublicKey)
		assert.Error(t, err)
	})

}

func TestSetNativeAsset(t *testing.T) {
	resp, err := xdrHelper.SetNativeAsset()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCreateAccount(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		destinationPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
		var startingBalance uint64
		startingBalance = 100
		resp, err := xdrHelper.CreateAccount(destinationPublicKey, startingBalance)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong issuerPublicKey", func(t *testing.T) {
		destinationPublicKey := "asdqw"
		var startingBalance uint64
		startingBalance = 100
		_, err := xdrHelper.CreateAccount(destinationPublicKey, startingBalance)
		assert.Error(t, err)
	})

}

func TestPayment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		destinationPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
		asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		var amount uint64
		amount = 100
		resp, err := xdrHelper.Payment(destinationPublicKey, asset, amount)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Fail, wrong issuerPublicKey", func(t *testing.T) {
		destinationPublicKey := "ssASdasewq"
		asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		var amount uint64
		amount = 100
		_, err = xdrHelper.Payment(destinationPublicKey, asset, amount)
		assert.Error(t, err)
	})

}

func TestManageOffer(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sellingAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		buyingAsset, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		var amount uint64
		amount = 1000
		pricestring := "7.2"
		resp, err := xdrHelper.ManageOffer(sellingAsset, buyingAsset, amount, pricestring)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong priceString", func(t *testing.T) {
		sellingAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		buyingAsset, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		var amount uint64
		amount = 1000
		pricestring := "ABC"
		_, err = xdrHelper.ManageOffer(sellingAsset, buyingAsset, amount, pricestring)
		assert.Error(t, err)
	})
}

func TestCreatePassiveOffer(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sellingAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		buyingAsset, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		var amount uint64
		amount = 10000
		pricestring := "6.2"
		resp, err := xdrHelper.CreatePassiveOffer(sellingAsset, buyingAsset, amount, pricestring)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong priceString", func(t *testing.T) {
		sellingAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		buyingAsset, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		var amount uint64
		amount = 10000
		pricestring := "AWSD"
		_, err = xdrHelper.CreatePassiveOffer(sellingAsset, buyingAsset, amount, pricestring)
		assert.Error(t, err)
	})
}

func TestAllowTrust(t *testing.T) {
	t.Run("Success, AlphaNum4", func(t *testing.T) {
		trustorPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
		asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		authorize := true
		resp, err := xdrHelper.AllowTrust(trustorPublicKey, asset, authorize)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Success, AlphaNum12", func(t *testing.T) {
		trustorPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
		asset, err := xdrHelper.SetAsset("ABCDEFGHSSS", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		authorize := true
		resp, err := xdrHelper.AllowTrust(trustorPublicKey, asset, authorize)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Success, Native", func(t *testing.T) {
		trustorPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
		asset, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		authorize := true
		resp, err := xdrHelper.AllowTrust(trustorPublicKey, asset, authorize)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Fail, wrong trustorPublicKey", func(t *testing.T) {
		trustorPublicKey := "sdwqwwqssasd"
		asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		authorize := true
		_, err = xdrHelper.AllowTrust(trustorPublicKey, asset, authorize)
		assert.Error(t, err)
	})
}

func TestPathPayment(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sendAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		var sendMax uint64
		sendMax = 100
		destinationPublicKey := "GCIQJ3JRXEEAKFL22C43X66B4NKACPWZ27WIMNXGA5CIEHOYWNXD3EQR"
		destAsset, err := xdrHelper.SetAsset("CDF", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		var destAmount uint64
		destAmount = 30
		tempAsset1, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		var path xdrHelper.Path
		path.XDRAsset = append(path.XDRAsset, tempAsset1.XDRAsset)

		resp, err := xdrHelper.PathPayment(sendAsset, sendMax, destinationPublicKey, destAsset, destAmount, path)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong destinationPublicKey", func(t *testing.T) {
		sendAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		var sendMax uint64
		sendMax = 100
		destinationPublicKey := "sdfse"
		destAsset, err := xdrHelper.SetAsset("CDF", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		var destAmount uint64
		destAmount = 30
		tempAsset1, err := xdrHelper.SetNativeAsset()
		if err != nil {
			panic(err)
		}
		var path xdrHelper.Path
		path.XDRAsset = append(path.XDRAsset, tempAsset1.XDRAsset)

		_, err = xdrHelper.PathPayment(sendAsset, sendMax, destinationPublicKey, destAsset, destAmount, path)
		assert.Error(t, err)
	})
}

func TestChangeTrust(t *testing.T) {
	asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	var limit uint64
	limit = 1000
	resp, err := xdrHelper.ChangeTrust(asset, limit)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestBumpSequence(t *testing.T) {
	var sequenceNumber uint64
	sequenceNumber = 1008097543847999
	resp, err := xdrHelper.BumpSequence(sequenceNumber)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestManageData(t *testing.T) {
	name := "name"
	value := "value"
	resp, err := xdrHelper.ManageData(name, value)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSetOptionInflation(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		inflationDestinationPublicKey := "GCIQJ3JRXEEAKFL22C43X66B4NKACPWZ27WIMNXGA5CIEHOYWNXD3EQR"
		resp, err := xdrHelper.SetOptionInflation(inflationDestinationPublicKey)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong inflationDestinationPublicKey", func(t *testing.T) {
		inflationDestinationPublicKey := "Sqwes"
		_, err := xdrHelper.SetOptionInflation(inflationDestinationPublicKey)
		assert.Error(t, err)
	})
}

func TestSetOptionClearFlags(t *testing.T) {
	var clearFlag uint32
	clearFlag = 2
	resp, err := xdrHelper.SetOptionClearFlags(clearFlag)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSetOptionSetFlags(t *testing.T) {
	var setFlag uint32
	setFlag = 1
	resp, err := xdrHelper.SetOptionSetFlags(setFlag)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSetOptionMasterWeight(t *testing.T) {
	var masterWeight uint32
	masterWeight = 4
	resp, err := xdrHelper.SetOptionMasterWeight(masterWeight)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSetOptionThreshold(t *testing.T) {
	var lowThreshold uint32
	var medThreshold uint32
	var highThreshold uint32
	lowThreshold = 1
	medThreshold = 2
	highThreshold = 4
	resp, err := xdrHelper.SetOptionThreshold(lowThreshold, medThreshold, highThreshold)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSetOptionHomeDomain(t *testing.T) {
	domain := "r0ix.com"
	resp, err := xdrHelper.SetOptionHomeDomain(domain)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestSetOptionSigner(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		singerPublicKey := "GBBDXCEBXA3HGMB5NJ6VUOPXJQDL76YXII6HKBHRXBBJK4MB22WE7AFO"
		var singerWeight uint32
		singerWeight = 2
		resp, err := xdrHelper.SetOptionSigner(singerPublicKey, singerWeight)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong singerPublicKey", func(t *testing.T) {
		singerPublicKey := "sadsawq"
		var singerWeight uint32
		singerWeight = 2
		_, err := xdrHelper.SetOptionSigner(singerPublicKey, singerWeight)
		assert.Error(t, err)
	})

}
