package armnetworkfunction_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction/v3"
)

// Generated from example definition: 2022-11-01/CollectorPolicyCreate.json
func ExampleCollectorPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkfunction.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCollectorPoliciesClient().BeginCreateOrUpdate(ctx, "rg1", "atc", "cp1", armnetworkfunction.CollectorPolicy{
		Location: to.Ptr("West US"),
		Properties: &armnetworkfunction.CollectorPolicyPropertiesFormat{
			EmissionPolicies: []*armnetworkfunction.EmissionPoliciesPropertiesFormat{
				{
					EmissionDestinations: []*armnetworkfunction.EmissionPolicyDestination{
						{
							DestinationType: to.Ptr(armnetworkfunction.DestinationTypeAzureMonitor),
						},
					},
					EmissionType: to.Ptr(armnetworkfunction.EmissionTypeIPFIX),
				},
			},
			IngestionPolicy: &armnetworkfunction.IngestionPolicyPropertiesFormat{
				IngestionSources: []*armnetworkfunction.IngestionSourcesPropertiesFormat{
					{
						ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
						SourceType: to.Ptr(armnetworkfunction.SourceTypeResource),
					},
				},
				IngestionType: to.Ptr(armnetworkfunction.IngestionTypeIPFIX),
			},
		},
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
	// res = armnetworkfunction.CollectorPoliciesClientCreateOrUpdateResponse{
	// 	CollectorPolicy: &armnetworkfunction.CollectorPolicy{
	// 		Name: to.Ptr("cp1"),
	// 		Type: to.Ptr("Microsoft.NetworkFunction/azureTrafficCollectors/collectorPolicies"),
	// 		Etag: to.Ptr("w/\\72090554-7e3b-43f2-80ad-99a9020dcb11\\"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.NetworkFunction/AzureTrafficCollector/atc/collectorPolicies/cp1"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetworkfunction.CollectorPolicyPropertiesFormat{
	// 			EmissionPolicies: []*armnetworkfunction.EmissionPoliciesPropertiesFormat{
	// 				{
	// 					EmissionDestinations: []*armnetworkfunction.EmissionPolicyDestination{
	// 						{
	// 							DestinationType: to.Ptr(armnetworkfunction.DestinationTypeAzureMonitor),
	// 						},
	// 					},
	// 					EmissionType: to.Ptr(armnetworkfunction.EmissionTypeIPFIX),
	// 				},
	// 			},
	// 			IngestionPolicy: &armnetworkfunction.IngestionPolicyPropertiesFormat{
	// 				IngestionSources: []*armnetworkfunction.IngestionSourcesPropertiesFormat{
	// 					{
	// 						ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
	// 						SourceType: to.Ptr(armnetworkfunction.SourceTypeResource),
	// 					},
	// 				},
	// 				IngestionType: to.Ptr(armnetworkfunction.IngestionTypeIPFIX),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetworkfunction.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
