package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/Instances_Get.json
func ExampleInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstancesClient().Get(ctx, "test-rg", "contoso", "blue", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Instance = armdeviceupdate.Instance{
	// 	Name: to.Ptr("blue"),
	// 	Type: to.Ptr("Microsoft.DeviceUpdate/accounts/instances"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/instances/blue"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armdeviceupdate.InstanceProperties{
	// 		AccountName: to.Ptr("contoso"),
	// 		DiagnosticStorageProperties: &armdeviceupdate.DiagnosticStorageProperties{
	// 			AuthenticationType: to.Ptr(armdeviceupdate.AuthenticationTypeKeyBased),
	// 			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/adu-resource-group/providers/Microsoft.Storage/storageAccounts/testAccount"),
	// 		},
	// 		EnableDiagnostics: to.Ptr(false),
	// 		IotHubs: []*armdeviceupdate.IotHubSettings{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Devices/IotHubs/blue-contoso-hub"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armdeviceupdate.ProvisioningStateSucceeded),
	// 	},
	// }
}
