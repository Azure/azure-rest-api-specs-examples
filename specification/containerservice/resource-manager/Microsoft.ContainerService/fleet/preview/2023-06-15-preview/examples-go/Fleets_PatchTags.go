package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/Fleets_PatchTags.json
func ExampleFleetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginUpdate(ctx, "rg1", "fleet1", armcontainerservicefleet.FleetPatch{
		Tags: map[string]*string{
			"env":  to.Ptr("prod"),
			"tier": to.Ptr("secure"),
		},
	}, &armcontainerservicefleet.FleetsClientBeginUpdateOptions{IfMatch: to.Ptr("dfjkwelr7384")})
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
	// res.Fleet = armcontainerservicefleet.Fleet{
	// 	Name: to.Ptr("fleet-1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/fleets"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
	// 	SystemData: &armcontainerservicefleet.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		CreatedBy: to.Ptr("someUser"),
	// 		CreatedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("someOtherUser"),
	// 		LastModifiedByType: to.Ptr(armcontainerservicefleet.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"env": to.Ptr("prod"),
	// 		"tier": to.Ptr("secure"),
	// 	},
	// 	ETag: to.Ptr("23ujdflewrj3="),
	// 	Properties: &armcontainerservicefleet.FleetProperties{
	// 		HubProfile: &armcontainerservicefleet.FleetHubProfile{
	// 			DNSPrefix: to.Ptr("dnsprefix1"),
	// 			Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
	// 			KubernetesVersion: to.Ptr("1.22.4"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerservicefleet.FleetProvisioningStateSucceeded),
	// 	},
	// }
}
