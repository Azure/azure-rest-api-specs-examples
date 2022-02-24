Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevtestlabs%2Farmdevtestlabs%2Fv0.2.1/sdk/resourcemanager/devtestlabs/armdevtestlabs/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_GenerateArmTemplate.json
func ExampleArtifactsClient_GenerateArmTemplate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewArtifactsClient("<subscription-id>", cred, nil)
	res, err := client.GenerateArmTemplate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<artifact-source-name>",
		"<name>",
		armdevtestlabs.GenerateArmTemplateRequest{
			FileUploadOptions:  armdevtestlabs.FileUploadOptions("None").ToPtr(),
			Location:           to.StringPtr("<location>"),
			VirtualMachineName: to.StringPtr("<virtual-machine-name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ArtifactsClientGenerateArmTemplateResult)
}
```
