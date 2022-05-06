Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdevops%2Farmdevops%2Fv0.4.0/sdk/resourcemanager/devops/armdevops/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/CreateAzurePipeline-Sample-AspNet-WindowsWebApp.json
func ExamplePipelinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<pipeline-name>",
		armdevops.Pipeline{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armdevops.PipelineProperties{
				BootstrapConfiguration: &armdevops.BootstrapConfiguration{
					Template: &armdevops.PipelineTemplate{
						ID: to.Ptr("<id>"),
						Parameters: map[string]*string{
							"appInsightLocation": to.Ptr("South India"),
							"appServicePlan":     to.Ptr("S1 Standard"),
							"azureAuth":          to.Ptr("{\"scheme\":\"ServicePrincipal\",\"parameters\":{\"tenantid\":\"{subscriptionTenantId}\",\"objectid\":\"{appObjectId}\",\"serviceprincipalid\":\"{appId}\",\"serviceprincipalkey\":\"{appSecret}\"}}"),
							"location":           to.Ptr("South India"),
							"resourceGroup":      to.Ptr("myAspNetWebAppPipeline-rg"),
							"subscriptionId":     to.Ptr("{subscriptionId}"),
							"webAppName":         to.Ptr("myAspNetWebApp"),
						},
					},
				},
				Organization: &armdevops.OrganizationReference{
					Name: to.Ptr("<name>"),
				},
				Project: &armdevops.ProjectReference{
					Name: to.Ptr("<name>"),
				},
			},
		},
		&armdevops.PipelinesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
