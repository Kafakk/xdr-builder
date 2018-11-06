package main

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

func CreateAccount(destinationPublicKey string, startingBalance uint64) (respXDR string, err error) {
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

func Payment(destinationPublicKey string, asset Asset, amount uint64) (respXDR string, err error) {
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

func ManageOffer(selling Asset, buying Asset, amount uint64, priceString string) (respXDR string, err error) {

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

func CreatePassiveOffer(selling Asset, buying Asset, amount uint64, priceString string) (respXDR string, err error) {

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

	var allowTrustAsset xdr.AllowTrustOpAsset

	assetType := asset.XDRAsset.Type
	fmt.Println("asset Type: ", assetType)

	switch assetType {
	case xdr.AssetTypeAssetTypeNative:
		return respXDR, nil
	case xdr.AssetTypeAssetTypeCreditAlphanum4:
		code, ok := asset.XDRAsset.GetAlphaNum4()
		if !ok {
			return "The value is not valid AlphaNum4", err
		}
		allowTrustAsset = xdr.AllowTrustOpAsset{
			Type:       xdr.AssetTypeAssetTypeCreditAlphanum4,
			AssetCode4: &code.AssetCode,
		}
	case xdr.AssetTypeAssetTypeCreditAlphanum12:
		code, ok := asset.XDRAsset.GetAlphaNum12()
		if !ok {
			return "The value is not valid AlphaNum12", err
		}
		allowTrustAsset = xdr.AllowTrustOpAsset{
			Type:        xdr.AssetTypeAssetTypeCreditAlphanum12,
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

func PathPayment(sendAsset Asset, sendMax uint64, destinationPublicKey string, destAsset Asset, destAmount uint64, path Path) (respXDR string, err error) {

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

func ChangeTrust(asset Asset, limit uint64) (respXDR string, err error) {
	op := xdr.ChangeTrustOp{
		Line:  asset.XDRAsset,
		Limit: xdr.Int64(limit),
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func BumpSequence(sequenceNumber uint64) (respXDR string, err error) {
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

func SetOptionInflation(inflationDestinationPublicKey string) (respXDR string, err error) {
	var destination xdr.AccountId
	err = destination.SetAddress(inflationDestinationPublicKey)
	if err != nil {
		return respXDR, err
	}
	op := xdr.SetOptionsOp{
		InflationDest: &destination,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func SetOptionClearFlags(clearF uint32) (respXDR string, err error) {
	clearFlag := xdr.Uint32(clearF)
	op := xdr.SetOptionsOp{
		ClearFlags: &clearFlag,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func SetOptionSetFlags(setF uint32) (respXDR string, err error) {
	setFlag := xdr.Uint32(setF)
	op := xdr.SetOptionsOp{
		SetFlags: &setFlag,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func SetOptionMasterWeight(masterW uint32) (respXDR string, err error) {
	masterWeight := xdr.Uint32(masterW)
	op := xdr.SetOptionsOp{
		MasterWeight: &masterWeight,
	}
	respXDR, err = xdr.MarshalBase64(op)
	return respXDR, err
}

func SetOptionThreshold(lowT, mediumT, highT uint32) (respXDR string, err error) {
	lowThreshold := xdr.Uint32(lowT)
	medThreshold := xdr.Uint32(mediumT)
	highThreshold := xdr.Uint32(highT)

	op := xdr.SetOptionsOp{
		LowThreshold:  &lowThreshold,
		MedThreshold:  &medThreshold,
		HighThreshold: &highThreshold,
	}
	respXDR, err = xdr.MarshalBase64(op)
	return respXDR, err
}

func SetOptionHomeDomain(domain string) (respXDR string, err error) {

	homeDomain := xdr.String32(domain)
	op := xdr.SetOptionsOp{
		HomeDomain: &homeDomain,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}

func SetOptionSigner(singerPublicKey string, singerWeight uint32) (respXDR string, err error) {
	var signer xdr.Signer
	signer.Weight = xdr.Uint32(singerWeight)
	err = signer.Key.SetAddress(singerPublicKey)
	if err != nil {
		return respXDR, err
	}

	op := xdr.SetOptionsOp{
		Signer: &signer,
	}
	respXDR, err = xdr.MarshalBase64(op)

	return respXDR, err
}
