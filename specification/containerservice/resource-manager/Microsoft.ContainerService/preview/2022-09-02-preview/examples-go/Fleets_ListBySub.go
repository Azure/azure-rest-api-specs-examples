package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/df863270270ad5b54fa8cce71d2c33becee0c097/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-09-02-preview/examples/Fleets_ListBySub.json
func ExampleFleetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListPager(nil)
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
		// page.FleetListResult = armcontainerservice.FleetListResult{
		// 	Value: []*armcontainerservice.Fleet{
		// 		{
		// 			Name: to.Ptr("fleet-1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/fleets"),
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/fleets/fleet-1"),
		// 			SystemData: &armcontainerservice.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				CreatedBy: to.Ptr("someUser"),
		// 				CreatedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-23T05:40:40.657Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("someOtherUser"),
		// 				LastModifiedByType: to.Ptr(armcontainerservice.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"archv2": to.Ptr(""),
		// 				"tier": to.Ptr("production"),
		// 			},
		// 			Etag: to.Ptr("23ujdflewrj3="),
		// 			Properties: &armcontainerservice.FleetProperties{
		// 				HubProfile: &armcontainerservice.FleetHubProfile{
		// 					DNSPrefix: to.Ptr("dnsprefix1"),
		// 					Fqdn: to.Ptr("dnsprefix1-abcd1234.flt.eastus.azmk8s.io"),
		// 					KubernetesVersion: to.Ptr("1.22.4"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerservice.FleetProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
