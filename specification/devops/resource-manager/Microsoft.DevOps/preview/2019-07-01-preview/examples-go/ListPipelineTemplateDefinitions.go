package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/ListPipelineTemplateDefinitions.json
func ExamplePipelineTemplateDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPipelineTemplateDefinitionsClient().NewListPager(nil)
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
		// page.PipelineTemplateDefinitionListResult = armdevops.PipelineTemplateDefinitionListResult{
		// 	Value: []*armdevops.PipelineTemplateDefinition{
		// 		{
		// 			Description: to.Ptr("Template for configuring CI/CD pipeline for ASP.Net app on Azure windows app service"),
		// 			ID: to.Ptr("ms.vss-continuous-delivery-pipeline-templates.aspnet-windowswebapp"),
		// 			Inputs: []*armdevops.InputDescriptor{
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("Authorization for Azure ARM endpoints."),
		// 					ID: to.Ptr("azureAuth"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("Id of subscription where azure resources will be created."),
		// 					ID: to.Ptr("subscriptionId"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("A resource group is a collection of resources that share the same lifecycle, permissions, and policies. Name of resource group which should contain web app."),
		// 					ID: to.Ptr("resourceGroup"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("Name of web app to be created"),
		// 					ID: to.Ptr("webAppName"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("Choose the Azure region that's right for you and your customers."),
		// 					ID: to.Ptr("location"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("Details of cost and compute resource associated with the web app"),
		// 					ID: to.Ptr("appServicePlan"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 						{
		// 							DisplayValue: to.Ptr("P1 Premium (1 Core, 1.75 GB RAM)"),
		// 							Value: to.Ptr("P1 Premium"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("P2 Premium (2 Core, 3.5 GB RAM)"),
		// 							Value: to.Ptr("P2 Premium"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("P3 Premium (4 Core, 7 GB RAM)"),
		// 							Value: to.Ptr("P3 Premium"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("S1 Standard (1 Core, 1.75 GB RAM)"),
		// 							Value: to.Ptr("S1 Standard"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("S2 Standard (2 Core, 3.5 GB RAM)"),
		// 							Value: to.Ptr("S2 Standard"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("S3 Standard (4 Core, 7 GB RAM)"),
		// 							Value: to.Ptr("S3 Standard"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("B1 Basic (1 Core, 1.75 GB RAM)"),
		// 							Value: to.Ptr("B1 Basic"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("B2 Basic (2 Core, 3.5 GB RAM)"),
		// 							Value: to.Ptr("B2 Basic"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("B3 Basic (4 Core, 7 GB RAM)"),
		// 							Value: to.Ptr("B3 Basic"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("F1 Free"),
		// 							Value: to.Ptr("F1 Free"),
		// 						},
		// 						{
		// 							DisplayValue: to.Ptr("D1 Shared"),
		// 							Value: to.Ptr("D1 Shared"),
		// 					}},
		// 				},
		// 				{
		// 					Type: to.Ptr(armdevops.InputDataTypeString),
		// 					Description: to.Ptr("Collect application monitoring data using Application Insights."),
		// 					ID: to.Ptr("appInsightLocation"),
		// 					PossibleValues: []*armdevops.InputValue{
		// 					},
		// 			}},
		// 	}},
		// }
	}
}
