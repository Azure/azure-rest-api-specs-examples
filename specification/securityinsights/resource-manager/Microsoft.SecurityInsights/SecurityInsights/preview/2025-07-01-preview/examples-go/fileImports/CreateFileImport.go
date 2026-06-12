package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/fileImports/CreateFileImport.json
func ExampleFileImportsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileImportsClient().Create(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.FileImport{
		Properties: &armsecurityinsights.FileImportProperties{
			ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
			ImportFile: &armsecurityinsights.FileMetadata{
				FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
				FileName:   to.Ptr("myFile.json"),
				FileSize:   to.Ptr[int32](4653),
			},
			IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
			Source:        to.Ptr("mySource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.FileImportsClientCreateResponse{
	// 	FileImport: armsecurityinsights.FileImport{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/FileImports"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/FileImports/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Properties: &armsecurityinsights.FileImportProperties{
	// 			ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
	// 			CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-04T20:05:59.847136Z"); return t}()),
	// 			FilesValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-05T20:05:59.8471361Z"); return t}()),
	// 			ImportFile: &armsecurityinsights.FileMetadata{
	// 				DeleteStatus: to.Ptr(armsecurityinsights.DeleteStatusNotDeleted),
	// 				FileContentURI: to.Ptr("https://sentinelimportswus2.blob.core.windows.net/78c2e51a-3cd3-4ca0-a2d4-e7effb9a05fe/43967a5e-47a7-474e-afb8-2081e9b99ca1/fileName.json?skoid=<skoid>&sktid=<sktid>&skt=2022-03-25T21%3A12%3A51Z&ske=2022-03-25T22%3A12%3A51Z&sks=b&skv=2020-10-02&sv=2020-08-04&st=2022-03-25T21%3A12%3A51Z&se=2022-03-25T22%3A12%3A51Z&sr=b&sp=c&sig=<signature>"),
	// 				FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
	// 				FileName: to.Ptr("myFile.json"),
	// 				FileSize: to.Ptr[int32](4653),
	// 			},
	// 			ImportValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-04T20:05:59.8471366Z"); return t}()),
	// 			IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
	// 			Source: to.Ptr("mySource"),
	// 			State: to.Ptr(armsecurityinsights.FileImportStateWaitingForUpload),
	// 		},
	// 	},
	// }
}
