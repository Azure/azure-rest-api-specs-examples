package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_Get.json
func ExampleFilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().Get(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectFile = armdatamigration.ProjectFile{
	// 	Name: to.Ptr("x114d023d8"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects/files"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/files/x114d023d8"),
	// 	Properties: &armdatamigration.ProjectFileProperties{
	// 		Extension: to.Ptr("sql"),
	// 		FilePath: to.Ptr("SchemaInput/DmsSdkFile.sql"),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T20:01:33.000Z"); return t}()),
	// 		MediaType: to.Ptr("text/plain"),
	// 		Size: to.Ptr[int64](51835),
	// 	},
	// }
}
