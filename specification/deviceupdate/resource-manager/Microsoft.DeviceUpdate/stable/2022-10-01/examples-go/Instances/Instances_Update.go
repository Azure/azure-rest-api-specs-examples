package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Instances/Instances_Update.json
func ExampleInstancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstancesClient().Update(ctx, "test-rg", "contoso", "blue", armdeviceupdate.TagUpdate{
		Tags: map[string]*string{
			"tagKey": to.Ptr("tagValue"),
		},
	}, nil)
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
	// 	Tags: map[string]*string{
	// 		"tagKey": to.Ptr("tagValue"),
	// 	},
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
	// 	},
	// }
}
