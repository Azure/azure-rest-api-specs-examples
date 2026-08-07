package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/ChatModelDeployments_Update_MaximumSet_Gen.json
func ExampleChatModelDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewChatModelDeploymentsClient().BeginUpdate(ctx, "rgdiscovery", "308882f04c9bcf36d5", "985d52cec9acb72ebe", armdiscovery.ChatModelDeployment{
		Properties: &armdiscovery.ChatModelDeploymentProperties{
			Capacity: to.Ptr[int32](15),
		},
		Tags: map[string]*string{
			"key5692": to.Ptr("gayu"),
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
	// res = armdiscovery.ChatModelDeploymentsClientUpdateResponse{
	// 	ChatModelDeployment: armdiscovery.ChatModelDeployment{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/workspaces/308882f04c9bcf36d5/chatModelDeployments/985d52cec9acb72ebe"),
	// 		Name: to.Ptr("985d52cec9acb72ebe"),
	// 		Tags: map[string]*string{
	// 			"key984": to.Ptr("sqzgsgykyhltqwmpgvhlyp"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/workspaces/chatModelDeployments"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.ChatModelDeploymentProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			ModelFormat: to.Ptr("zo"),
	// 			ModelName: to.Ptr("ijzwlirrkr"),
	// 			ModelVersion: to.Ptr("seiduxog"),
	// 			SKUName: to.Ptr("dymgademiauwwacz"),
	// 			Capacity: to.Ptr[int32](8),
	// 		},
	// 	},
	// }
}
