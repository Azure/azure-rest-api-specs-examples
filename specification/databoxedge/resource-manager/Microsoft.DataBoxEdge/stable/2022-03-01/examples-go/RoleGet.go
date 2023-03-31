package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RoleGet.json
func ExampleRolesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRolesClient().Get(ctx, "testedgedevice", "IoTRole1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataboxedge.RolesClientGetResponse{
	// 	                            RoleClassification: &armdataboxedge.IoTRole{
	// 		Name: to.Ptr("IoTRole1"),
	// 		Type: to.Ptr("dataBoxEdgeDevices/roles"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/IoTRole1"),
	// 		Kind: to.Ptr(armdataboxedge.RoleTypesIOT),
	// 		Properties: &armdataboxedge.IoTRoleProperties{
	// 			HostPlatform: to.Ptr(armdataboxedge.PlatformTypeLinux),
	// 			IoTDeviceDetails: &armdataboxedge.IoTDeviceInfo{
	// 				Authentication: &armdataboxedge.Authentication{
	// 					SymmetricKey: &armdataboxedge.SymmetricKey{
	// 					},
	// 				},
	// 				DeviceID: to.Ptr("iotdevice"),
	// 				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
	// 			},
	// 			IoTEdgeDeviceDetails: &armdataboxedge.IoTDeviceInfo{
	// 				Authentication: &armdataboxedge.Authentication{
	// 					SymmetricKey: &armdataboxedge.SymmetricKey{
	// 					},
	// 				},
	// 				DeviceID: to.Ptr("iotEdge"),
	// 				IoTHostHub: to.Ptr("iothub.azure-devices.net"),
	// 			},
	// 			RoleStatus: to.Ptr(armdataboxedge.RoleStatusEnabled),
	// 			ShareMappings: []*armdataboxedge.MountPointMap{
	// 			},
	// 		},
	// 	},
	// 	                        }
}
