package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/MobileNetworkUpdateTags.json
func ExampleMobileNetworksClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMobileNetworksClient().UpdateTags(ctx, "rg1", "testMobileNetwork", armmobilenetwork.IdentityAndTagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MobileNetwork = armmobilenetwork.MobileNetwork{
	// 	Name: to.Ptr("testMobileNetwork"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/mobileNetworks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armmobilenetwork.PropertiesFormat{
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 		PublicLandMobileNetworkIdentifier: &armmobilenetwork.PlmnID{
	// 			Mcc: to.Ptr("001"),
	// 			Mnc: to.Ptr("01"),
	// 		},
	// 		PublicLandMobileNetworks: []*armmobilenetwork.PublicLandMobileNetwork{
	// 			{
	// 				Mcc: to.Ptr("001"),
	// 				Mnc: to.Ptr("01"),
	// 				HomeNetworkPublicKeys: &armmobilenetwork.PublicLandMobileNetworkHomeNetworkPublicKeys{
	// 					ProfileA: []*armmobilenetwork.HomeNetworkPublicKey{
	// 						{
	// 							ID: to.Ptr[int32](1),
	// 							URL: to.Ptr("https://contosovault.vault.azure.net/secrets/exampleHnpk"),
	// 						},
	// 						{
	// 							ID: to.Ptr[int32](2),
	// 							URL: to.Ptr("https://contosovault.vault.azure.net/secrets/exampleHnpk2/5e4876e9140e4e16bfe6e2cf92e0cbd2"),
	// 					}},
	// 					ProfileB: []*armmobilenetwork.HomeNetworkPublicKey{
	// 						{
	// 							ID: to.Ptr[int32](1),
	// 							URL: to.Ptr("https://contosovault.vault.azure.net/secrets/exampleHnpkProfileB"),
	// 					}},
	// 				},
	// 		}},
	// 	},
	// }
}
