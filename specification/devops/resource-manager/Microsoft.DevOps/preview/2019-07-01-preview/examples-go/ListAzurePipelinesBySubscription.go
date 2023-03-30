package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/ListAzurePipelinesBySubscription.json
func ExamplePipelinesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPipelinesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PipelineListResult = armdevops.PipelineListResult{
		// 	Value: []*armdevops.Pipeline{
		// 		{
		// 			Name: to.Ptr("myAspNetWebAppPipeline"),
		// 			Type: to.Ptr("Microsoft.DevOps/pipelines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myAspNetWebAppPipeline-rg/providers/Microsoft.DevOps/pipelines/myAspNetWebAppPipeline"),
		// 			Location: to.Ptr("South India"),
		// 			Properties: &armdevops.PipelineProperties{
		// 				BootstrapConfiguration: &armdevops.BootstrapConfiguration{
		// 					Template: &armdevops.PipelineTemplate{
		// 						ID: to.Ptr("ms.vss-continuous-delivery-pipeline-templates.aspnet-windowswebapp"),
		// 						Parameters: map[string]*string{
		// 							"appInsightLocation": to.Ptr("South India"),
		// 							"appServicePlan": to.Ptr("S1 Standard"),
		// 							"azureAuth": nil,
		// 							"location": to.Ptr("South India"),
		// 							"resourceGroup": to.Ptr("myAspNetWebAppPipeline-rg"),
		// 							"subscriptionId": to.Ptr("{subscriptionId}"),
		// 							"webAppName": to.Ptr("myAspNetWebApp"),
		// 						},
		// 					},
		// 				},
		// 				Organization: &armdevops.OrganizationReference{
		// 					Name: to.Ptr("myAspNetWebAppPipeline-org"),
		// 				},
		// 				Project: &armdevops.ProjectReference{
		// 					Name: to.Ptr("myAspNetWebAppPipeline-project"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAspNetWebAppPipeline1"),
		// 			Type: to.Ptr("Microsoft.DevOps/pipelines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myAspNetWebAppPipeline-rg1/providers/Microsoft.DevOps/pipelines/myAspNetWebAppPipeline1"),
		// 			Location: to.Ptr("South India"),
		// 			Properties: &armdevops.PipelineProperties{
		// 				BootstrapConfiguration: &armdevops.BootstrapConfiguration{
		// 					Template: &armdevops.PipelineTemplate{
		// 						ID: to.Ptr("ms.vss-continuous-delivery-pipeline-templates.aspnet-windowswebapp"),
		// 						Parameters: map[string]*string{
		// 							"appInsightLocation": to.Ptr("South India"),
		// 							"appServicePlan": to.Ptr("S1 Standard"),
		// 							"azureAuth": nil,
		// 							"location": to.Ptr("South India"),
		// 							"resourceGroup": to.Ptr("myAspNetWebAppPipeline-rg"),
		// 							"subscriptionId": to.Ptr("{subscriptionId}"),
		// 							"webAppName": to.Ptr("myAspNetWebApp"),
		// 						},
		// 					},
		// 				},
		// 				Organization: &armdevops.OrganizationReference{
		// 					Name: to.Ptr("myAspNetWebAppPipeline-org1"),
		// 				},
		// 				Project: &armdevops.ProjectReference{
		// 					Name: to.Ptr("myAspNetWebAppPipeline-project1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
