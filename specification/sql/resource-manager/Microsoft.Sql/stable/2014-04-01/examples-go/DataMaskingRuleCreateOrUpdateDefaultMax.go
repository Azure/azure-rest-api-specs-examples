package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingRuleCreateOrUpdateDefaultMax.json
func ExampleDataMaskingRulesClient_CreateOrUpdate_createUpdateDataMaskingRuleForDefaultMax() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataMaskingRulesClient().CreateOrUpdate(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", "rule1", armsql.DataMaskingRule{
		Properties: &armsql.DataMaskingRuleProperties{
			AliasName:       to.Ptr("nickname"),
			ColumnName:      to.Ptr("test1"),
			MaskingFunction: to.Ptr(armsql.DataMaskingFunctionDefault),
			RuleState:       to.Ptr(armsql.DataMaskingRuleStateEnabled),
			SchemaName:      to.Ptr("dbo"),
			TableName:       to.Ptr("Table_1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataMaskingRule = armsql.DataMaskingRule{
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/dataMaskingPolicies/rules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Sql/servers/sqlcrudtest-6852/databases/sqlcrudtest-331/dataMaskingPolicies/Default/rules/"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armsql.DataMaskingRuleProperties{
	// 		AliasName: to.Ptr("nickname"),
	// 		ColumnName: to.Ptr("test1"),
	// 		ID: to.Ptr("dbo_Table_1_test1"),
	// 		MaskingFunction: to.Ptr(armsql.DataMaskingFunctionDefault),
	// 		RuleState: to.Ptr(armsql.DataMaskingRuleStateEnabled),
	// 		SchemaName: to.Ptr("dbo"),
	// 		TableName: to.Ptr("Table_1"),
	// 	},
	// }
}
