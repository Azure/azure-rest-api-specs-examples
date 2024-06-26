package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListAvailabilities.json
func ExampleSKUsAvailabilityClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsAvailabilityClient().NewListPager("westus2", &armvmwarecloudsimple.SKUsAvailabilityClientListOptions{SKUID: nil})
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
		// page.SKUAvailabilityListResponse = armvmwarecloudsimple.SKUAvailabilityListResponse{
		// 	Value: []*armvmwarecloudsimple.SKUAvailability{
		// 		{
		// 			DedicatedAvailabilityZoneID: to.Ptr("az1"),
		// 			DedicatedAvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 			DedicatedPlacementGroupID: to.Ptr("n1"),
		// 			DedicatedPlacementGroupName: to.Ptr("Placement Group 1"),
		// 			Limit: to.Ptr[int32](0),
		// 			ResourceType: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			SKUID: to.Ptr("general"),
		// 			SKUName: to.Ptr("CS28-Node"),
		// 		},
		// 		{
		// 			DedicatedAvailabilityZoneID: to.Ptr("az1"),
		// 			DedicatedAvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 			DedicatedPlacementGroupID: to.Ptr("n1"),
		// 			DedicatedPlacementGroupName: to.Ptr("Placement Group 1"),
		// 			Limit: to.Ptr[int32](0),
		// 			ResourceType: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			SKUID: to.Ptr("large"),
		// 			SKUName: to.Ptr("CS36-Node"),
		// 	}},
		// }
	}
}
