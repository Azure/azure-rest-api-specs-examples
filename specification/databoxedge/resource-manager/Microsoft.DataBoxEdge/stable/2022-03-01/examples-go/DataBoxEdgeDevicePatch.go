package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDevicePatch.json
func ExampleDevicesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Update(ctx, "testedgedevice", "GroupForEdgeAutomation", armdataboxedge.DevicePatch{
		Properties: &armdataboxedge.DevicePropertiesPatch{
			EdgeProfile: &armdataboxedge.EdgeProfilePatch{
				Subscription: &armdataboxedge.EdgeProfileSubscriptionPatch{
					ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourceGroups/rapvs-rg/providers/Microsoft.AzureStack/linkedSubscriptions/ca014ddc-5cf2-45f8-b390-e901e4a0ae87"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armdataboxedge.Device{
	// 	Name: to.Ptr("testedgedevice"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourcegroups/VmMgmtTestPass/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice"),
	// 	Etag: to.Ptr("W/\"datetime'2020-11-19T04%3A43%3A38.6457308Z'\"_W/\"datetime'2020-11-19T04%3A43%3A38.6507339Z'\""),
	// 	Identity: &armdataboxedge.ResourceIdentity{
	// 		Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("b3e34fcc-ab02-4ca4-9d22-5de115419091"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureStackEdge),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armdataboxedge.DeviceProperties{
	// 		DataBoxEdgeDeviceStatus: to.Ptr(armdataboxedge.DataBoxEdgeDeviceStatusReadyToSetup),
	// 		DeviceLocalCapacity: to.Ptr[int64](0),
	// 		DeviceType: to.Ptr(armdataboxedge.DeviceTypeDataBoxEdgeDevice),
	// 		EdgeProfile: &armdataboxedge.EdgeProfile{
	// 			Subscription: &armdataboxedge.EdgeProfileSubscription{
	// 				ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/VmMgmtTestPass/providers/Microsoft.AzureStack/edgeSubscriptions/daaac4b0-35c4-4008-bdc6-b72ca5890f16"),
	// 				RegistrationID: to.Ptr("91d8753b-af42-4908-9a5e-2a61f08b20de"),
	// 				State: to.Ptr(armdataboxedge.SubscriptionStateRegistered),
	// 			},
	// 		},
	// 		NodeCount: to.Ptr[int32](0),
	// 		ResourceMoveDetails: &armdataboxedge.ResourceMoveDetails{
	// 			OperationInProgress: to.Ptr(armdataboxedge.ResourceMoveStatusNone),
	// 		},
	// 		TimeZone: to.Ptr("Pacific Standard Time"),
	// 	},
	// 	SKU: &armdataboxedge.SKUInfo{
	// 		Name: to.Ptr(armdataboxedge.SKUNameEdge),
	// 		Tier: to.Ptr(armdataboxedge.SKUTierStandard),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// }
}
