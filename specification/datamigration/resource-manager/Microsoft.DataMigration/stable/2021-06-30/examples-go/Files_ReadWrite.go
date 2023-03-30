package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_ReadWrite.json
func ExampleFilesClient_ReadWrite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().ReadWrite(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileStorageInfo = armdatamigration.FileStorageInfo{
	// 	Headers: map[string]*string{
	// 		"x-ms-blob-type": to.Ptr("BlockBlob"),
	// 	},
	// 	URI: to.Ptr("https://dmssdkservicestorage.blob.core.windows.net/dmssdkservicecontainer/_rpfiles/dmssdkproject/pfpu7fxqcpziyg2h3qj2vb7d8jpbbg7p?sv=2016-05-31&sr=b&sig=sassignature&se=2018-10-05T18%3A21%3A42Z&sp=racwd"),
	// }
}
