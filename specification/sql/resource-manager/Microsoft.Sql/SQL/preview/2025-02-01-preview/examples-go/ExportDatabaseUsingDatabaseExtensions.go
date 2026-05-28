package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ExportDatabaseUsingDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_BeginCreateOrUpdate_exportDatabaseUsingDatabaseExtension() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("0ca8cd24-0b47-4ad5-bc7e-d70e35c44adf", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseExtensionsClient().BeginCreateOrUpdate(ctx, "rg_d1ef9eae-044d-4710-ba59-b82e84ad3157", "srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c", "db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68", "Export", armsql.DatabaseExtensions{
		Properties: &armsql.DatabaseExtensionsProperties{
			AdministratorLogin:         to.Ptr("login"),
			AdministratorLoginPassword: to.Ptr("password"),
			AuthenticationType:         to.Ptr("Sql"),
			OperationMode:              to.Ptr(armsql.OperationModeExport),
			StorageKey:                 to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			StorageKeyType:             to.Ptr(armsql.StorageKeyTypeStorageAccessKey),
			StorageURI:                 to.Ptr("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
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
	// res = armsql.DatabaseExtensionsClientCreateOrUpdateResponse{
	// 	ImportExportExtensionsOperationResult: armsql.ImportExportExtensionsOperationResult{
	// 		Name: to.Ptr("Export"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/extensions"),
	// 		ID: to.Ptr("/subscriptions/0ca8cd24-0b47-4ad5-bc7e-d70e35c44adf/resourceGroups/rg_d1ef9eae-044d-4710-ba59-b82e84ad3157/providers/Microsoft.Sql/servers/srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c/databases/db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68/extensions/Export"),
	// 		Properties: &armsql.ImportExportExtensionsOperationResultProperties{
	// 			BlobURI: to.Ptr("https://teststorage.blob.core.windows.net/testcontainer/samplebacpac.bacpac"),
	// 			DatabaseName: to.Ptr("db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68"),
	// 			LastModifiedTime: to.Ptr("lastModifiedTime"),
	// 			PrivateEndpointConnections: []*armsql.PrivateEndpointConnectionRequestStatus{
	// 			},
	// 			QueuedTime: to.Ptr("09/10/2021 18:35:10"),
	// 			RequestID: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 			RequestType: to.Ptr("Export"),
	// 			ServerName: to.Ptr("srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c"),
	// 			Status: to.Ptr("Completed"),
	// 		},
	// 	},
	// }
}
