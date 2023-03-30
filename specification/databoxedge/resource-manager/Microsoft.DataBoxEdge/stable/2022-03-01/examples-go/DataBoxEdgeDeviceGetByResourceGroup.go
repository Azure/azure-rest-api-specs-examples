package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDeviceGetByResourceGroup.json
func ExampleDevicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDevicesClient().NewListByResourceGroupPager("GroupForEdgeAutomation", &armdataboxedge.DevicesClientListByResourceGroupOptions{Expand: nil})
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
		// 			Name: to.Ptr("EdgeTestPassResource"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/VmMgmtTestPass/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/EdgeTestPassResource"),
		// 			Etag: to.Ptr("W/\"datetime'2020-09-29T18%3A50%3A31.1203818Z'\"_W/\"datetime'2020-09-29T18%3A50%3A31.1343914Z'\""),
		// 			Identity: &armdataboxedge.ResourceIdentity{
		// 				Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("b3e34fcc-ab02-4ca4-9d22-5de115419091"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureStackEdge),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armdataboxedge.SKUInfo{
		// 				Name: to.Ptr(armdataboxedge.SKUNameEdge),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("FPGAResource"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/VmMgmtTestPass/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/FPGAResource"),
		// 			Etag: to.Ptr("W/\"datetime'2020-11-18T23%3A31%3A30.3266766Z'\"_W/\"datetime'2020-11-18T23%3A31%3A30.3326804Z'\""),
		// 			Identity: &armdataboxedge.ResourceIdentity{
		// 				Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("d97a6de5-f5c0-485a-8f5e-b7f705d2dbc4"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureStackEdge),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armdataboxedge.SKUInfo{
		// 				Name: to.Ptr(armdataboxedge.SKUNameEdge),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestVMEdgeResource"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
		// 			ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/VmMgmtTestPass/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestVMEdgeResource"),
		// 			Etag: to.Ptr("W/\"datetime'2020-11-18T23%3A31%3A23.1715672Z'\"_W/\"datetime'2020-11-18T23%3A31%3A23.2396152Z'\""),
		// 			Identity: &armdataboxedge.ResourceIdentity{
		// 				Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("6f2b341b-aded-4ec2-a1bd-d09438d6cc8f"),
		// 				TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 			},
		// 			Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureStackEdge),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armdataboxedge.SKUInfo{
		// 				Name: to.Ptr(armdataboxedge.SKUNameEdge),
		// 				Tier: to.Ptr(armdataboxedge.SKUTierStandard),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}
