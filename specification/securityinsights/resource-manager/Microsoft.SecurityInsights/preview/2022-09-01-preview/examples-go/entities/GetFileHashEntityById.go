package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetFileHashEntityById.json
func ExampleEntitiesClient_Get_getAFileHashEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "myRg", "myWorkspace", "ea359fa6-c1e5-f878-e105-6344f3e399a1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntitiesClientGetResponse{
	// 	                            EntityClassification: &armsecurityinsights.FileHashEntity{
	// 		Name: to.Ptr("ea359fa6-c1e5-f878-e105-6344f3e399a1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/ea359fa6-c1e5-f878-e105-6344f3e399a1"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindFileHash),
	// 		Properties: &armsecurityinsights.FileHashEntityProperties{
	// 			FriendlyName: to.Ptr("E923636F1093C414AAB39F846E9D7A372BEEFA7B628B28179197E539C56AA0F0(SHA256)"),
	// 			Algorithm: to.Ptr(armsecurityinsights.FileHashAlgorithmSHA256),
	// 			HashValue: to.Ptr("E923636F1093C414AAB39F846E9D7A372BEEFA7B628B28179197E539C56AA0F0"),
	// 		},
	// 	},
	// 	                        }
}
