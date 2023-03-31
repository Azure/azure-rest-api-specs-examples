package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDeviceGetBySubscription.json
func ExampleDevicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListBySubscriptionPager(&armdataboxedge.DevicesClientListBySubscriptionOptions{Expand: nil})
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
		// page.DeviceList = armdataboxedge.DeviceList{
		// 	Value: []*armdataboxedge.Device{
		// 		{
		// 			Name: to.Ptr("linksub01"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/abhudeda-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/linksub01"),
		// 			Etag: to.Ptr("W/\"datetime'2020-11-19T01%3A39%3A55.1270082Z'\"_W/\"datetime'2020-11-19T01%3A39%3A55.1320118Z'\""),
		// 			Identity: &armdataboxedge.ResourceIdentity{
		// 				Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("083df009-06d9-4e3c-ae72-f9249a814334"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureDataBoxGateway),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armdataboxedge.SKUInfo{
		// 				Name: to.Ptr(armdataboxedge.SKUNameGateway),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mergeazstest"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/abhudeda-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/mergeazstest"),
		// 			Etag: to.Ptr("W/\"datetime'2020-11-19T01%3A40%3A05.9246966Z'\"_W/\"datetime'2020-11-19T01%3A40%3A05.9317011Z'\""),
		// 			Identity: &armdataboxedge.ResourceIdentity{
		// 				Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("8536db30-78d2-4759-95b7-896a66e14c24"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureDataBoxGateway),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armdataboxedge.SKUInfo{
		// 				Name: to.Ptr(armdataboxedge.SKUNameGateway),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mergeazstest2"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/abhudeda-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/mergeazstest2"),
		// 			Etag: to.Ptr("W/\"datetime'2020-11-19T01%3A40%3A06.496102Z'\"_W/\"datetime'2020-11-19T01%3A40%3A06.503107Z'\""),
		// 			Identity: &armdataboxedge.ResourceIdentity{
		// 				Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("fc4c1205-35c8-4f35-abc7-f5a78945f676"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureDataBoxGateway),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armdataboxedge.SKUInfo{
		// 				Name: to.Ptr(armdataboxedge.SKUNameGateway),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}
