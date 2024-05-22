package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CreateOrUpdateDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseExtensionsClient().BeginCreateOrUpdate(ctx, "rg_20cbe0f0-c2d9-4522-9177-5469aad53029", "srv_1ffd1cf8-9951-47fb-807d-a9c384763849", "878e303f-1ea0-4f17-aa3d-a22ac5e9da08", "polybaseimport", armsql.DatabaseExtensions{
		Properties: &armsql.DatabaseExtensionsProperties{
			OperationMode:  to.Ptr(armsql.OperationModePolybaseImport),
			StorageKey:     to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			StorageKeyType: to.Ptr(armsql.StorageKeyTypeStorageAccessKey),
			StorageURI:     to.Ptr("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
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
	// res.ImportExportExtensionsOperationResult = armsql.ImportExportExtensionsOperationResult{
	// 	Name: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extensions"),
	// 	ID: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 	Properties: &armsql.ImportExportExtensionsOperationResultProperties{
	// 		DatabaseName: to.Ptr("878e303f-1ea0-4f17-aa3d-a22ac5e9da08"),
	// 		LastModifiedTime: to.Ptr("lastModifiedTime"),
	// 		RequestID: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 		RequestType: to.Ptr("PolybaseImport"),
	// 		ServerName: to.Ptr("srv_1ffd1cf8-9951-47fb-807d-a9c384763849"),
	// 		Status: to.Ptr("succeeded"),
	// 	},
	// }
}
