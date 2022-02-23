Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleCreateOrUpdateDefaultMax.json
func ExampleDataMaskingRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewDataMaskingRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		"<data-masking-rule-name>",
		armsynapse.DataMaskingRule{
			Properties: &armsynapse.DataMaskingRuleProperties{
				AliasName:       to.StringPtr("<alias-name>"),
				ColumnName:      to.StringPtr("<column-name>"),
				MaskingFunction: armsynapse.DataMaskingFunctionDefault.ToPtr(),
				RuleState:       armsynapse.DataMaskingRuleStateEnabled.ToPtr(),
				SchemaName:      to.StringPtr("<schema-name>"),
				TableName:       to.StringPtr("<table-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataMaskingRulesClientCreateOrUpdateResult)
}
```
