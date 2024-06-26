package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/MobileNetworkListByResourceGroup.json
func ExampleMobileNetworksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMobileNetworksClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ListResult = armmobilenetwork.ListResult{
		// 	Value: []*armmobilenetwork.MobileNetwork{
		// 		{
		// 			Name: to.Ptr("testMobileNetwork"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/mobileNetworks"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
		// 			SystemData: &armmobilenetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmobilenetwork.PropertiesFormat{
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 				PublicLandMobileNetworkIdentifier: &armmobilenetwork.PlmnID{
		// 					Mcc: to.Ptr("001"),
		// 					Mnc: to.Ptr("01"),
		// 				},
		// 				PublicLandMobileNetworks: []*armmobilenetwork.PublicLandMobileNetwork{
		// 					{
		// 						Mcc: to.Ptr("001"),
		// 						Mnc: to.Ptr("01"),
		// 						HomeNetworkPublicKeys: &armmobilenetwork.PublicLandMobileNetworkHomeNetworkPublicKeys{
		// 							ProfileA: []*armmobilenetwork.HomeNetworkPublicKey{
		// 								{
		// 									ID: to.Ptr[int32](1),
		// 									URL: to.Ptr("https://contosovault.vault.azure.net/secrets/exampleHnpk"),
		// 								},
		// 								{
		// 									ID: to.Ptr[int32](2),
		// 									URL: to.Ptr("https://contosovault.vault.azure.net/secrets/exampleHnpk2/5e4876e9140e4e16bfe6e2cf92e0cbd2"),
		// 							}},
		// 							ProfileB: []*armmobilenetwork.HomeNetworkPublicKey{
		// 								{
		// 									ID: to.Ptr[int32](1),
		// 									URL: to.Ptr("https://contosovault.vault.azure.net/secrets/exampleHnpkProfileB"),
		// 							}},
		// 						},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
