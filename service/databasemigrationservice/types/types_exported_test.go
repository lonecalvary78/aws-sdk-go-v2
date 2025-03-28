// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

func ExampleDataProviderSettings_outputUsage() {
	var union types.DataProviderSettings
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.DataProviderSettingsMemberDocDbSettings:
		_ = v.Value // Value is types.DocDbDataProviderSettings

	case *types.DataProviderSettingsMemberIbmDb2LuwSettings:
		_ = v.Value // Value is types.IbmDb2LuwDataProviderSettings

	case *types.DataProviderSettingsMemberIbmDb2zOsSettings:
		_ = v.Value // Value is types.IbmDb2zOsDataProviderSettings

	case *types.DataProviderSettingsMemberMariaDbSettings:
		_ = v.Value // Value is types.MariaDbDataProviderSettings

	case *types.DataProviderSettingsMemberMicrosoftSqlServerSettings:
		_ = v.Value // Value is types.MicrosoftSqlServerDataProviderSettings

	case *types.DataProviderSettingsMemberMongoDbSettings:
		_ = v.Value // Value is types.MongoDbDataProviderSettings

	case *types.DataProviderSettingsMemberMySqlSettings:
		_ = v.Value // Value is types.MySqlDataProviderSettings

	case *types.DataProviderSettingsMemberOracleSettings:
		_ = v.Value // Value is types.OracleDataProviderSettings

	case *types.DataProviderSettingsMemberPostgreSqlSettings:
		_ = v.Value // Value is types.PostgreSqlDataProviderSettings

	case *types.DataProviderSettingsMemberRedshiftSettings:
		_ = v.Value // Value is types.RedshiftDataProviderSettings

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.OracleDataProviderSettings
var _ *types.DocDbDataProviderSettings
var _ *types.IbmDb2LuwDataProviderSettings
var _ *types.MariaDbDataProviderSettings
var _ *types.PostgreSqlDataProviderSettings
var _ *types.MongoDbDataProviderSettings
var _ *types.MicrosoftSqlServerDataProviderSettings
var _ *types.IbmDb2zOsDataProviderSettings
var _ *types.RedshiftDataProviderSettings
var _ *types.MySqlDataProviderSettings

func ExampleErrorDetails_outputUsage() {
	var union types.ErrorDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ErrorDetailsMemberDefaultErrorDetails:
		_ = v.Value // Value is types.DefaultErrorDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.DefaultErrorDetails
