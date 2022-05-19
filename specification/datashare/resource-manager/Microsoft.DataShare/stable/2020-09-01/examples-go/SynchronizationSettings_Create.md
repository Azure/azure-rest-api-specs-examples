Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatashare%2Farmdatashare%2Fv1.0.0/sdk/resourcemanager/datashare/armdatashare/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatashare_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/SynchronizationSettings_Create.json
func ExampleSynchronizationSettingsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewSynchronizationSettingsClient("433a8dfd-e5d5-4e77-ad86-90acdc75eb1a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"SampleResourceGroup",
		"Account1",
		"Share1",
		"Dataset1",
		&armdatashare.ScheduledSynchronizationSetting{
			Kind: to.Ptr(armdatashare.SynchronizationSettingKindScheduleBased),
			Properties: &armdatashare.ScheduledSynchronizationSettingProperties{
				RecurrenceInterval:  to.Ptr(armdatashare.RecurrenceIntervalDay),
				SynchronizationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T04:47:52.9614956Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
