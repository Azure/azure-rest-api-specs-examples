Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappplatform%2Farmappplatform%2Fv1.0.0/sdk/resourcemanager/appplatform/armappplatform/README.md) on how to add the SDK to your project and authenticate.

```go
package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_Update.json
func ExampleAppsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewAppsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myservice",
		"myapp",
		armappplatform.AppResource{
			Identity: &armappplatform.ManagedIdentityProperties{
				Type: to.Ptr(armappplatform.ManagedIdentityTypeSystemAssigned),
			},
			Location: to.Ptr("eastus"),
			Properties: &armappplatform.AppResourceProperties{
				EnableEndToEndTLS: to.Ptr(false),
				Fqdn:              to.Ptr("myapp.mydomain.com"),
				HTTPSOnly:         to.Ptr(false),
				PersistentDisk: &armappplatform.PersistentDisk{
					MountPath: to.Ptr("/mypersistentdisk"),
					SizeInGB:  to.Ptr[int32](2),
				},
				Public: to.Ptr(true),
				TemporaryDisk: &armappplatform.TemporaryDisk{
					MountPath: to.Ptr("/mytemporarydisk"),
					SizeInGB:  to.Ptr[int32](2),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
