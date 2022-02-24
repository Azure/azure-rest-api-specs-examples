Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevops%2Farmdevops%2Fv0.2.1/sdk/resourcemanager/devops/armdevops/README.md) on how to add the SDK to your project and authenticate.

```go
package armdevops_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// x-ms-original-file: specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/CreateAzurePipeline-Sample-AspNet-WindowsWebApp.json
func ExamplePipelinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevops.NewPipelinesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<pipeline-name>",
		armdevops.Pipeline{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armdevops.PipelineProperties{
				BootstrapConfiguration: &armdevops.BootstrapConfiguration{
					Template: &armdevops.PipelineTemplate{
						ID: to.StringPtr("<id>"),
						Parameters: map[string]*string{
							"appInsightLocation": to.StringPtr("South India"),
							"appServicePlan":     to.StringPtr("S1 Standard"),
							"azureAuth":          to.StringPtr("{\"scheme\":\"ServicePrincipal\",\"parameters\":{\"tenantid\":\"{subscriptionTenantId}\",\"objectid\":\"{appObjectId}\",\"serviceprincipalid\":\"{appId}\",\"serviceprincipalkey\":\"{appSecret}\"}}"),
							"location":           to.StringPtr("South India"),
							"resourceGroup":      to.StringPtr("myAspNetWebAppPipeline-rg"),
							"subscriptionId":     to.StringPtr("{subscriptionId}"),
							"webAppName":         to.StringPtr("myAspNetWebApp"),
						},
					},
				},
				Organization: &armdevops.OrganizationReference{
					Name: to.StringPtr("<name>"),
				},
				Project: &armdevops.ProjectReference{
					Name: to.StringPtr("<name>"),
				},
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
	log.Printf("Response result: %#v\n", res.PipelinesClientCreateOrUpdateResult)
}
```
