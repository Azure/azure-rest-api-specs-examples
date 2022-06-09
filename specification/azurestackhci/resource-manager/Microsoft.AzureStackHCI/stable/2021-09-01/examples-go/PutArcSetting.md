```go
package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2021-09-01/examples/PutArcSetting.json
func ExampleArcSettingsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurestackhci.NewArcSettingsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<arc-setting-name>",
		armazurestackhci.ArcSetting{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ArcSettingsClientCreateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fazurestackhci%2Farmazurestackhci%2Fv0.2.0/sdk/resourcemanager/azurestackhci/armazurestackhci/README.md) on how to add the SDK to your project and authenticate.
