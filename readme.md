# XDR Builder

XDR Builder use to generate Stellar operations into xdr format.
for more information about operations: https://www.stellar.org/developers/guides/concepts/list-of-operations.html

# Usage

## SetAsset
```bash
func SetAsset(code, issuerPublicKey string) (respAsset Asset, err error)
```
Generate credit asset

## SetNativeAsset
```bash
SetNativeAsset() (respAsset Asset, err error)
```
Generate native asset

## CreateAccount
```bash
CreateAccount(destinationPublicKey string, startingBalance uint64) (respXDR string, err error)
```
- destinationPublicKey: address of new account
- startingBalance: starting amount of new account

## Payment
```bash
Payment(destinationPublicKey string, asset Asset, amount uint64) (respXDR string, err error)
```
- destinationPublicKey is address of receiver
- asset: sending asset
- amount: number of tranfer amount

## ManageOffer
```bash
ManageOffer(selling Asset, buying Asset, amount uint64, priceString string) (respXDR string, err error)
```
- selling: selling asset
- buying: buying asset
- amount: number of sell amount
- priceString: rate of selling asset

## CreatePassiveOffer
```bash
CreatePassiveOffer(selling Asset, buying Asset, amount uint64, priceString string) (respXDR string, err error)
```
- selling: selling asset
- buying: buying asset
- amount: number of sell amount
- priceString: rate of selling asset


## AllowTrust
```bash
AllowTrust(trustorPublicKey string, asset Asset, authorize bool) (respXDR string, err error)
```
- trustorPublicKey: address of trustor
- asset: asset of the trustline
- authorize: flag indicating whether the trustline


## PathPayment
```bash
PathPayment(sendAsset Asset, sendMax uint64, destinationPublicKey string, destAsset Asset, destAmount uint64, path Path) (respXDR string, err error) 
```
- sendAsset: asset from source account
- sendMaX: the maximum amount of send asset
- destinationPublicKey: account of receiver
- destAsset: asset of receiver
- destAmount: amount of destination asset
- path: array of asset path (maximum is 5)

## ChangeTrust
```bash
ChangeTrust(asset Asset, limit uint64) (respXDR string, err error)
```
- asset: asset of the trustline
- limit: limit of the trustline

## BumpSequence
```bash
BumpSequence(sequenceNumber uint64) (respXDR string, err error)
```
- BumpSequence: jump sequence number of source account to this sequence number


## ManageData
```bash
ManageData(name, value string) (respXDR string, err error)
```
- name: the name of data entry that is attached to account
- value: if not present then the Name will be deleted. If present then this value will be set in the DataEntry.


## SetOptionInflation
```bash
SetOptionInflation(inflationDestinationPublicKey string) (respXDR string, err error)
```
- inflationDestinationPublicKey: account of the inflation


## SetOptionClearFlags
```bash
SetOptionClearFlags(clearF uint32) (respXDR string, err error)
```
- clearF: point to which flags to clear

## SetOptionSetFlags
```bash
SetOptionSetFlags(setF uint32) (respXDR string, err error)
```
- setF: point to which flags to set

## SetOptionMasterWeight
```bash
SetOptionMasterWeight(masterW uint32) (respXDR string, err error)
```
- masterW: weight of source account

## SetOptionThreshold
```bash
SetOptionThreshold(lowT, mediumT, highT uint32) (respXDR string, err error)
```
- lowT: a number from 0-255 representing the low threshold 
- mediumT: a number from 0-255 representing the medium threshold 
- highT: a number from 0-255 representing the high threshold 

## SetOptionHomeDomain
```bash
SetOptionHomeDomain(domain string) (respXDR string, err error)
```
- domain: home domain of the account

## SetOptionSigner
```bash
SetOptionSigner(singerPublicKey string, singerWeight uint32) (respXDR string, err error) 
```
- singerPublicKey: Add, update, or remove a signer from an account.



## Running test
```bash
make test
```
or with coverage report
```bash
make test-report
```

