package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/GetFileImports.json
func ExampleFileImportsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileImportsClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.FileImportsClientListOptions{Filter: nil,
		Orderby:   to.Ptr("properties/createdTimeUtc desc"),
		Top:       to.Ptr[int32](1),
		SkipToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FileImportList = armsecurityinsights.FileImportList{
		// 	Value: []*armsecurityinsights.FileImport{
		// 		{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/FileImports"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/FileImports/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Properties: &armsecurityinsights.FileImportProperties{
		// 				ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
		// 				CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-25T21:02:38.8350631Z"); return t}()),
		// 				FilesValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-26T21:02:38.8350632Z"); return t}()),
		// 				ImportFile: &armsecurityinsights.FileMetadata{
		// 					DeleteStatus: to.Ptr(armsecurityinsights.DeleteStatusNotDeleted),
		// 					FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
		// 					FileName: to.Ptr("fileName.json"),
		// 					FileSize: to.Ptr[int32](5146),
		// 				},
		// 				ImportValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-24T21:02:38.8350636Z"); return t}()),
		// 				IngestedRecordCount: to.Ptr[int32](5),
		// 				IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
		// 				Source: to.Ptr("mySource"),
		// 				State: to.Ptr(armsecurityinsights.FileImportStateIngested),
		// 				TotalRecordCount: to.Ptr[int32](5),
		// 				ValidRecordCount: to.Ptr[int32](5),
		// 			},
		// 	}},
		// }
	}
}
