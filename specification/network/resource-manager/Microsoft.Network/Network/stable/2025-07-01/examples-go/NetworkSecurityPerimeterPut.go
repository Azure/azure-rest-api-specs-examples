package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkSecurityPerimeterPut.json
func ExampleSecurityPerimetersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityPerimetersClient().CreateOrUpdate(ctx, "rg1", "nsp1", armnetwork.SecurityPerimeter{
		Location:   to.Ptr("location1"),
		Properties: &armnetwork.SecurityPerimeterProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.SecurityPerimetersClientCreateOrUpdateResponse{
	// 	SecurityPerimeter: armnetwork.SecurityPerimeter{
	// 		Name: to.Ptr("TestNetworkSecurityPerimeter"),
	// 		Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/TestNetworkSecurityPerimeter"),
	// 		Location: to.Ptr("East US 2 EUAP"),
	// 		Properties: &armnetwork.SecurityPerimeterProperties{
	// 			PerimeterGUID: to.Ptr("guid"),
	// 			ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armnetwork.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
	// 			CreatedBy: to.Ptr("user"),
	// 			CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user"),
	// 			LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
