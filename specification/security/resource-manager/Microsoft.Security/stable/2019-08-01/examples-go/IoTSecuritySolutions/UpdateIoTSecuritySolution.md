Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewIotSecuritySolutionClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<solution-name>",
		armsecurity.UpdateIotSecuritySolutionData{
			Tags: map[string]*string{
				"foo": to.StringPtr("bar"),
			},
			Properties: &armsecurity.UpdateIoTSecuritySolutionProperties{
				RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
					{
						RecommendationType: armsecurity.RecommendationType("IoT_OpenPorts").ToPtr(),
						Status:             armsecurity.RecommendationConfigStatus("Disabled").ToPtr(),
					},
					{
						RecommendationType: armsecurity.RecommendationType("IoT_SharedCredentials").ToPtr(),
						Status:             armsecurity.RecommendationConfigStatus("Disabled").ToPtr(),
					}},
				UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
					Query: to.StringPtr("<query>"),
					QuerySubscriptions: []*string{
						to.StringPtr("075423e9-7d33-4166-8bdf-3920b04e3735")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IotSecuritySolutionClientUpdateResult)
}
```
