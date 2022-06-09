```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"myRg",
		"default",
		armsecurity.UpdateIotSecuritySolutionData{
			Tags: map[string]*string{
				"foo": to.Ptr("bar"),
			},
			Properties: &armsecurity.UpdateIoTSecuritySolutionProperties{
				RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
					{
						RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTOpenPorts),
						Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
					},
					{
						RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTSharedCredentials),
						Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
					}},
				UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
					Query: to.Ptr("where type != \"microsoft.devices/iothubs\" | where name contains \"v2\""),
					QuerySubscriptions: []*string{
						to.Ptr("075423e9-7d33-4166-8bdf-3920b04e3735")},
				},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.7.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.
