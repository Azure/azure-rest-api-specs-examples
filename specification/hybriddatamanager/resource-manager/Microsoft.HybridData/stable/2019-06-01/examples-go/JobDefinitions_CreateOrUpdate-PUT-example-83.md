Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybriddatamanager%2Farmhybriddatamanager%2Fv0.2.0/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybriddatamanager_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// x-ms-original-file: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_CreateOrUpdate-PUT-example-83.json
func ExampleJobDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybriddatamanager.NewJobDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<data-service-name>",
		"<job-definition-name>",
		"<resource-group-name>",
		"<data-manager-name>",
		armhybriddatamanager.JobDefinition{
			Properties: &armhybriddatamanager.JobDefinitionProperties{
				DataServiceInput: map[string]interface{}{
					"AzureStorageType": "Blob",
					"BackupChoice":     "UseExistingLatest",
					"ContainerName":    "containerfromtest",
					"DeviceName":       "8600-SHG0997877L71FC",
					"FileNameFilter":   "*",
					"IsDirectoryMode":  false,
					"RootDirectories": []interface{}{
						"\\",
					},
					"VolumeNames": []interface{}{
						"TestAutomation",
					},
				},
				DataSinkID:       to.StringPtr("<data-sink-id>"),
				DataSourceID:     to.StringPtr("<data-source-id>"),
				RunLocation:      armhybriddatamanager.RunLocationWestus.ToPtr(),
				State:            armhybriddatamanager.StateEnabled.ToPtr(),
				UserConfirmation: armhybriddatamanager.UserConfirmationRequired.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobDefinitionsClientCreateOrUpdateResult)
}
```
