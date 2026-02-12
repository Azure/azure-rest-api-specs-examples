package armdisconnectedoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/disconnectedoperations/armdisconnectedoperations"
)

// Generated from example definition: 2025-06-01-preview/DisconnectedOperations_ListDeploymentManifest_MaximumSet_Gen.json
func ExampleClient_ListDeploymentManifest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdisconnectedoperations.NewClientFactory("301DCB09-82EC-4777-A56C-6FFF26BCC814", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListDeploymentManifest(ctx, "rgdisconnectedoperations", "demo-resource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdisconnectedoperations.ClientListDeploymentManifestResponse{
	// 	DisconnectedOperationDeploymentManifest: &armdisconnectedoperations.DisconnectedOperationDeploymentManifest{
	// 		ResourceID: to.Ptr("/subscriptions/1F6CACA0-5FFA-47AD-94FD-42368F71E49E/resourceGroups/rgdisconnectedOperations/providers/Microsoft.Edge/disconnectedOperations/demo-resource"),
	// 		ResourceName: to.Ptr("demo-resource"),
	// 		StampID: to.Ptr("401ECB09-83EC-4777-A56C-6FFF26BCC815"),
	// 		Location: to.Ptr("eastus"),
	// 		BillingModel: to.Ptr(armdisconnectedoperations.BillingModelCapacity),
	// 		ConnectionIntent: to.Ptr(armdisconnectedoperations.ConnectionIntentConnected),
	// 		Cloud: to.Ptr("Public"),
	// 	},
	// }
}
