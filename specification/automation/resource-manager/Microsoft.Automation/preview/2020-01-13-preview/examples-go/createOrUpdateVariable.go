package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateVariable.json
func ExampleVariableClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVariableClient().CreateOrUpdate(ctx, "rg", "sampleAccount9", "sampleVariable", armautomation.VariableCreateOrUpdateParameters{
		Name: to.Ptr("sampleVariable"),
		Properties: &armautomation.VariableCreateOrUpdateProperties{
			Description: to.Ptr("my description"),
			IsEncrypted: to.Ptr(false),
			Value:       to.Ptr("\"ComputerName.domain.com\""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Variable = armautomation.Variable{
	// 	Name: to.Ptr("sampleVariable"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/sampleAccount9/variables/sampleVariable"),
	// 	Properties: &armautomation.VariableProperties{
	// 		Description: to.Ptr("my description"),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
	// 		IsEncrypted: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T22:59:00.937Z"); return t}()),
	// 		Value: to.Ptr("\"ComputerName2.domain.com\""),
	// 	},
	// }
}
