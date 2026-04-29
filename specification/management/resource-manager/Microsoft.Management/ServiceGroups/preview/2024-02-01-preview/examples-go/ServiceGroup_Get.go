package armservicegroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicegroups/armservicegroups"
)

// Generated from example definition: 2024-02-01-preview/ServiceGroup_Get.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicegroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicegroups.ClientGetResponse{
	// 	ServiceGroup: &armservicegroups.ServiceGroup{
	// 		Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
	// 		Type: to.Ptr("Microsoft.Management/serviceGroups"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/20000000-0001-0000-0000-000000000000"),
	// 		Properties: &armservicegroups.ServiceGroupProperties{
	// 			DisplayName: to.Ptr("ServiceGroup 1 Tenant 2"),
	// 			Parent: &armservicegroups.ParentServiceGroupProperties{
	// 				ResourceID: to.Ptr("/providers/Microsoft.Management/serviceGroups/RootGroup"),
	// 			},
	// 			ProvisioningState: to.Ptr(armservicegroups.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
