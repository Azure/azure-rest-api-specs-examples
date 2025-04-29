package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/Fleets_CreateOrUpdate.json
func ExampleFleetsClient_BeginCreate_createsAFleetResourceWithALongRunningOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginCreate(ctx, "rg1", "fleet1", armcontainerservicefleet.Fleet{
		Tags: map[string]*string{
			"tier":   to.Ptr("production"),
			"archv2": to.Ptr(""),
		},
		Location: to.Ptr("East US"),
		Properties: &armcontainerservicefleet.FleetProperties{
			HubProfile: &armcontainerservicefleet.FleetHubProfile{
				DNSPrefix: to.Ptr("dnsprefix1"),
				AgentProfile: &armcontainerservicefleet.AgentProfile{
					VMSize: to.Ptr("Standard_DS1"),
				},
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
	// res = armcontainerservicefleet.FleetsClientCreateResponse{
	// 	Fleet: &armcontainerservicefleet.Fleet{
	// 		ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
	// 		Name: to.Ptr("fleet-1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/fleets"),
	// 		SystemData: &armcontainerservicefleet.SystemData{
	// 			CreatedBy: to.Ptr("someUser"),
	// 			CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("someOtherUser"),
	// 			LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tier": to.Ptr("production"),
	// 			"archv2": to.Ptr(""),
	// 		},
	// 		Location: to.Ptr("East US"),
	// 		ETag: to.Ptr("23ujdflewrj3="),
	// 		Properties: &armcontainerservicefleet.FleetProperties{
	// 			HubProfile: &armcontainerservicefleet.FleetHubProfile{
	// 				DNSPrefix: to.Ptr("dnsprefix1"),
	// 				Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 				PortalFqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 				KubernetesVersion: to.Ptr("1.22.4"),
	// 				AgentProfile: &armcontainerservicefleet.AgentProfile{
	// 					VMSize: to.Ptr("Standard_DS1"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateUpdating),
	// 			Status: &armcontainerservicefleet.FleetStatus{
	// 				LastOperationID: to.Ptr("operation-12345"),
	// 				LastOperationError: &armcontainerservicefleet.ErrorDetail{
	// 					Code: to.Ptr("None"),
	// 					Message: to.Ptr("No error"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
