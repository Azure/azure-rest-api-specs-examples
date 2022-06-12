```go
package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/AccessPoliciesCreate.json
func ExampleAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtimeseriesinsights.NewAccessPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"env1",
		"ap1",
		armtimeseriesinsights.AccessPolicyCreateOrUpdateParameters{
			Properties: &armtimeseriesinsights.AccessPolicyResourceProperties{
				Description:       to.Ptr("some description"),
				PrincipalObjectID: to.Ptr("aGuid"),
				Roles: []*armtimeseriesinsights.AccessPolicyRole{
					to.Ptr(armtimeseriesinsights.AccessPolicyRoleReader)},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Ftimeseriesinsights%2Farmtimeseriesinsights%2Fv1.0.0/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights/README.md) on how to add the SDK to your project and authenticate.
