package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/DataBoxEdgeDeviceGetByNameWithDataResidency.json
func ExampleDevicesClient_Get_dataBoxEdgeDeviceGetByNameWithDataResidency() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().Get(ctx, "testedgedevice", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Device = armdataboxedge.Device{
	// 	Name: to.Ptr("EdgeTestPassResource"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices"),
	// 	ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourcegroups/VmMgmtTestPass/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/EdgeTestPassResource"),
	// 	Etag: to.Ptr("W/\"datetime'2020-09-29T18%3A50%3A31.1203818Z'\"_W/\"datetime'2020-09-29T18%3A50%3A31.1343914Z'\""),
	// 	Identity: &armdataboxedge.ResourceIdentity{
	// 		Type: to.Ptr(armdataboxedge.MsiIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("b3e34fcc-ab02-4ca4-9d22-5de115419091"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// 	Kind: to.Ptr(armdataboxedge.DataBoxEdgeDeviceKindAzureStackEdge),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armdataboxedge.DeviceProperties{
	// 		ConfiguredRoleTypes: []*armdataboxedge.RoleTypes{
	// 			to.Ptr(armdataboxedge.RoleTypesCloudEdgeManagement)},
	// 			DataBoxEdgeDeviceStatus: to.Ptr(armdataboxedge.DataBoxEdgeDeviceStatusOffline),
	// 			DataResidency: &armdataboxedge.DataResidency{
	// 				Type: to.Ptr(armdataboxedge.DataResidencyTypeZoneReplication),
	// 			},
	// 			DeviceHcsVersion: to.Ptr("2.1.1361.23408"),
	// 			DeviceLocalCapacity: to.Ptr[int64](8042259),
	// 			DeviceModel: to.Ptr("Physical"),
	// 			DeviceSoftwareVersion: to.Ptr("Azure Stack Edge 2009"),
	// 			DeviceType: to.Ptr(armdataboxedge.DeviceTypeDataBoxEdgeDevice),
	// 			EdgeProfile: &armdataboxedge.EdgeProfile{
	// 				Subscription: &armdataboxedge.EdgeProfileSubscription{
	// 					ID: to.Ptr("/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourceGroups/VmMgmtTestPass/providers/Microsoft.AzureStack/edgeSubscriptions/daaac4b0-35c4-4008-bdc6-b72ca5890f16"),
	// 					RegistrationID: to.Ptr("91d8753b-af42-4908-9a5e-2a61f08b20de"),
	// 					State: to.Ptr(armdataboxedge.SubscriptionStateRegistered),
	// 				},
	// 			},
	// 			FriendlyName: to.Ptr("DBE-1D6QHQ2"),
	// 			NodeCount: to.Ptr[int32](1),
	// 			SerialNumber: to.Ptr("1D6QHQ2"),
	// 			TimeZone: to.Ptr("Pacific Standard Time"),
	// 		},
	// 		SKU: &armdataboxedge.SKUInfo{
	// 			Name: to.Ptr(armdataboxedge.SKUNameEdge),
	// 			Tier: to.Ptr(armdataboxedge.SKUTierStandard),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	}
}
