package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/emailServices/listByResourceGroup.json
func ExampleEmailServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailServicesClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page.EmailServiceResourceList = armcommunication.EmailServiceResourceList{
		// 	Value: []*armcommunication.EmailServiceResource{
		// 		{
		// 			Name: to.Ptr("MyEmailServiceResource"),
		// 			Type: to.Ptr("Microsoft.Communication/EmailServices"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource"),
		// 			Location: to.Ptr("Global"),
		// 			Properties: &armcommunication.EmailServiceProperties{
		// 				DataLocation: to.Ptr("United States"),
		// 				ProvisioningState: to.Ptr(armcommunication.EmailServicesProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
