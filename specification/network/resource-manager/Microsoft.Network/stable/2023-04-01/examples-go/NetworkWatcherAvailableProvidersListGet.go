package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NetworkWatcherAvailableProvidersListGet.json
func ExampleWatchersClient_BeginListAvailableProviders() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWatchersClient().BeginListAvailableProviders(ctx, "rg1", "nw1", armnetwork.AvailableProvidersListParameters{
		AzureLocations: []*string{
			to.Ptr("West US")},
		City:    to.Ptr("seattle"),
		Country: to.Ptr("United States"),
		State:   to.Ptr("washington"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailableProvidersList = armnetwork.AvailableProvidersList{
	// 	Countries: []*armnetwork.AvailableProvidersListCountry{
	// 		{
	// 			CountryName: to.Ptr("United States"),
	// 			States: []*armnetwork.AvailableProvidersListState{
	// 				{
	// 					Cities: []*armnetwork.AvailableProvidersListCity{
	// 						{
	// 							CityName: to.Ptr("seattle"),
	// 							Providers: []*string{
	// 								to.Ptr("Comcast Cable Communications, Inc. - ASN 7922"),
	// 								to.Ptr("Comcast Cable Communications, LLC - ASN 7922"),
	// 								to.Ptr("Level 3 Communications, Inc. (GBLX) - ASN 3549"),
	// 								to.Ptr("Qwest Communications Company, LLC - ASN 209")},
	// 						}},
	// 						StateName: to.Ptr("washington"),
	// 				}},
	// 		}},
	// 	}
}
