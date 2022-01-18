package nft

import (
	"fmt"
	"regexp"
	"strings"

	sdk "plugchain-sdk-go/types"
)

const (
	ModuleName = "nft"

	DefaultStringValue = "[do-not-modify]"

	MinClassLen  = 3
	MaxClassLen  = 64
	MaxNFTURILen = 256
)

var (
	_ sdk.Msg = &MsgIssueClass{}
	_ sdk.Msg = &MsgTransferNFT{}
	_ sdk.Msg = &MsgTransferClass{}
	_ sdk.Msg = &MsgEditNFT{}
	_ sdk.Msg = &MsgIssueNFT{}
	_ sdk.Msg = &MsgBurnNFT{}

	RegexAlphaNumeric = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	RegexAlphaTop     = regexp.MustCompile(`^[a-z].*`).MatchString

	keyWords         = strings.Join([]string{"ibc", "plug"}, "|")
	regexpKeywordFmt = fmt.Sprintf("^(%s).*", keyWords)
	regexpKeyword    = regexp.MustCompile(regexpKeywordFmt).MatchString

	URIMatchWords = strings.Join([]string{"http://", "https://"}, "|")
	regexURIFmt   = fmt.Sprintf("^(%s).*", URIMatchWords)
	regexpURI     = regexp.MustCompile(regexURIFmt).MatchString
)

func (m MsgIssueClass) Route() string {
	return ModuleName
}

func (m MsgIssueClass) Type() string {
	return "issue_class"
}

func (m MsgIssueClass) ValidateBasic() error {
	if len(m.Owner) == 0 {
		return sdk.Wrapf("missing sender address")
	}

	if err := sdk.ValidateAccAddress(m.Owner); err != nil {
		return sdk.Wrap(err)
	}
	id := strings.TrimSpace(m.ID)
	if len(id) == 0 {
		return sdk.Wrapf("missing id")
	}
	return nil
}

func (m MsgIssueClass) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgIssueClass) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (m MsgTransferNFT) Route() string {
	return ModuleName
}

func (m MsgTransferNFT) Type() string {
	return "transfer_nft"
}

func (m MsgTransferNFT) ValidateBasic() error {
	if len(m.Owner) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Owner); err != nil {
		return sdk.Wrap(err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return sdk.Wrapf("invalid recipient address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdk.Wrapf("invalid sender address (%s)", err)
	}

	return ValidateNFTID(m.ID)
}

func (m MsgTransferNFT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgTransferNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (m MsgTransferClass) Route() string {
	return ModuleName
}

func (m MsgTransferClass) Type() string {
	return "transfer_class"
}

func (m MsgTransferClass) ValidateBasic() error {
	if len(m.Owner) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Owner); err != nil {
		return sdk.Wrap(err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdk.Wrapf("invalid sender address (%s)", err)
	}

	return ValidateNFTID(m.ID)
}

func (m MsgTransferClass) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgTransferClass) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (m MsgEditNFT) Route() string {
	return ModuleName
}

func (m MsgEditNFT) Type() string {
	return "edit_nft"
}

func (m MsgEditNFT) ValidateBasic() error {
	if len(m.Owner) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Owner); err != nil {
		return sdk.Wrap(err)
	}

	if err := ValidateClassID(m.ClassID); err != nil {
		return err
	}

	if err := ValidateNFTID(m.ID); err != nil {
		return err
	}

	return ValidateNFTURI(m.URI)
}

func (m MsgEditNFT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgEditNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (m MsgIssueNFT) Route() string {
	return ModuleName
}

func (m MsgIssueNFT) Type() string {
	return "issue_nft"
}

func (m MsgIssueNFT) ValidateBasic() error {
	if len(m.Owner) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Owner); err != nil {
		return sdk.Wrap(err)
	}
	if err := sdk.ValidateAccAddress(m.Recipient); err != nil {
		return sdk.Wrap(err)
	}

	if err := ValidateClassID(m.ClassID); err != nil {
		return err
	}

	if err := ValidateNFTID(m.ID); err != nil {
		return err
	}
	return ValidateNFTURI(m.URI)
}

func (m MsgIssueNFT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgIssueNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (m MsgBurnNFT) Route() string {
	return ModuleName
}

func (m MsgBurnNFT) Type() string {
	return "burn_nft"
}

func (m MsgBurnNFT) ValidateBasic() error {
	if len(m.Owner) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Owner); err != nil {
		return sdk.Wrap(err)
	}

	if err := ValidateClassID(m.ClassID); err != nil {
		return err
	}

	return ValidateNFTID(m.ID)
}

func (m MsgBurnNFT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Owner)}
}

func (o Owner) Convert() interface{} {
	var idcs []IDC
	for _, idc := range o.CollectionIDs {
		idcs = append(idcs, IDC{
			Denom:    idc.ClassID,
			TokenIDs: idc.NFTIDs,
		})
	}
	return QueryOwnerResp{
		Address: o.Address,
		IDCs:    idcs,
	}
}

func (this NFT) Convert() interface{} {
	return QueryNFTResp{
		ID:      this.ID,
		Name:    this.Name,
		URI:     this.URI,
		Data:    this.Data,
		Creator: this.Owner,
	}
}

type NFTs []NFT

func (this Class) Convert() interface{} {
	return QueryDenomResp{
		ID:             this.ID,
		Name:           this.Name,
		Schema:         this.Schema,
		Symbol:         this.Symbol,
		Owner:          this.Owner,
		MintRestricted: this.MintRestricted,
		EditRestricted: this.EditRestricted,
	}
}

type denoms []Class

func (this denoms) Convert() interface{} {
	var denoms []QueryDenomResp
	for _, denom := range this {
		denoms = append(denoms, denom.Convert().(QueryDenomResp))
	}
	return denoms
}

func (c Collection) Convert() interface{} {
	var nfts []QueryNFTResp
	for _, nft := range c.NFTs {
		nfts = append(nfts, QueryNFTResp{
			ID:      nft.ID,
			Name:    nft.Name,
			URI:     nft.URI,
			Data:    nft.Data,
			Creator: nft.Owner,
		})
	}
	return QueryCollectionResp{
		Denom: c.Class.Convert().(QueryDenomResp),
		NFTs:  nfts,
	}
}

//ValidateNFTID verify that the nftID is legal
func ValidateNFTID(nftID string) error {
	if len(nftID) < MinClassLen || len(nftID) > MaxClassLen {
		return sdk.Wrapf("invalid nft ID", "the length of nft id(%s) only accepts value [%d, %d]", nftID, MinClassLen, MaxClassLen)
	}
	if !RegexAlphaNumeric(nftID) || !RegexAlphaTop(nftID) {
		return sdk.Wrapf("invalid nft ID", "nft id(%s) only accepts alphanumeric characters, and begin with an english letter", nftID)
	}
	return nil
}

// ValidateClassID verifies whether the  parameters are legal
func ValidateClassID(classID string) error {
	if len(classID) < MinClassLen || len(classID) > MaxClassLen {
		return sdk.Wrapf("invalid denom", "the length of Class(%s) only accepts value [%d, %d]", classID, MinClassLen, MaxClassLen)
	}
	if !RegexAlphaNumeric(classID) || !RegexAlphaTop(classID) {
		return sdk.Wrapf("invalid denom", "the Class(%s) only accepts alphanumeric characters, and begin with an english letter", classID)
	}
	return ValidateKeywords(classID)
}

// ValidateKeywords checks if the given classID begins with `DenomKeywords`
func ValidateKeywords(classID string) error {
	if regexpKeyword(classID) {
		return sdk.Wrapf("invalid denom", "invalid classID: %s, can not begin with keyword: (%s)", classID, keyWords)
	}
	return nil
}

func ValidateNFTURI(uri string) error {
	if len(uri) > MaxNFTURILen {
		return sdk.Wrapf("invalid nft uri", "the length of nft uri(%s) only accepts value [0, %d]", uri, MaxNFTURILen)
	}
	if !regexpURI(uri) {
		return sdk.Wrapf("invalid nft uri", "uri begin with: (%s) ", URIMatchWords)
	}
	return nil
}
