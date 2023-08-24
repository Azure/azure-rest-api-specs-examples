package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9403296f0b0e112b0d8222ad05fd1d79ee10e03/specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/UpdateAccountGen2.json
func ExampleAccountsClient_Update_updateToGen2Account() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Update(ctx, "myResourceGroup", "myMapsAccount", armmaps.AccountUpdateParameters{
		Kind: to.Ptr(armmaps.KindGen2),
		SKU: &armmaps.SKU{
			Name: to.Ptr(armmaps.NameG2),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armmaps.Account{
	// 	Name: to.Ptr("myMapsAccount"),
	// 	Type: to.Ptr("Microsoft.Maps/accounts"),
	// 	ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount"),
	// 	Location: to.Ptr("eastus"),
	// 	Kind: to.Ptr(armmaps.KindGen2),
	// 	Properties: &armmaps.AccountProperties{
	// 		DisableLocalAuth: to.Ptr(false),
	// 		LinkedResources: []*armmaps.LinkedResource{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UniqueID: to.Ptr("b2e763e6-d6f3-4858-9e2b-7cf8df85c593"),
	// 	},
	// 	SKU: &armmaps.SKU{
	// 		Name: to.Ptr(armmaps.NameG2),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// 	SystemData: &armmaps.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-02T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmaps.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-02T01:01:01.1075056Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmaps.CreatedByTypeApplication),
	// 	},
	// }
}
