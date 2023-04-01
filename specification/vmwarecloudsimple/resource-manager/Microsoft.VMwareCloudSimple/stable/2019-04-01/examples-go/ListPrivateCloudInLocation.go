package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListPrivateCloudInLocation.json
func ExamplePrivateCloudsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateCloudsClient().NewListPager("eastus", nil)
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
		// page.PrivateCloudList = armvmwarecloudsimple.PrivateCloudList{
		// 	Value: []*armvmwarecloudsimple.PrivateCloud{
		// 		{
		// 			Name: to.Ptr("myPrivateCloud"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/privateClouds"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/eastus/privateClouds/myPrivateCloud"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armvmwarecloudsimple.PrivateCloudProperties{
		// 				Type: to.Ptr("vSphere"),
		// 				AvailabilityZoneID: to.Ptr("az1"),
		// 				AvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 				ClustersNumber: to.Ptr[int32](1),
		// 				CreatedBy: to.Ptr("john.doe@cloudsimple.com"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-21T07:32:59.491Z"); return t}()),
		// 				DNSServers: []*string{
		// 					to.Ptr("100.100.100.8"),
		// 					to.Ptr("100.100.100.9")},
		// 					Expires: to.Ptr("n/a"),
		// 					NsxType: to.Ptr("Advanced"),
		// 					PlacementGroupID: to.Ptr("n2"),
		// 					PlacementGroupName: to.Ptr("Placement Group 2"),
		// 					PrivateCloudID: to.Ptr("private-cloud-id"),
		// 					State: to.Ptr("operational"),
		// 					TotalCPUCores: to.Ptr[int32](84),
		// 					TotalNodes: to.Ptr[int32](3),
		// 					TotalRAM: to.Ptr[int32](768),
		// 					TotalStorage: to.Ptr[float32](16.875),
		// 					VSphereVersion: to.Ptr("6.7u1"),
		// 					VcenterFqdn: to.Ptr("vcsa-eastus.az.cloudsimple.io"),
		// 					VcenterRefid: to.Ptr("100.100.100.6"),
		// 				},
		// 		}},
		// 	}
	}
}
