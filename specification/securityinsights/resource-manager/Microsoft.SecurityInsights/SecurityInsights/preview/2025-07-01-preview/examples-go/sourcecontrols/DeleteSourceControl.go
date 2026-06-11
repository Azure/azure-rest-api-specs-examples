package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/sourcecontrols/DeleteSourceControl.json
func ExampleSourceControlsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("b28fbe4a-0bb1-4593-960b-061c8655a550", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourceControlsClient().Delete(ctx, "myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a", armsecurityinsights.RepositoryAccessProperties{
		Properties: &armsecurityinsights.RepositoryAccessObject{
			RepositoryAccess: &armsecurityinsights.RepositoryAccess{
				ClientID: to.Ptr("54b3c2c0-1f48-4a1c-af9f-6399c3240b73"),
				Code:     to.Ptr("939fd7c6caf754f4f41f"),
				Kind:     to.Ptr(armsecurityinsights.RepositoryAccessKindOAuth),
				State:    to.Ptr("state"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.SourceControlsClientDeleteResponse{
	// 	Warning: armsecurityinsights.Warning{
	// 		Warning: &armsecurityinsights.WarningBody{
	// 			Code: to.Ptr(armsecurityinsights.WarningCodeSourceControlWarningDeleteServicePrincipal),
	// 			Message: to.Ptr("ServicePrincipal has not been removed due to insufficient permissions."),
	// 		},
	// 	},
	// }
}
