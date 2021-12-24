Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_GetExecutionLogs.json
func ExampleScriptExecutionsClient_GetExecutionLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	res, err := client.GetExecutionLogs(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<script-execution-name>",
		&armavs.ScriptExecutionsGetExecutionLogsOptions{ScriptOutputStreamType: []*armavs.ScriptOutputStreamType{
			armavs.ScriptOutputStreamTypeInformation.ToPtr(),
			armavs.ScriptOutputStreamTypeInformation.ToPtr(),
			armavs.ScriptOutputStreamTypeInformation.ToPtr(),
			armavs.ScriptOutputStreamTypeOutput.ToPtr()},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ScriptExecution.ID: %s\n", *res.ID)
}
```
