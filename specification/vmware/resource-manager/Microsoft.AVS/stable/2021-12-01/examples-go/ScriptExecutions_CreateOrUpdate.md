Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_CreateOrUpdate.json
func ExampleScriptExecutionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<script-execution-name>",
		armavs.ScriptExecution{
			Properties: &armavs.ScriptExecutionProperties{
				HiddenParameters: []armavs.ScriptExecutionParameterClassification{
					&armavs.ScriptSecureStringExecutionParameter{
						ScriptExecutionParameter: armavs.ScriptExecutionParameter{
							Name: to.StringPtr("<name>"),
							Type: armavs.ScriptExecutionParameterTypeSecureValue.ToPtr(),
						},
						SecureValue: to.StringPtr("<secure-value>"),
					}},
				Parameters: []armavs.ScriptExecutionParameterClassification{
					&armavs.ScriptStringExecutionParameter{
						ScriptExecutionParameter: armavs.ScriptExecutionParameter{
							Name: to.StringPtr("<name>"),
							Type: armavs.ScriptExecutionParameterTypeValue.ToPtr(),
						},
						Value: to.StringPtr("<value>"),
					},
					&armavs.ScriptStringExecutionParameter{
						ScriptExecutionParameter: armavs.ScriptExecutionParameter{
							Name: to.StringPtr("<name>"),
							Type: armavs.ScriptExecutionParameterTypeValue.ToPtr(),
						},
						Value: to.StringPtr("<value>"),
					}},
				Retention:      to.StringPtr("<retention>"),
				ScriptCmdletID: to.StringPtr("<script-cmdlet-id>"),
				Timeout:        to.StringPtr("<timeout>"),
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
	log.Printf("ScriptExecution.ID: %s\n", *res.ID)
}
```
