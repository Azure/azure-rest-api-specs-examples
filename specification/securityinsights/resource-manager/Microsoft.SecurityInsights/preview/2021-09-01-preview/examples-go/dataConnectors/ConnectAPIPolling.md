```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/ConnectAPIPolling.json
func ExampleDataConnectorsClient_Connect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewDataConnectorsClient("<subscription-id>", cred, nil)
	_, err = client.Connect(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<data-connector-id>",
		armsecurityinsight.DataConnectorConnectBody{
			APIKey: to.StringPtr("<apikey>"),
			Kind:   armsecurityinsight.ConnectAuthKind("APIKey").ToPtr(),
			RequestConfigUserInputValues: []map[string]interface{}{
				{
					"displayText":      "Organization Name",
					"placeHolderName":  "{{placeHolder1}}",
					"placeHolderValue": "somePlaceHolderValue",
					"requestObjectKey": "apiEndpoint",
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.
