Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdesktopvirtualization%2Farmdesktopvirtualization%2Fv0.2.0/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```go
package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Application_Create.json
func ExampleApplicationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewApplicationsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<application-group-name>",
		"<application-name>",
		armdesktopvirtualization.Application{
			Properties: &armdesktopvirtualization.ApplicationProperties{
				Description:          to.StringPtr("<description>"),
				CommandLineArguments: to.StringPtr("<command-line-arguments>"),
				CommandLineSetting:   armdesktopvirtualization.CommandLineSetting("Allow").ToPtr(),
				FilePath:             to.StringPtr("<file-path>"),
				FriendlyName:         to.StringPtr("<friendly-name>"),
				IconIndex:            to.Int32Ptr(1),
				IconPath:             to.StringPtr("<icon-path>"),
				ShowInPortal:         to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationsClientCreateOrUpdateResult)
}
```
