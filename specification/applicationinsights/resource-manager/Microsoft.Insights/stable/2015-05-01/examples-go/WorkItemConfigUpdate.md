```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigUpdate.json
func ExampleWorkItemConfigurationsClient_UpdateItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateItem(ctx,
		"my-resource-group",
		"my-component",
		"Visual Studio Team Services",
		armapplicationinsights.WorkItemCreateConfiguration{
			ConnectorDataConfiguration: to.Ptr("{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"d0662b05-439a-4a1b-840b-33a7f8b42ebf\",\"Custom\":\"{\\\"/fields/System.WorkItemType\\\":\\\"Bug\\\",\\\"/fields/System.AreaPath\\\":\\\"todeletefirst\\\",\\\"/fields/System.AssignedTo\\\":\\\"\\\"}\"}"),
			ConnectorID:                to.Ptr("d334e2a4-6733-488e-8645-a9fdc1694f41"),
			ValidateOnly:               to.Ptr(true),
			WorkItemProperties: map[string]*string{
				"0": to.Ptr("[object Object]"),
				"1": to.Ptr("[object Object]"),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv2.0.0-beta.1/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.
