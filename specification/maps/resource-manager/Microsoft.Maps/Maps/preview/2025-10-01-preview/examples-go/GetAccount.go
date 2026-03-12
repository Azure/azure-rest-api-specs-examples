package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/GetAccount.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "myResourceGroup", "myMapsAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmaps.AccountsClientGetResponse{
	// 	Account: &armmaps.Account{
	// 		Name: to.Ptr("myMapsAccount"),
	// 		Type: to.Ptr("Microsoft.Maps/accounts"),
	// 		ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount"),
	// 		Kind: to.Ptr(armmaps.KindGen2),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmaps.AccountProperties{
	// 			DisableLocalAuth: to.Ptr(false),
	// 			LinkedResources: []*armmaps.LinkedResource{
	// 			},
	// 			Locations: []*armmaps.LocationsItem{
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			UniqueID: to.Ptr("string"),
	// 		},
	// 		SKU: &armmaps.SKU{
	// 			Name: to.Ptr(armmaps.NameG2),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		SystemData: &armmaps.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-02T01:01:01.1075056Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmaps.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-02T01:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmaps.CreatedByTypeApplication),
	// 		},
	// 		Tags: map[string]*string{
	// 			"test": to.Ptr("true"),
	// 		},
	// 	},
	// }
}
