Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorageimportexport%2Farmstorageimportexport%2Fv0.4.0/sdk/resourcemanager/storageimportexport/armstorageimportexport/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/UpdateExportJob.json
func ExampleJobsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstorageimportexport.NewJobsClient("<subscription-id>",
		nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<job-name>",
		"<resource-group-name>",
		armstorageimportexport.UpdateJobParameters{
			Properties: &armstorageimportexport.UpdateJobParametersProperties{
				BackupDriveManifest: to.Ptr(true),
				LogLevel:            to.Ptr("<log-level>"),
				State:               to.Ptr("<state>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
