package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
)

// Generated from example definition: 2025-05-01/listVault.json
func ExampleVaultsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVaultsClient().NewListPager(&armkeyvault.VaultsClientListOptions{
		Top: to.Ptr[int32](1)})
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
		// page = armkeyvault.VaultsClientListResponse{
		// 	ResourceListResult: armkeyvault.ResourceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resources?%24filter=resourceType+eq+%27Microsoft.KeyVault%2fvaults%27&%24top=4&api-version=2015-11-01&%24skiptoken=eyJuZXh0UGFydGl0aW9uS2V5IjoiMSE4IVEwTTJNVGMtIiwibmV4dFJvd0tleSI6IjEhMTMyIU5rUTRSVEU1UXpJelFUWXdORGd4UmpoRlFrSXhNVGhFUXpNd01EZEROVVJmVkVkT1RDMU5TVU5TVDFOUFJsUTZNa1ZMUlZsV1FWVk1WRG95UmxaQlZVeFVVeTFUVTFKSExVcEpUVG95UkZSRlUxUTZNa1JXUVZWTVZDMVhSVk5VVlZNLSJ9"),
		// 		Value: []*armkeyvault.TrackedResource{
		// 			{
		// 				Name: to.Ptr("sample-vault"),
		// 				Type: to.Ptr("Microsoft.KeyVault/vaults"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/vaults/sample-vault"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
