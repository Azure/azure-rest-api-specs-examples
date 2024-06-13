package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetProcessEntityById.json
func ExampleEntitiesClient_Get_getAProcessEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "myRg", "myWorkspace", "7264685c-038c-42c6-948c-38e14ef1fb98", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntitiesClientGetResponse{
	// 	                            EntityClassification: &armsecurityinsights.ProcessEntity{
	// 		Name: to.Ptr("7264685c-038c-42c6-948c-38e14ef1fb98"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/7264685c-038c-42c6-948c-38e14ef1fb98"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindProcess),
	// 		Properties: &armsecurityinsights.ProcessEntityProperties{
	// 			FriendlyName: to.Ptr("cmd.exe"),
	// 			CommandLine: to.Ptr("\"cmd\""),
	// 			ImageFileEntityID: to.Ptr("bba7b47b-c1c1-4021-b568-5b07b9292f5e"),
	// 			ProcessID: to.Ptr("0x2aa48"),
	// 		},
	// 	},
	// 	                        }
}
