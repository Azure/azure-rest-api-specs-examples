Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybriddatamanager%2Farmhybriddatamanager%2Fv1.0.0/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_Run-POST-example-132.json
func ExampleJobDefinitionsClient_BeginRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybriddatamanager.NewJobDefinitionsClient("6e0219f5-327a-4365-904f-05eed4227ad7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRun(ctx,
		"DataTransformation",
		"jobdeffromtestcode1",
		"ResourceGroupForSDKTest",
		"TestAzureSDKOperations",
		armhybriddatamanager.RunParameters{
			CustomerSecrets: []*armhybriddatamanager.CustomerSecret{},
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
			UserConfirmation: to.Ptr(armhybriddatamanager.UserConfirmationNotRequired),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
