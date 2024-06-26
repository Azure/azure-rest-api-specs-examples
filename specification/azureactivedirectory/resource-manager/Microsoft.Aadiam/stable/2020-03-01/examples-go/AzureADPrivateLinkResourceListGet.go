package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateLinkResourceListGet.json
func ExamplePrivateLinkResourcesClient_NewListByPrivateLinkPolicyPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armaad.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByPrivateLinkPolicyPager("rg1", "ddb1", nil)
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
		// page.PrivateLinkResourceListResult = armaad.PrivateLinkResourceListResult{
		// 	Value: []*armaad.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("azureactivedirectory"),
		// 			Type: to.Ptr("microsoft.aadiam/privateLinkforAzureAD/privateLinkResources"),
		// 			ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/microsoft.aadiam/privateLinkForAzureAD/ddb1/privateLinkResources/azureactivedirectory"),
		// 			Properties: &armaad.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("azureactivedirectory"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("authentication"),
		// 					to.Ptr("nonauthentication")},
		// 				},
		// 		}},
		// 	}
	}
}
