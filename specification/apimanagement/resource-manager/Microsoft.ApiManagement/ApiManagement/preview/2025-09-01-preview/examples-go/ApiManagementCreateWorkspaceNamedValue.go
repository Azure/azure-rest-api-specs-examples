package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateWorkspaceNamedValue.json
func ExampleWorkspaceNamedValueClient_BeginCreateOrUpdate_apiManagementCreateWorkspaceNamedValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceNamedValueClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "wks1", "testprop2", armapimanagement.NamedValueCreateContract{
		Properties: &armapimanagement.NamedValueCreateContractProperties{
			DisplayName: to.Ptr("prop3name"),
			Secret:      to.Ptr(false),
			Tags: []*string{
				to.Ptr("foo"),
				to.Ptr("bar"),
			},
			Value: to.Ptr("propValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.WorkspaceNamedValueClientCreateOrUpdateResponse{
	// 	NamedValueContract: armapimanagement.NamedValueContract{
	// 		Name: to.Ptr("testprop2"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/namedValues"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/namedValues/testprop2"),
	// 		Properties: &armapimanagement.NamedValueContractProperties{
	// 			DisplayName: to.Ptr("prop3name"),
	// 			ProvisioningState: to.Ptr("InProgress"),
	// 			Secret: to.Ptr(false),
	// 			Tags: []*string{
	// 				to.Ptr("foo"),
	// 				to.Ptr("bar"),
	// 			},
	// 			Value: to.Ptr("propValue"),
	// 		},
	// 	},
	// }
}
