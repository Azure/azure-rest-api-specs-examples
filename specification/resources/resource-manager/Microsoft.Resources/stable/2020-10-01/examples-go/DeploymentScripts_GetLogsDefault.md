Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmdeploymentscripts%2Fv0.1.1/sdk/resourcemanager/resources/armdeploymentscripts/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeploymentscripts_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2020-10-01/examples/DeploymentScripts_GetLogsDefault.json
func ExampleDeploymentScriptsClient_GetLogsDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentscripts.NewDeploymentScriptsClient("<subscription-id>", cred, nil)
	res, err := client.GetLogsDefault(ctx,
		"<resource-group-name>",
		"<script-name>",
		&armdeploymentscripts.DeploymentScriptsGetLogsDefaultOptions{Tail: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ScriptLog.ID: %s\n", *res.ID)
}
```
