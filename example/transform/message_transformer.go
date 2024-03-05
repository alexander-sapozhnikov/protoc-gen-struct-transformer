// Code generated by protoc-gen-struct-transformer, version: 1.0.7-dev. DO NOT EDIT.
// source file: example/message.proto
// source package: svc.example

package transform

import (
	"strconv"

	"github.com/alexander-sapozhnikov/protoc-gen-struct-transformer/example"
	"github.com/alexander-sapozhnikov/protoc-gen-struct-transformer/example/helpers"
	"github.com/alexander-sapozhnikov/protoc-gen-struct-transformer/example/model"
)

// Oneof: "the_decl"

// message "TheOne" has no option "transformer.go_struct", skipped...
// Oneof: "the_decl"

// message "NotSupportedOneOf" has no option "transformer.go_struct", skipped...
// Oneof: "value"

// message "CustomOneof" has no option "transformer.go_struct", skipped...
// message "CustomType" has no option "transformer.go_struct", skipped...
// field skipped: some_field
// message "SkippedMessageOne" has no option "transformer.go_struct", skipped...
// message "SkippedMessageTwo" has no option "transformer.go_struct", skipped...
func PbToProductPtr(src *example.Product, opts ...TransformParam) *model.Product {
	if src == nil {
		return nil
	}

	d := PbToProduct(*src, opts...)
	return &d
}

func PbToProductPtrList(src []*example.Product, opts ...TransformParam) []*model.Product {
	resp := make([]*model.Product, len(src))

	for i, s := range src {
		resp[i] = PbToProductPtr(s, opts...)
	}

	return resp
}

func PbToProductPtrVal(src *example.Product, opts ...TransformParam) model.Product {
	if src == nil {
		return model.Product{}
	}

	return PbToProduct(*src, opts...)
}

func PbToProductPtrValList(src []*example.Product, opts ...TransformParam) []model.Product {
	resp := make([]model.Product, len(src))

	for i, s := range src {
		resp[i] = PbToProduct(*s)
	}

	return resp
}

// PbToProductList is DEPRECATED. Use PbToProductPtrValList instead.
func PbToProductList(src []*example.Product, opts ...TransformParam) []model.Product {
	return PbToProductPtrValList(src)
}

func PbToProduct(src example.Product, opts ...TransformParam) model.Product {
	s := model.Product{
		ID:                int(src.Id),
		Name:              src.Name,
		One:               TheOneToString(src.One),
		SecondID:          TheOneToString(src.SecondId),
		CustomField:       PbCustomTypeToStringPtrVal(src.CustomField, opts...),
		CustomOneof:       PbCustomOneofToStringPtrVal(src.CustomOneof, opts...),
		NotsupportedOneof: PbToPtrVal(src.NotsupportedOneof, opts...),
	}

	applyOptions(opts...)

	return s
}

func PbToProductValPtr(src example.Product, opts ...TransformParam) *model.Product {
	d := PbToProduct(src, opts...)
	return &d
}

func PbToProductValList(src []example.Product, opts ...TransformParam) []model.Product {
	resp := make([]model.Product, len(src))

	for i, s := range src {
		resp[i] = PbToProduct(s, opts...)
	}

	return resp
}

func ProductToPbPtr(src *model.Product, opts ...TransformParam) *example.Product {
	if src == nil {
		return nil
	}

	d := ProductToPb(*src, opts...)
	return &d
}

func ProductToPbPtrList(src []*model.Product, opts ...TransformParam) []*example.Product {
	resp := make([]*example.Product, len(src))

	for i, s := range src {
		resp[i] = ProductToPbPtr(s, opts...)
	}

	return resp
}

func ProductToPbPtrVal(src *model.Product, opts ...TransformParam) example.Product {
	if src == nil {
		return example.Product{}
	}

	return ProductToPb(*src, opts...)
}

func ProductToPbValPtrList(src []model.Product, opts ...TransformParam) []*example.Product {
	resp := make([]*example.Product, len(src))

	for i, s := range src {
		g := ProductToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// ProductToPbList is DEPRECATED. Use ProductToPbValPtrList instead.
func ProductToPbList(src []model.Product, opts ...TransformParam) []*example.Product {
	return ProductToPbValPtrList(src)
}

func ProductToPb(src model.Product, opts ...TransformParam) example.Product {
	s := example.Product{
		Id:                int32(src.ID),
		Name:              src.Name,
		One:               &example.TheOne{},
		SecondId:          &example.TheOne{},
		CustomField:       StringToPbCustomTypeValPtr(src.CustomField, opts...),
		CustomOneof:       StringToPbCustomOneofValPtr(src.CustomOneof, opts...),
		NotsupportedOneof: ToPbValPtr(src.NotsupportedOneof, opts...),
	}

	applyOptions(opts...)

	StringToTheOne(src.One, s.One, version)
	StringToTheOne(src.SecondID, s.SecondId, version)

	return s
}

func ProductToPbValPtr(src model.Product, opts ...TransformParam) *example.Product {
	d := ProductToPb(src, opts...)
	return &d
}

func ProductToPbValList(src []model.Product, opts ...TransformParam) []example.Product {
	resp := make([]example.Product, len(src))

	for i, s := range src {
		resp[i] = ProductToPb(s, opts...)
	}

	return resp
}

func PbToOrderPtr(src *example.Order, opts ...TransformParam) *model.Order {
	if src == nil {
		return nil
	}

	d := PbToOrder(*src, opts...)
	return &d
}

func PbToOrderPtrList(src []*example.Order, opts ...TransformParam) []*model.Order {
	resp := make([]*model.Order, len(src))

	for i, s := range src {
		resp[i] = PbToOrderPtr(s, opts...)
	}

	return resp
}

func PbToOrderPtrVal(src *example.Order, opts ...TransformParam) model.Order {
	if src == nil {
		return model.Order{}
	}

	return PbToOrder(*src, opts...)
}

func PbToOrderPtrValList(src []*example.Order, opts ...TransformParam) []model.Order {
	resp := make([]model.Order, len(src))

	for i, s := range src {
		resp[i] = PbToOrder(*s)
	}

	return resp
}

// PbToOrderList is DEPRECATED. Use PbToOrderPtrValList instead.
func PbToOrderList(src []*example.Order, opts ...TransformParam) []model.Order {
	return PbToOrderPtrValList(src)
}

func PbToOrder(src example.Order, opts ...TransformParam) model.Order {
	s := model.Order{
		ID:       int(src.Id),
		FirstID:  TheOneToString(src.FirstId),
		SecondID: TheOneToString(src.SecondId),
		ThirdURL: TheOneToString(src.ThirdUrl),
	}

	applyOptions(opts...)

	return s
}

func PbToOrderValPtr(src example.Order, opts ...TransformParam) *model.Order {
	d := PbToOrder(src, opts...)
	return &d
}

func PbToOrderValList(src []example.Order, opts ...TransformParam) []model.Order {
	resp := make([]model.Order, len(src))

	for i, s := range src {
		resp[i] = PbToOrder(s, opts...)
	}

	return resp
}

func OrderToPbPtr(src *model.Order, opts ...TransformParam) *example.Order {
	if src == nil {
		return nil
	}

	d := OrderToPb(*src, opts...)
	return &d
}

func OrderToPbPtrList(src []*model.Order, opts ...TransformParam) []*example.Order {
	resp := make([]*example.Order, len(src))

	for i, s := range src {
		resp[i] = OrderToPbPtr(s, opts...)
	}

	return resp
}

func OrderToPbPtrVal(src *model.Order, opts ...TransformParam) example.Order {
	if src == nil {
		return example.Order{}
	}

	return OrderToPb(*src, opts...)
}

func OrderToPbValPtrList(src []model.Order, opts ...TransformParam) []*example.Order {
	resp := make([]*example.Order, len(src))

	for i, s := range src {
		g := OrderToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// OrderToPbList is DEPRECATED. Use OrderToPbValPtrList instead.
func OrderToPbList(src []model.Order, opts ...TransformParam) []*example.Order {
	return OrderToPbValPtrList(src)
}

func OrderToPb(src model.Order, opts ...TransformParam) example.Order {
	s := example.Order{
		Id:       int64(src.ID),
		FirstId:  &example.TheOne{},
		SecondId: &example.TheOne{},
		ThirdUrl: &example.TheOne{},
	}

	applyOptions(opts...)

	StringToTheOne(src.FirstID, s.FirstId, version)
	StringToTheOne(src.SecondID, s.SecondId, version)
	StringToTheOne(src.ThirdURL, s.ThirdUrl, version)
	return s
}

func OrderToPbValPtr(src model.Order, opts ...TransformParam) *example.Order {
	d := OrderToPb(src, opts...)
	return &d
}

func OrderToPbValList(src []model.Order, opts ...TransformParam) []example.Order {
	resp := make([]example.Order, len(src))

	for i, s := range src {
		resp[i] = OrderToPb(s, opts...)
	}

	return resp
}

func PbToAddressPtr(src *example.Address, opts ...TransformParam) *model.Address {
	if src == nil {
		return nil
	}

	d := PbToAddress(*src, opts...)
	return &d
}

func PbToAddressPtrList(src []*example.Address, opts ...TransformParam) []*model.Address {
	resp := make([]*model.Address, len(src))

	for i, s := range src {
		resp[i] = PbToAddressPtr(s, opts...)
	}

	return resp
}

func PbToAddressPtrVal(src *example.Address, opts ...TransformParam) model.Address {
	if src == nil {
		return model.Address{}
	}

	return PbToAddress(*src, opts...)
}

func PbToAddressPtrValList(src []*example.Address, opts ...TransformParam) []model.Address {
	resp := make([]model.Address, len(src))

	for i, s := range src {
		resp[i] = PbToAddress(*s)
	}

	return resp
}

// PbToAddressList is DEPRECATED. Use PbToAddressPtrValList instead.
func PbToAddressList(src []*example.Address, opts ...TransformParam) []model.Address {
	return PbToAddressPtrValList(src)
}

func PbToAddress(src example.Address, opts ...TransformParam) model.Address {
	s := model.Address{
		ID:   int(src.Id),
		Type: src.Type,
	}

	applyOptions(opts...)

	return s
}

func PbToAddressValPtr(src example.Address, opts ...TransformParam) *model.Address {
	d := PbToAddress(src, opts...)
	return &d
}

func PbToAddressValList(src []example.Address, opts ...TransformParam) []model.Address {
	resp := make([]model.Address, len(src))

	for i, s := range src {
		resp[i] = PbToAddress(s, opts...)
	}

	return resp
}

func AddressToPbPtr(src *model.Address, opts ...TransformParam) *example.Address {
	if src == nil {
		return nil
	}

	d := AddressToPb(*src, opts...)
	return &d
}

func AddressToPbPtrList(src []*model.Address, opts ...TransformParam) []*example.Address {
	resp := make([]*example.Address, len(src))

	for i, s := range src {
		resp[i] = AddressToPbPtr(s, opts...)
	}

	return resp
}

func AddressToPbPtrVal(src *model.Address, opts ...TransformParam) example.Address {
	if src == nil {
		return example.Address{}
	}

	return AddressToPb(*src, opts...)
}

func AddressToPbValPtrList(src []model.Address, opts ...TransformParam) []*example.Address {
	resp := make([]*example.Address, len(src))

	for i, s := range src {
		g := AddressToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// AddressToPbList is DEPRECATED. Use AddressToPbValPtrList instead.
func AddressToPbList(src []model.Address, opts ...TransformParam) []*example.Address {
	return AddressToPbValPtrList(src)
}

func AddressToPb(src model.Address, opts ...TransformParam) example.Address {
	s := example.Address{
		Id:   int64(src.ID),
		Type: src.Type,
	}

	applyOptions(opts...)

	return s
}

func AddressToPbValPtr(src model.Address, opts ...TransformParam) *example.Address {
	d := AddressToPb(src, opts...)
	return &d
}

func AddressToPbValList(src []model.Address, opts ...TransformParam) []example.Address {
	resp := make([]example.Address, len(src))

	for i, s := range src {
		resp[i] = AddressToPb(s, opts...)
	}

	return resp
}

func PbToCustomerPtr(src *example.Customer, opts ...TransformParam) *model.Customer {
	if src == nil {
		return nil
	}

	d := PbToCustomer(*src, opts...)
	return &d
}

func PbToCustomerPtrList(src []*example.Customer, opts ...TransformParam) []*model.Customer {
	resp := make([]*model.Customer, len(src))

	for i, s := range src {
		resp[i] = PbToCustomerPtr(s, opts...)
	}

	return resp
}

func PbToCustomerPtrVal(src *example.Customer, opts ...TransformParam) model.Customer {
	if src == nil {
		return model.Customer{}
	}

	return PbToCustomer(*src, opts...)
}

func PbToCustomerPtrValList(src []*example.Customer, opts ...TransformParam) []model.Customer {
	resp := make([]model.Customer, len(src))

	for i, s := range src {
		resp[i] = PbToCustomer(*s)
	}

	return resp
}

// PbToCustomerList is DEPRECATED. Use PbToCustomerPtrValList instead.
func PbToCustomerList(src []*example.Customer, opts ...TransformParam) []model.Customer {
	return PbToCustomerPtrValList(src)
}

func PbToCustomer(src example.Customer, opts ...TransformParam) model.Customer {
	s := model.Customer{
		ID:             int(src.Id),
		Name:           src.Name,
		Addresses:      PbToAddressPtrValList(src.Addresses, opts...),
		DefaultAddress: PbToAddressPtr(src.DefaultAddress, opts...),
		BillingAddress: PbToAddress(src.BillingAddress, opts...),
		MapField1:      src.MapField_1,
		MapField2:      src.MapFieldToWithoutDigits,
	}

	applyOptions(opts...)

	return s
}

func PbToCustomerValPtr(src example.Customer, opts ...TransformParam) *model.Customer {
	d := PbToCustomer(src, opts...)
	return &d
}

func PbToCustomerValList(src []example.Customer, opts ...TransformParam) []model.Customer {
	resp := make([]model.Customer, len(src))

	for i, s := range src {
		resp[i] = PbToCustomer(s, opts...)
	}

	return resp
}

func CustomerToPbPtr(src *model.Customer, opts ...TransformParam) *example.Customer {
	if src == nil {
		return nil
	}

	d := CustomerToPb(*src, opts...)
	return &d
}

func CustomerToPbPtrList(src []*model.Customer, opts ...TransformParam) []*example.Customer {
	resp := make([]*example.Customer, len(src))

	for i, s := range src {
		resp[i] = CustomerToPbPtr(s, opts...)
	}

	return resp
}

func CustomerToPbPtrVal(src *model.Customer, opts ...TransformParam) example.Customer {
	if src == nil {
		return example.Customer{}
	}

	return CustomerToPb(*src, opts...)
}

func CustomerToPbValPtrList(src []model.Customer, opts ...TransformParam) []*example.Customer {
	resp := make([]*example.Customer, len(src))

	for i, s := range src {
		g := CustomerToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// CustomerToPbList is DEPRECATED. Use CustomerToPbValPtrList instead.
func CustomerToPbList(src []model.Customer, opts ...TransformParam) []*example.Customer {
	return CustomerToPbValPtrList(src)
}

func CustomerToPb(src model.Customer, opts ...TransformParam) example.Customer {
	s := example.Customer{
		Id:                      int64(src.ID),
		Name:                    src.Name,
		Addresses:               AddressToPbValPtrList(src.Addresses, opts...),
		DefaultAddress:          AddressToPbPtr(src.DefaultAddress, opts...),
		BillingAddress:          AddressToPb(src.BillingAddress, opts...),
		MapField_1:              src.MapField1,
		MapFieldToWithoutDigits: src.MapField2,
	}

	applyOptions(opts...)

	return s
}

func CustomerToPbValPtr(src model.Customer, opts ...TransformParam) *example.Customer {
	d := CustomerToPb(src, opts...)
	return &d
}

func CustomerToPbValList(src []model.Customer, opts ...TransformParam) []example.Customer {
	resp := make([]example.Customer, len(src))

	for i, s := range src {
		resp[i] = CustomerToPb(s, opts...)
	}

	return resp
}

func PbToMyLineItemUsagePtr(src *example.LineItemUsage, opts ...TransformParam) *model.MyLineItemUsage {
	if src == nil {
		return nil
	}

	d := PbToMyLineItemUsage(*src, opts...)
	return &d
}

func PbToMyLineItemUsagePtrList(src []*example.LineItemUsage, opts ...TransformParam) []*model.MyLineItemUsage {
	resp := make([]*model.MyLineItemUsage, len(src))

	for i, s := range src {
		resp[i] = PbToMyLineItemUsagePtr(s, opts...)
	}

	return resp
}

func PbToMyLineItemUsagePtrVal(src *example.LineItemUsage, opts ...TransformParam) model.MyLineItemUsage {
	if src == nil {
		return model.MyLineItemUsage{}
	}

	return PbToMyLineItemUsage(*src, opts...)
}

func PbToMyLineItemUsagePtrValList(src []*example.LineItemUsage, opts ...TransformParam) []model.MyLineItemUsage {
	resp := make([]model.MyLineItemUsage, len(src))

	for i, s := range src {
		resp[i] = PbToMyLineItemUsage(*s)
	}

	return resp
}

// PbToMyLineItemUsageList is DEPRECATED. Use PbToMyLineItemUsagePtrValList instead.
func PbToMyLineItemUsageList(src []*example.LineItemUsage, opts ...TransformParam) []model.MyLineItemUsage {
	return PbToMyLineItemUsagePtrValList(src)
}

func PbToMyLineItemUsage(src example.LineItemUsage, opts ...TransformParam) model.MyLineItemUsage {
	s := model.MyLineItemUsage{
		Item: PbToMyLineItemPtr(src.Item, opts...),
		List: PbToMyLineItemPtrValList(src.List, opts...),
	}

	applyOptions(opts...)

	return s
}

func PbToMyLineItemUsageValPtr(src example.LineItemUsage, opts ...TransformParam) *model.MyLineItemUsage {
	d := PbToMyLineItemUsage(src, opts...)
	return &d
}

func PbToMyLineItemUsageValList(src []example.LineItemUsage, opts ...TransformParam) []model.MyLineItemUsage {
	resp := make([]model.MyLineItemUsage, len(src))

	for i, s := range src {
		resp[i] = PbToMyLineItemUsage(s, opts...)
	}

	return resp
}

func MyLineItemUsageToPbPtr(src *model.MyLineItemUsage, opts ...TransformParam) *example.LineItemUsage {
	if src == nil {
		return nil
	}

	d := MyLineItemUsageToPb(*src, opts...)
	return &d
}

func MyLineItemUsageToPbPtrList(src []*model.MyLineItemUsage, opts ...TransformParam) []*example.LineItemUsage {
	resp := make([]*example.LineItemUsage, len(src))

	for i, s := range src {
		resp[i] = MyLineItemUsageToPbPtr(s, opts...)
	}

	return resp
}

func MyLineItemUsageToPbPtrVal(src *model.MyLineItemUsage, opts ...TransformParam) example.LineItemUsage {
	if src == nil {
		return example.LineItemUsage{}
	}

	return MyLineItemUsageToPb(*src, opts...)
}

func MyLineItemUsageToPbValPtrList(src []model.MyLineItemUsage, opts ...TransformParam) []*example.LineItemUsage {
	resp := make([]*example.LineItemUsage, len(src))

	for i, s := range src {
		g := MyLineItemUsageToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// MyLineItemUsageToPbList is DEPRECATED. Use MyLineItemUsageToPbValPtrList instead.
func MyLineItemUsageToPbList(src []model.MyLineItemUsage, opts ...TransformParam) []*example.LineItemUsage {
	return MyLineItemUsageToPbValPtrList(src)
}

func MyLineItemUsageToPb(src model.MyLineItemUsage, opts ...TransformParam) example.LineItemUsage {
	s := example.LineItemUsage{
		Item: MyLineItemToPbPtr(src.Item, opts...),
		List: MyLineItemToPbValPtrList(src.List, opts...),
	}

	applyOptions(opts...)

	return s
}

func MyLineItemUsageToPbValPtr(src model.MyLineItemUsage, opts ...TransformParam) *example.LineItemUsage {
	d := MyLineItemUsageToPb(src, opts...)
	return &d
}

func MyLineItemUsageToPbValList(src []model.MyLineItemUsage, opts ...TransformParam) []example.LineItemUsage {
	resp := make([]example.LineItemUsage, len(src))

	for i, s := range src {
		resp[i] = MyLineItemUsageToPb(s, opts...)
	}

	return resp
}

func PbToMyLineItemPtr(src *example.LineItem, opts ...TransformParam) *model.MyLineItem {
	if src == nil {
		return nil
	}

	d := PbToMyLineItem(*src, opts...)
	return &d
}

func PbToMyLineItemPtrList(src []*example.LineItem, opts ...TransformParam) []*model.MyLineItem {
	resp := make([]*model.MyLineItem, len(src))

	for i, s := range src {
		resp[i] = PbToMyLineItemPtr(s, opts...)
	}

	return resp
}

func PbToMyLineItemPtrVal(src *example.LineItem, opts ...TransformParam) model.MyLineItem {
	if src == nil {
		return model.MyLineItem{}
	}

	return PbToMyLineItem(*src, opts...)
}

func PbToMyLineItemPtrValList(src []*example.LineItem, opts ...TransformParam) []model.MyLineItem {
	resp := make([]model.MyLineItem, len(src))

	for i, s := range src {
		resp[i] = PbToMyLineItem(*s)
	}

	return resp
}

// PbToMyLineItemList is DEPRECATED. Use PbToMyLineItemPtrValList instead.
func PbToMyLineItemList(src []*example.LineItem, opts ...TransformParam) []model.MyLineItem {
	return PbToMyLineItemPtrValList(src)
}

func PbToMyLineItem(src example.LineItem, opts ...TransformParam) model.MyLineItem {
	s := model.MyLineItem{
		ID:   int(src.ID),
		Type: src.Type,
		URL:  src.URL,
		SKU:  int(src.SKU),
	}

	applyOptions(opts...)

	return s
}

func PbToMyLineItemValPtr(src example.LineItem, opts ...TransformParam) *model.MyLineItem {
	d := PbToMyLineItem(src, opts...)
	return &d
}

func PbToMyLineItemValList(src []example.LineItem, opts ...TransformParam) []model.MyLineItem {
	resp := make([]model.MyLineItem, len(src))

	for i, s := range src {
		resp[i] = PbToMyLineItem(s, opts...)
	}

	return resp
}

func MyLineItemToPbPtr(src *model.MyLineItem, opts ...TransformParam) *example.LineItem {
	if src == nil {
		return nil
	}

	d := MyLineItemToPb(*src, opts...)
	return &d
}

func MyLineItemToPbPtrList(src []*model.MyLineItem, opts ...TransformParam) []*example.LineItem {
	resp := make([]*example.LineItem, len(src))

	for i, s := range src {
		resp[i] = MyLineItemToPbPtr(s, opts...)
	}

	return resp
}

func MyLineItemToPbPtrVal(src *model.MyLineItem, opts ...TransformParam) example.LineItem {
	if src == nil {
		return example.LineItem{}
	}

	return MyLineItemToPb(*src, opts...)
}

func MyLineItemToPbValPtrList(src []model.MyLineItem, opts ...TransformParam) []*example.LineItem {
	resp := make([]*example.LineItem, len(src))

	for i, s := range src {
		g := MyLineItemToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// MyLineItemToPbList is DEPRECATED. Use MyLineItemToPbValPtrList instead.
func MyLineItemToPbList(src []model.MyLineItem, opts ...TransformParam) []*example.LineItem {
	return MyLineItemToPbValPtrList(src)
}

func MyLineItemToPb(src model.MyLineItem, opts ...TransformParam) example.LineItem {
	s := example.LineItem{
		ID:   int64(src.ID),
		Type: src.Type,
		URL:  src.URL,
		SKU:  int64(src.SKU),
	}

	applyOptions(opts...)

	return s
}

func MyLineItemToPbValPtr(src model.MyLineItem, opts ...TransformParam) *example.LineItem {
	d := MyLineItemToPb(src, opts...)
	return &d
}

func MyLineItemToPbValList(src []model.MyLineItem, opts ...TransformParam) []example.LineItem {
	resp := make([]example.LineItem, len(src))

	for i, s := range src {
		resp[i] = MyLineItemToPb(s, opts...)
	}

	return resp
}

func PbToValue2PointerPtr(src *example.Value2Pointer, opts ...TransformParam) *model.Value2Pointer {
	if src == nil {
		return nil
	}

	d := PbToValue2Pointer(*src, opts...)
	return &d
}

func PbToValue2PointerPtrList(src []*example.Value2Pointer, opts ...TransformParam) []*model.Value2Pointer {
	resp := make([]*model.Value2Pointer, len(src))

	for i, s := range src {
		resp[i] = PbToValue2PointerPtr(s, opts...)
	}

	return resp
}

func PbToValue2PointerPtrVal(src *example.Value2Pointer, opts ...TransformParam) model.Value2Pointer {
	if src == nil {
		return model.Value2Pointer{}
	}

	return PbToValue2Pointer(*src, opts...)
}

func PbToValue2PointerPtrValList(src []*example.Value2Pointer, opts ...TransformParam) []model.Value2Pointer {
	resp := make([]model.Value2Pointer, len(src))

	for i, s := range src {
		resp[i] = PbToValue2Pointer(*s)
	}

	return resp
}

// PbToValue2PointerList is DEPRECATED. Use PbToValue2PointerPtrValList instead.
func PbToValue2PointerList(src []*example.Value2Pointer, opts ...TransformParam) []model.Value2Pointer {
	return PbToValue2PointerPtrValList(src)
}

func PbToValue2Pointer(src example.Value2Pointer, opts ...TransformParam) model.Value2Pointer {
	s := model.Value2Pointer{
		AddressNil: PbToAddressValPtr(src.AddressNil, opts...),
	}

	applyOptions(opts...)

	return s
}

func PbToValue2PointerValPtr(src example.Value2Pointer, opts ...TransformParam) *model.Value2Pointer {
	d := PbToValue2Pointer(src, opts...)
	return &d
}

func PbToValue2PointerValList(src []example.Value2Pointer, opts ...TransformParam) []model.Value2Pointer {
	resp := make([]model.Value2Pointer, len(src))

	for i, s := range src {
		resp[i] = PbToValue2Pointer(s, opts...)
	}

	return resp
}

func Value2PointerToPbPtr(src *model.Value2Pointer, opts ...TransformParam) *example.Value2Pointer {
	if src == nil {
		return nil
	}

	d := Value2PointerToPb(*src, opts...)
	return &d
}

func Value2PointerToPbPtrList(src []*model.Value2Pointer, opts ...TransformParam) []*example.Value2Pointer {
	resp := make([]*example.Value2Pointer, len(src))

	for i, s := range src {
		resp[i] = Value2PointerToPbPtr(s, opts...)
	}

	return resp
}

func Value2PointerToPbPtrVal(src *model.Value2Pointer, opts ...TransformParam) example.Value2Pointer {
	if src == nil {
		return example.Value2Pointer{}
	}

	return Value2PointerToPb(*src, opts...)
}

func Value2PointerToPbValPtrList(src []model.Value2Pointer, opts ...TransformParam) []*example.Value2Pointer {
	resp := make([]*example.Value2Pointer, len(src))

	for i, s := range src {
		g := Value2PointerToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// Value2PointerToPbList is DEPRECATED. Use Value2PointerToPbValPtrList instead.
func Value2PointerToPbList(src []model.Value2Pointer, opts ...TransformParam) []*example.Value2Pointer {
	return Value2PointerToPbValPtrList(src)
}

func Value2PointerToPb(src model.Value2Pointer, opts ...TransformParam) example.Value2Pointer {
	s := example.Value2Pointer{
		AddressNil: AddressToPbPtrVal(src.AddressNil, opts...),
	}

	applyOptions(opts...)

	return s
}

func Value2PointerToPbValPtr(src model.Value2Pointer, opts ...TransformParam) *example.Value2Pointer {
	d := Value2PointerToPb(src, opts...)
	return &d
}

func Value2PointerToPbValList(src []model.Value2Pointer, opts ...TransformParam) []example.Value2Pointer {
	resp := make([]example.Value2Pointer, len(src))

	for i, s := range src {
		resp[i] = Value2PointerToPb(s, opts...)
	}

	return resp
}

func PbToPointer2ValuePtr(src *example.Pointer2Value, opts ...TransformParam) *model.Pointer2Value {
	if src == nil {
		return nil
	}

	d := PbToPointer2Value(*src, opts...)
	return &d
}

func PbToPointer2ValuePtrList(src []*example.Pointer2Value, opts ...TransformParam) []*model.Pointer2Value {
	resp := make([]*model.Pointer2Value, len(src))

	for i, s := range src {
		resp[i] = PbToPointer2ValuePtr(s, opts...)
	}

	return resp
}

func PbToPointer2ValuePtrVal(src *example.Pointer2Value, opts ...TransformParam) model.Pointer2Value {
	if src == nil {
		return model.Pointer2Value{}
	}

	return PbToPointer2Value(*src, opts...)
}

func PbToPointer2ValuePtrValList(src []*example.Pointer2Value, opts ...TransformParam) []model.Pointer2Value {
	resp := make([]model.Pointer2Value, len(src))

	for i, s := range src {
		resp[i] = PbToPointer2Value(*s)
	}

	return resp
}

// PbToPointer2ValueList is DEPRECATED. Use PbToPointer2ValuePtrValList instead.
func PbToPointer2ValueList(src []*example.Pointer2Value, opts ...TransformParam) []model.Pointer2Value {
	return PbToPointer2ValuePtrValList(src)
}

func PbToPointer2Value(src example.Pointer2Value, opts ...TransformParam) model.Pointer2Value {
	s := model.Pointer2Value{
		AddressNotNil: PbToAddressPtrVal(src.AddressNotNil, opts...),
	}

	applyOptions(opts...)

	return s
}

func PbToPointer2ValueValPtr(src example.Pointer2Value, opts ...TransformParam) *model.Pointer2Value {
	d := PbToPointer2Value(src, opts...)
	return &d
}

func PbToPointer2ValueValList(src []example.Pointer2Value, opts ...TransformParam) []model.Pointer2Value {
	resp := make([]model.Pointer2Value, len(src))

	for i, s := range src {
		resp[i] = PbToPointer2Value(s, opts...)
	}

	return resp
}

func Pointer2ValueToPbPtr(src *model.Pointer2Value, opts ...TransformParam) *example.Pointer2Value {
	if src == nil {
		return nil
	}

	d := Pointer2ValueToPb(*src, opts...)
	return &d
}

func Pointer2ValueToPbPtrList(src []*model.Pointer2Value, opts ...TransformParam) []*example.Pointer2Value {
	resp := make([]*example.Pointer2Value, len(src))

	for i, s := range src {
		resp[i] = Pointer2ValueToPbPtr(s, opts...)
	}

	return resp
}

func Pointer2ValueToPbPtrVal(src *model.Pointer2Value, opts ...TransformParam) example.Pointer2Value {
	if src == nil {
		return example.Pointer2Value{}
	}

	return Pointer2ValueToPb(*src, opts...)
}

func Pointer2ValueToPbValPtrList(src []model.Pointer2Value, opts ...TransformParam) []*example.Pointer2Value {
	resp := make([]*example.Pointer2Value, len(src))

	for i, s := range src {
		g := Pointer2ValueToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// Pointer2ValueToPbList is DEPRECATED. Use Pointer2ValueToPbValPtrList instead.
func Pointer2ValueToPbList(src []model.Pointer2Value, opts ...TransformParam) []*example.Pointer2Value {
	return Pointer2ValueToPbValPtrList(src)
}

func Pointer2ValueToPb(src model.Pointer2Value, opts ...TransformParam) example.Pointer2Value {
	s := example.Pointer2Value{
		AddressNotNil: AddressToPbValPtr(src.AddressNotNil, opts...),
	}

	applyOptions(opts...)

	return s
}

func Pointer2ValueToPbValPtr(src model.Pointer2Value, opts ...TransformParam) *example.Pointer2Value {
	d := Pointer2ValueToPb(src, opts...)
	return &d
}

func Pointer2ValueToPbValList(src []model.Pointer2Value, opts ...TransformParam) []example.Pointer2Value {
	resp := make([]example.Pointer2Value, len(src))

	for i, s := range src {
		resp[i] = Pointer2ValueToPb(s, opts...)
	}

	return resp
}

func PbToTimeModelPtr(src *example.Timer, opts ...TransformParam) *model.TimeModel {
	if src == nil {
		return nil
	}

	d := PbToTimeModel(*src, opts...)
	return &d
}

func PbToTimeModelPtrList(src []*example.Timer, opts ...TransformParam) []*model.TimeModel {
	resp := make([]*model.TimeModel, len(src))

	for i, s := range src {
		resp[i] = PbToTimeModelPtr(s, opts...)
	}

	return resp
}

func PbToTimeModelPtrVal(src *example.Timer, opts ...TransformParam) model.TimeModel {
	if src == nil {
		return model.TimeModel{}
	}

	return PbToTimeModel(*src, opts...)
}

func PbToTimeModelPtrValList(src []*example.Timer, opts ...TransformParam) []model.TimeModel {
	resp := make([]model.TimeModel, len(src))

	for i, s := range src {
		resp[i] = PbToTimeModel(*s)
	}

	return resp
}

// PbToTimeModelList is DEPRECATED. Use PbToTimeModelPtrValList instead.
func PbToTimeModelList(src []*example.Timer, opts ...TransformParam) []model.TimeModel {
	return PbToTimeModelPtrValList(src)
}

func PbToTimeModel(src example.Timer, opts ...TransformParam) model.TimeModel {
	s := model.TimeModel{
		TimeTime:      src.Time,
		PtrTimeTime:   src.PtrTime,
		NullsTime:     helpers.TimeToNullsTime(src.TimeToStruct),
		PtrNullsTime:  helpers.TimePtrToNullsTimePtr(src.TimeToStructPtr),
		NullsTime2:    helpers.TimePtrToNullsTime(src.TimePtrToStruct),
		PtrNullsTime2: helpers.TimePtrToNullsTimePtr(src.TimePtrToPtrStruct),
	}

	applyOptions(opts...)

	return s
}

func PbToTimeModelValPtr(src example.Timer, opts ...TransformParam) *model.TimeModel {
	d := PbToTimeModel(src, opts...)
	return &d
}

func PbToTimeModelValList(src []example.Timer, opts ...TransformParam) []model.TimeModel {
	resp := make([]model.TimeModel, len(src))

	for i, s := range src {
		resp[i] = PbToTimeModel(s, opts...)
	}

	return resp
}

func TimeModelToPbPtr(src *model.TimeModel, opts ...TransformParam) *example.Timer {
	if src == nil {
		return nil
	}

	d := TimeModelToPb(*src, opts...)
	return &d
}

func TimeModelToPbPtrList(src []*model.TimeModel, opts ...TransformParam) []*example.Timer {
	resp := make([]*example.Timer, len(src))

	for i, s := range src {
		resp[i] = TimeModelToPbPtr(s, opts...)
	}

	return resp
}

func TimeModelToPbPtrVal(src *model.TimeModel, opts ...TransformParam) example.Timer {
	if src == nil {
		return example.Timer{}
	}

	return TimeModelToPb(*src, opts...)
}

func TimeModelToPbValPtrList(src []model.TimeModel, opts ...TransformParam) []*example.Timer {
	resp := make([]*example.Timer, len(src))

	for i, s := range src {
		g := TimeModelToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// TimeModelToPbList is DEPRECATED. Use TimeModelToPbValPtrList instead.
func TimeModelToPbList(src []model.TimeModel, opts ...TransformParam) []*example.Timer {
	return TimeModelToPbValPtrList(src)
}

func TimeModelToPb(src model.TimeModel, opts ...TransformParam) example.Timer {
	s := example.Timer{
		Time:               src.TimeTime,
		PtrTime:            src.PtrTimeTime,
		TimeToStruct:       helpers.NullsTimeToTime(src.NullsTime),
		TimeToStructPtr:    helpers.NullsTimePtrToTimePtr(src.PtrNullsTime),
		TimePtrToStruct:    helpers.NullsTimeToTimePtr(src.NullsTime2),
		TimePtrToPtrStruct: helpers.NullsTimePtrToTimePtr(src.PtrNullsTime2),
	}

	applyOptions(opts...)

	return s
}

func TimeModelToPbValPtr(src model.TimeModel, opts ...TransformParam) *example.Timer {
	d := TimeModelToPb(src, opts...)
	return &d
}

func TimeModelToPbValList(src []model.TimeModel, opts ...TransformParam) []example.Timer {
	resp := make([]example.Timer, len(src))

	for i, s := range src {
		resp[i] = TimeModelToPb(s, opts...)
	}

	return resp
}

func PbToIntsModelPtr(src *example.Ints, opts ...TransformParam) *model.IntsModel {
	if src == nil {
		return nil
	}

	d := PbToIntsModel(*src, opts...)
	return &d
}

func PbToIntsModelPtrList(src []*example.Ints, opts ...TransformParam) []*model.IntsModel {
	resp := make([]*model.IntsModel, len(src))

	for i, s := range src {
		resp[i] = PbToIntsModelPtr(s, opts...)
	}

	return resp
}

func PbToIntsModelPtrVal(src *example.Ints, opts ...TransformParam) model.IntsModel {
	if src == nil {
		return model.IntsModel{}
	}

	return PbToIntsModel(*src, opts...)
}

func PbToIntsModelPtrValList(src []*example.Ints, opts ...TransformParam) []model.IntsModel {
	resp := make([]model.IntsModel, len(src))

	for i, s := range src {
		resp[i] = PbToIntsModel(*s)
	}

	return resp
}

// PbToIntsModelList is DEPRECATED. Use PbToIntsModelPtrValList instead.
func PbToIntsModelList(src []*example.Ints, opts ...TransformParam) []model.IntsModel {
	return PbToIntsModelPtrValList(src)
}

func PbToIntsModel(src example.Ints, opts ...TransformParam) model.IntsModel {
	s := model.IntsModel{
		IntFor32Value: int(src.IntFor_32Value),
		IntFor64Value: int(src.IntFor_64Value),
		Int32Value:    src.Int32Value,
		Int64Value:    src.Int64Value,
		StringValue:   helpers.Int64ToString(src.StringValue),
	}

	applyOptions(opts...)

	return s
}

func PbToIntsModelValPtr(src example.Ints, opts ...TransformParam) *model.IntsModel {
	d := PbToIntsModel(src, opts...)
	return &d
}

func PbToIntsModelValList(src []example.Ints, opts ...TransformParam) []model.IntsModel {
	resp := make([]model.IntsModel, len(src))

	for i, s := range src {
		resp[i] = PbToIntsModel(s, opts...)
	}

	return resp
}

func IntsModelToPbPtr(src *model.IntsModel, opts ...TransformParam) *example.Ints {
	if src == nil {
		return nil
	}

	d := IntsModelToPb(*src, opts...)
	return &d
}

func IntsModelToPbPtrList(src []*model.IntsModel, opts ...TransformParam) []*example.Ints {
	resp := make([]*example.Ints, len(src))

	for i, s := range src {
		resp[i] = IntsModelToPbPtr(s, opts...)
	}

	return resp
}

func IntsModelToPbPtrVal(src *model.IntsModel, opts ...TransformParam) example.Ints {
	if src == nil {
		return example.Ints{}
	}

	return IntsModelToPb(*src, opts...)
}

func IntsModelToPbValPtrList(src []model.IntsModel, opts ...TransformParam) []*example.Ints {
	resp := make([]*example.Ints, len(src))

	for i, s := range src {
		g := IntsModelToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// IntsModelToPbList is DEPRECATED. Use IntsModelToPbValPtrList instead.
func IntsModelToPbList(src []model.IntsModel, opts ...TransformParam) []*example.Ints {
	return IntsModelToPbValPtrList(src)
}

func IntsModelToPb(src model.IntsModel, opts ...TransformParam) example.Ints {
	s := example.Ints{
		IntFor_32Value: int32(src.IntFor32Value),
		IntFor_64Value: int64(src.IntFor64Value),
		Int32Value:     src.Int32Value,
		Int64Value:     src.Int64Value,
		StringValue:    helpers.StringToInt64(src.StringValue),
	}

	applyOptions(opts...)

	return s
}

func IntsModelToPbValPtr(src model.IntsModel, opts ...TransformParam) *example.Ints {
	d := IntsModelToPb(src, opts...)
	return &d
}

func IntsModelToPbValList(src []model.IntsModel, opts ...TransformParam) []example.Ints {
	resp := make([]example.Ints, len(src))

	for i, s := range src {
		resp[i] = IntsModelToPb(s, opts...)
	}

	return resp
}

type OneofTheDecl interface {
	GetStringValue() string
	GetInt64Value() int64
}

func TheOneToString(src OneofTheDecl) string {
	if s := src.GetStringValue(); s != "" {
		return s
	}

	if i := src.GetInt64Value(); i != 0 {
		return strconv.FormatInt(i, 10)
	}

	return "<nil>"
}

func StringToTheOne(s string, dst *example.TheOne, v string) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil || v == "v2" {
		dst.TheDecl = &example.TheOne_StringValue{StringValue: s}
		return
	}

	dst.TheDecl = &example.TheOne_Int64Value{Int64Value: i}
	return
}
