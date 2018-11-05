package xdrbuilder_test

import (
	xdrHelper "stellar-xdr-helper"
	"testing"

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
		startingBalance := 100
		resp, err := xdrHelper.CreateAccount(destinationPublicKey, startingBalance)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail, wrong issuerPublicKey", func(t *testing.T) {
		destinationPublicKey := "asdqw"
		startingBalance := 100
		_, err := xdrHelper.CreateAccount(destinationPublicKey, startingBalance)
		assert.Error(t, err)
	})
	t.Run("Fail, wrong startingBalance", func(t *testing.T) {
		destinationPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
		startingBalance := -1999
		_, err := xdrHelper.CreateAccount(destinationPublicKey, startingBalance)
		assert.Error(t, err)
	})

}

func TestPayment(t *testing.T) {
	destinationPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
	asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	amount := 100
	resp, err := xdrHelper.Payment(destinationPublicKey, asset, amount)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestManageOffer(t *testing.T) {
	sellingAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	buyingAsset, err := xdrHelper.SetNativeAsset()
	if err != nil {
		panic(err)
	}
	amount := 1000
	pricestring := "7.2"
	resp, err := xdrHelper.ManageOffer(sellingAsset, buyingAsset, amount, pricestring)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCreatePassiveOffer(t *testing.T) {
	sellingAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	buyingAsset, err := xdrHelper.SetNativeAsset()
	if err != nil {
		panic(err)
	}
	amount := 10000
	pricestring := "6.2"
	resp, err := xdrHelper.CreatePassiveOffer(sellingAsset, buyingAsset, amount, pricestring)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestAllowTrust(t *testing.T) {
	trustorPublicKey := "GDKV36XRERL7HVQ5GKRAV47ZLEPIZMFM7MMLEO4NKQOWFPL5NCIEW3GR"
	asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	authorize := true
	resp, err := xdrHelper.AllowTrust(trustorPublicKey, asset, authorize)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPathPayment(t *testing.T) {
	sendAsset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	sendMax := 100
	destinationPublicKey := "GCIQJ3JRXEEAKFL22C43X66B4NKACPWZ27WIMNXGA5CIEHOYWNXD3EQR"
	destAsset, err := xdrHelper.SetAsset("CDF", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	destAmount := 30
	tempAsset1, err := xdrHelper.SetNativeAsset()
	if err != nil {
		panic(err)
	}
	var path xdrHelper.Path
	path.XDRAsset = append(path.XDRAsset, tempAsset1.XDRAsset)

	resp, err := xdrHelper.PathPayment(sendAsset, sendMax, destinationPublicKey, destAsset, destAmount, path)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestChangeTrust(t *testing.T) {
	asset, err := xdrHelper.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
	if err != nil {
		panic(err)
	}
	limit := 1000
	resp, err := xdrHelper.ChangeTrust(asset, limit)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestBumpSequence(t *testing.T) {
	sequenceNumber := 1008097543847999
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

func TestSetOption(t *testing.T) {
	inflationDestinationPublicKey := "GCIQJ3JRXEEAKFL22C43X66B4NKACPWZ27WIMNXGA5CIEHOYWNXD3EQR"
	clearFlag := uint32(1)
	setFlag := uint32(2)
	masterWeight := uint32(7)
	lowThreshold := uint32(1)
	medThreshold := uint32(2)
	highThreshold := uint32(3)
	homeDomain := "r0ix.com"
	singerWeight := uint32(3)
	singerPublicKey := "GCWQICM6L52IERBVY3UHZVINP6VB3W52RHDKQGGNRNSXPXK4RB3SFAWE"
	resp, err := xdrHelper.SetOption(inflationDestinationPublicKey, clearFlag, setFlag, masterWeight, lowThreshold,
		medThreshold, highThreshold, homeDomain, singerPublicKey, singerWeight)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
