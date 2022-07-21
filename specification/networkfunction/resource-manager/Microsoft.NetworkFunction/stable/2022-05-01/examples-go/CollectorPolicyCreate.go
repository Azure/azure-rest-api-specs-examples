package armnetworkfunction_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/CollectorPolicyCreate.json
func ExampleCollectorPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetworkfunction.NewCollectorPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"atc",
		"cp1",
		armnetworkfunction.CollectorPolicy{
			Properties: &armnetworkfunction.CollectorPolicyPropertiesFormat{
				EmissionPolicies: []*armnetworkfunction.EmissionPoliciesPropertiesFormat{
					{
						EmissionDestinations: []*armnetworkfunction.EmissionPolicyDestination{
							{
								DestinationType: to.Ptr(armnetworkfunction.DestinationTypeAzureMonitor),
							}},
						EmissionType: to.Ptr(armnetworkfunction.EmissionTypeIPFIX),
					}},
				IngestionPolicy: &armnetworkfunction.IngestionPolicyPropertiesFormat{
					IngestionSources: []*armnetworkfunction.IngestionSourcesPropertiesFormat{
						{
							ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName"),
							SourceType: to.Ptr(armnetworkfunction.SourceTypeResource),
						}},
					IngestionType: to.Ptr(armnetworkfunction.IngestionTypeIPFIX),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
