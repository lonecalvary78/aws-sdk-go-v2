// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/keyspacesstreams/types"
)

func ExampleKeyspacesCellValue_outputUsage() {
	var union types.KeyspacesCellValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.KeyspacesCellValueMemberAsciiT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberBigintT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberBlobT:
		_ = v.Value // Value is []byte

	case *types.KeyspacesCellValueMemberBoolT:
		_ = v.Value // Value is bool

	case *types.KeyspacesCellValueMemberCounterT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberDateT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberDecimalT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberDoubleT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberFloatT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberInetT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberIntT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberListT:
		_ = v.Value // Value is []types.KeyspacesCell

	case *types.KeyspacesCellValueMemberMapT:
		_ = v.Value // Value is []types.KeyspacesCellMapDefinition

	case *types.KeyspacesCellValueMemberSetT:
		_ = v.Value // Value is []types.KeyspacesCell

	case *types.KeyspacesCellValueMemberSmallintT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberTextT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberTimestampT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberTimeT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberTimeuuidT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberTinyintT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberTupleT:
		_ = v.Value // Value is []types.KeyspacesCell

	case *types.KeyspacesCellValueMemberUdtT:
		_ = v.Value // Value is map[string]types.KeyspacesCell

	case *types.KeyspacesCellValueMemberUuidT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberVarcharT:
		_ = v.Value // Value is string

	case *types.KeyspacesCellValueMemberVarintT:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ map[string]types.KeyspacesCell
var _ []types.KeyspacesCell
var _ *string
var _ *bool
var _ []types.KeyspacesCellMapDefinition
var _ []byte
