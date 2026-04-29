package armservicegroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicegroups/armservicegroups"
)

// Generated from example definition: 2024-02-01-preview/ServiceGroup_Patch.json
func ExampleManagementClient_BeginUpdateServiceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicegroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginUpdateServiceGroup(ctx, "ServiceGroup1", armservicegroups.ServiceGroup{
		Properties: &armservicegroups.ServiceGroupProperties{
			DisplayName: to.Ptr("ServiceGroup 1 Name"),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicegroups.ManagementClientUpdateServiceGroupResponse{
	// 	ServiceGroup: &armservicegroups.ServiceGroup{
	// 		Name: to.Ptr("ServiceGroup1"),
	// 		Type: to.Ptr("Microsoft.Management/serviceGroups"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/ServiceGroup1"),
	// 		Kind: to.Ptr("App"),
	// 		Properties: &armservicegroups.ServiceGroupProperties{
	// 			DisplayName: to.Ptr("ServiceGroup 1 Name"),
	// 			Parent: &armservicegroups.ParentServiceGroupProperties{
	// 				ResourceID: to.Ptr("/providers/Microsoft.Management/serviceGroups/RootGroup"),
	// 			},
	// 			ProvisioningState: to.Ptr(armservicegroups.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
