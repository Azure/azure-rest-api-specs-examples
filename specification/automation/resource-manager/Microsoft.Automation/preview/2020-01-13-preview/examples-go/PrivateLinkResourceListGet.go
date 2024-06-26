package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateLinkResourceListGet.json
func ExamplePrivateLinkResourcesClient_NewAutomationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewAutomationPager("rg1", "ddb1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkResourceListResult = armautomation.PrivateLinkResourceListResult{
		// 	Value: []*armautomation.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("sql"),
		// 			Type: to.Ptr("Microsoft.Automation/automationAccounts/privateLinkResources"),
		// 			ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Automation/automationAccounts/ddb1/privateLinkResources/sql"),
		// 			Properties: &armautomation.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("sql"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("ddb1"),
		// 					to.Ptr("ddb1-westus")},
		// 				},
		// 		}},
		// 	}
	}
}
