package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/Fleets_ListBySub.json
func ExampleFleetsClient_NewListBySubscriptionPager_listsTheFleetResourcesInASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListBySubscriptionPager(nil)
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
		// page = armcontainerservicefleet.FleetsClientListBySubscriptionResponse{
		// 	FleetListResult: armcontainerservicefleet.FleetListResult{
		// 		Value: []*armcontainerservicefleet.Fleet{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet1"),
		// 				Name: to.Ptr("fleet-1"),
		// 				Type: to.Ptr("Microsoft.ContainerService/fleets"),
		// 				SystemData: &armcontainerservicefleet.SystemData{
		// 					CreatedBy: to.Ptr("someUser"),
		// 					CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("someOtherUser"),
		// 					LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tier": to.Ptr("production"),
		// 					"archv2": to.Ptr(""),
		// 				},
		// 				Location: to.Ptr("East US"),
		// 				ETag: to.Ptr("23ujdflewrj3="),
		// 				Properties: &armcontainerservicefleet.FleetProperties{
		// 					HubProfile: &armcontainerservicefleet.FleetHubProfile{
		// 						DNSPrefix: to.Ptr("dnsprefix1"),
		// 						Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
		// 						PortalFqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
		// 						KubernetesVersion: to.Ptr("1.22.4"),
		// 						AgentProfile: &armcontainerservicefleet.AgentProfile{
		// 							VMSize: to.Ptr("Standard_DS1"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
		// 					Status: &armcontainerservicefleet.FleetStatus{
		// 						LastOperationID: to.Ptr("operation-12345"),
		// 						LastOperationError: &armcontainerservicefleet.ErrorDetail{
		// 							Code: to.Ptr("None"),
		// 							Message: to.Ptr("No error"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
