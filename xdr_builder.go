package xdrbuilder

import (
	"fmt"

	"github.com/stellar/go/price"
	"github.com/stellar/go/xdr"
)

type Asset struct {
	XDRAsset xdr.Asset
}

type Path struct {
	XDRAsset []xdr.Asset //max size is 5
}

func SetAsset(code, issuerPublicKey string) (respAsset Asset, err error) {
	var issuer xdr.AccountId
	err = issuer.SetAddress(issuerPublicKey)
	if err != nil {
		return respAsset, err
	}
	// var asset AssetHelp

	err = respAsset.XDRAsset.SetCredit(code, issuer)

	return respAsset, err
}

func SetNativeAsset() (respAsset Asset, err error) {
	err = respAsset.XDRAsset.SetNative()
	return respAsset, err
}

func CreateAccount(destinationPublicKey string, startingBalance int) (respXDR string, err error) {
	var destination xdr.AccountId
	err = destination.SetAddress(destinationPublicKey)
	if err != nil {
		return respXDR, err
	}
	op := xdr.CreateAccountOp{
		Destination:     destination,
		StartingBalance: xdr.Int64(startingBalance) * 10000000,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err

}

func Payment(destinationPublicKey string, asset Asset, amount int) (respXDR string, err error) {
	var destination xdr.AccountId
	err = destination.SetAddress(destinationPublicKey)
	if err != nil {
		return respXDR, err
	}

	op := xdr.PaymentOp{
		Destination: destination,
		Asset:       asset.XDRAsset,
		Amount:      xdr.Int64(amount) * 10000000,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err

}

func ManageOffer(selling Asset, buying Asset, amount int, priceString string) (respXDR string, err error) {

	price, err := price.Parse(priceString)
	if err != nil {
		return respXDR, err
	}
	op := xdr.ManageOfferOp{
		Selling: selling.XDRAsset,
		Buying:  buying.XDRAsset,
		Amount:  xdr.Int64(amount) * 10000000,
		Price:   price,
	}
	respXDR, err = xdr.MarshalBase64(op)
	return respXDR, err
}

func CreatePassiveOffer(selling Asset, buying Asset, amount int, priceString string) (respXDR string, err error) {

	price, err := price.Parse(priceString)
	if err != nil {
		return respXDR, err
	}
	op := xdr.CreatePassiveOfferOp{
		Selling: selling.XDRAsset,
		Buying:  buying.XDRAsset,
		Amount:  xdr.Int64(amount) * 10000000,
		Price:   price,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func AllowTrust(trustorPublicKey string, asset Asset, authorize bool) (respXDR string, err error) {

	// asset.Extract(asset.Type)

	// var assetType xdr.AllowTrustOpAsset
	var allowTrustAsset xdr.AllowTrustOpAsset

	assetType := asset.XDRAsset.Type

	fmt.Println(assetType)

	switch assetType {
	case 0:
		return "This is Native Asset", nil
	case 1:
		code, ok := asset.XDRAsset.GetAlphaNum4()
		if ok == false {
			panic("The value is not valid")
		}
		allowTrustAsset = xdr.AllowTrustOpAsset{
			Type:       1,
			AssetCode4: &code.AssetCode,
		}
	case 2:
		code, ok := asset.XDRAsset.GetAlphaNum12()
		if ok == false {
			panic("The value is not valid")
		}
		allowTrustAsset = xdr.AllowTrustOpAsset{
			Type:        1,
			AssetCode12: &code.AssetCode,
		}
	}

	var trustor xdr.AccountId
	err = trustor.SetAddress(trustorPublicKey)
	if err != nil {
		return respXDR, err
	}

	op := xdr.AllowTrustOp{
		Trustor:   trustor,
		Asset:     allowTrustAsset,
		Authorize: authorize,
	}
	respXDR, err = xdr.MarshalBase64(op)
	return respXDR, err
}

func PathPayment(sendAsset Asset, sendMax int, destinationPublicKey string, destAsset Asset, destAmount int, path Path) (respXDR string, err error) {

	var destination xdr.AccountId
	err = destination.SetAddress(destinationPublicKey)
	if err != nil {
		return respXDR, err
	}
	op := xdr.PathPaymentOp{
		SendAsset:   sendAsset.XDRAsset,
		SendMax:     xdr.Int64(sendMax) * 10000000,
		Destination: destination,
		DestAsset:   destAsset.XDRAsset,
		DestAmount:  xdr.Int64(destAmount) * 10000000,
		Path:        path.XDRAsset,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func ChangeTrust(asset Asset, limit int) (respXDR string, err error) {
	op := xdr.ChangeTrustOp{
		Line:  asset.XDRAsset,
		Limit: xdr.Int64(limit),
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func BumpSequence(sequenceNumber int) (respXDR string, err error) {
	op := xdr.BumpSequenceOp{
		BumpTo: xdr.SequenceNumber(sequenceNumber),
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func ManageData(name, value string) (respXDR string, err error) {
	dataname := xdr.String64(name)

	valueByte := []byte(value)
	datavalue := xdr.DataValue(valueByte)

	op := xdr.ManageDataOp{
		DataName:  dataname,
		DataValue: &datavalue,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func SetOption(inflationDestinationPublicKey string, clearFlag, setFlag, masterWeight,
	lowThreshold, medThreshold, highThreshold uint32, homeDomain string, singerPublicKey string, singerWeight uint32) (respXDR string, err error) {

	var destination xdr.AccountId
	err = destination.SetAddress(inflationDestinationPublicKey)
	if err != nil {
		return respXDR, err
	}
	clearF := xdr.Uint32(clearFlag)
	setF := xdr.Uint32(setFlag)
	masterW := xdr.Uint32(masterWeight)
	lowT := xdr.Uint32(lowThreshold)
	medT := xdr.Uint32(medThreshold)
	highT := xdr.Uint32(highThreshold)
	homeD := xdr.String32(homeDomain)
	var signer xdr.Signer
	signer.Weight = xdr.Uint32(singerWeight)
	err = signer.Key.SetAddress(singerPublicKey)
	if err != nil {
		panic(err)
	}

	op := xdr.SetOptionsOp{
		InflationDest: &destination,
		ClearFlags:    &clearF,
		SetFlags:      &setF,
		MasterWeight:  &masterW,
		LowThreshold:  &lowT,
		MedThreshold:  &medT,
		HighThreshold: &highT,
		HomeDomain:    &homeD,
		Signer:        &signer,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

//Not found yet
// func AccountMerge() {
// }
