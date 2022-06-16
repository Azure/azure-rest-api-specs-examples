package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/CreateAzurePipeline-Sample-AspNet-WindowsWebApp.json
func ExamplePipelinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myAspNetWebAppPipeline-rg",
		"myAspNetWebAppPipeline",
		armdevops.Pipeline{
			Location: to.Ptr("South India"),
			Tags:     map[string]*string{},
			Properties: &armdevops.PipelineProperties{
				BootstrapConfiguration: &armdevops.BootstrapConfiguration{
					Template: &armdevops.PipelineTemplate{
						ID: to.Ptr("ms.vss-continuous-delivery-pipeline-templates.aspnet-windowswebapp"),
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
					Name: to.Ptr("myAspNetWebAppPipeline-org"),
				},
				Project: &armdevops.ProjectReference{
					Name: to.Ptr("myAspNetWebAppPipeline-project"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
