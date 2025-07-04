package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense/v2"
)

// Generated from example definition: 2025-04-01-preview/Workspaces_GenerateUploadUrl_MinimumSet_Gen.json
func ExampleWorkspacesClient_GenerateUploadURL_workspacesGenerateUploadUrlMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("5443A01A-5242-4950-AC1A-2DD362180254", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().GenerateUploadURL(ctx, "rgworkspaces", "E___-3", armiotfirmwaredefense.GenerateUploadURLRequest{
		FirmwareID: to.Ptr("ktnnf"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotfirmwaredefense.WorkspacesClientGenerateUploadURLResponse{
	// 	URLToken: &armiotfirmwaredefense.URLToken{
	// 	},
	// }
}
