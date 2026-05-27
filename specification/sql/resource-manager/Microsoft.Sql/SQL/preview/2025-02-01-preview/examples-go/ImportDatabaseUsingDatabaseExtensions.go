package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ImportDatabaseUsingDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_BeginCreateOrUpdate_importDatabaseUsingDatabaseExtension() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("17ca4d13-bf7d-4c33-a60e-b87a2820a325", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseExtensionsClient().BeginCreateOrUpdate(ctx, "rg_062866bf-c4f4-41f9-abf0-b59132ca7924", "srv_2d6be2d2-26c8-4930-8fb6-82a5e95e0e82", "db_2a47e946-e414-4c00-94ac-ed886bb78302", "Import", armsql.DatabaseExtensions{
		Properties: &armsql.DatabaseExtensionsProperties{
			AdministratorLogin:         to.Ptr("login"),
			AdministratorLoginPassword: to.Ptr("password"),
			AuthenticationType:         to.Ptr("Sql"),
			OperationMode:              to.Ptr(armsql.OperationModeImport),
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
	// 		Name: to.Ptr("Import"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases/extensions"),
	// 		ID: to.Ptr("/subscriptions/17ca4d13-bf7d-4c33-a60e-b87a2820a325/resourceGroups/rg_062866bf-c4f4-41f9-abf0-b59132ca7924/providers/Microsoft.Sql/servers/srv_2d6be2d2-26c8-4930-8fb6-82a5e95e0e82/databases/db_2a47e946-e414-4c00-94ac-ed886bb78302/extensions/Import"),
	// 		Properties: &armsql.ImportExportExtensionsOperationResultProperties{
	// 			BlobURI: to.Ptr("https://teststorage.blob.core.windows.net/testcontainer/samplebacpac.bacpac"),
	// 			DatabaseName: to.Ptr("db_2a47e946-e414-4c00-94ac-ed886bb78302"),
	// 			LastModifiedTime: to.Ptr("lastModifiedTime"),
	// 			PrivateEndpointConnections: []*armsql.PrivateEndpointConnectionRequestStatus{
	// 			},
	// 			QueuedTime: to.Ptr("09/10/2021 18:35:10"),
	// 			RequestID: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 			RequestType: to.Ptr("Import"),
	// 			ServerName: to.Ptr("srv_2d6be2d2-26c8-4930-8fb6-82a5e95e0e82"),
	// 			Status: to.Ptr("Completed"),
	// 		},
	// 	},
	// }
}
