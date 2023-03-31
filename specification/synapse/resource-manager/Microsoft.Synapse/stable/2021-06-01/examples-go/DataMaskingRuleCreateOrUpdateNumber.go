package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleCreateOrUpdateNumber.json
func ExampleDataMaskingRulesClient_CreateOrUpdate_createUpdateDataMaskingRuleForNumbers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataMaskingRulesClient().CreateOrUpdate(ctx, "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", "rule1", armsynapse.DataMaskingRule{
		Properties: &armsynapse.DataMaskingRuleProperties{
			ColumnName:      to.Ptr("test1"),
			MaskingFunction: to.Ptr(armsynapse.DataMaskingFunctionNumber),
			NumberFrom:      to.Ptr("0"),
			NumberTo:        to.Ptr("2"),
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
	// res.DataMaskingRule = armsynapse.DataMaskingRule{
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/dataMaskingPolicies/rules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6852/providers/Microsoft.Synapse/workspaces/sqlcrudtest-6852/sqlPools/sqlcrudtest-331/dataMaskingPolicies/Default/rules/"),
	// 	Location: to.Ptr("Central US"),
	// 	Properties: &armsynapse.DataMaskingRuleProperties{
	// 		ColumnName: to.Ptr("test1"),
	// 		ID: to.Ptr("dbo_Table_1_test1"),
	// 		MaskingFunction: to.Ptr(armsynapse.DataMaskingFunctionNumber),
	// 		NumberFrom: to.Ptr("0"),
	// 		NumberTo: to.Ptr("2"),
	// 		RuleState: to.Ptr(armsynapse.DataMaskingRuleStateEnabled),
	// 		SchemaName: to.Ptr("dbo"),
	// 		TableName: to.Ptr("Table_1"),
	// 	},
	// }
}
