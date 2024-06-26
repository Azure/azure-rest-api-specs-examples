package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armdevops.OperationListResult{
		// 	Value: []*armdevops.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DevOps/register/action"),
		// 			Display: &armdevops.OperationDisplayValue{
		// 				Description: to.Ptr("Registers the specified subscription with Microsoft.DevOps resource provider and enables the creation of Pipelines"),
		// 				Operation: to.Ptr("Register for Microsoft.DevOps"),
		// 				Provider: to.Ptr("Microsoft DevOps"),
		// 				Resource: to.Ptr("register"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevOps/pipelines/write"),
		// 			Display: &armdevops.OperationDisplayValue{
		// 				Description: to.Ptr("Creates or Updates any Pipeline"),
		// 				Operation: to.Ptr("Create or Update Pipeline"),
		// 				Provider: to.Ptr("Microsoft DevOps"),
		// 				Resource: to.Ptr("Pipelines"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevOps/pipelines/read"),
		// 			Display: &armdevops.OperationDisplayValue{
		// 				Description: to.Ptr("Reads any Pipeline"),
		// 				Operation: to.Ptr("Read Pipeline"),
		// 				Provider: to.Ptr("Microsoft DevOps"),
		// 				Resource: to.Ptr("Pipelines"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevOps/pipelines/delete"),
		// 			Display: &armdevops.OperationDisplayValue{
		// 				Description: to.Ptr("Deletes any Pipeline"),
		// 				Operation: to.Ptr("Delete Pipeline"),
		// 				Provider: to.Ptr("Microsoft DevOps"),
		// 				Resource: to.Ptr("Pipelines"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DevOps/pipelineTemplateDefinitions/read"),
		// 			Display: &armdevops.OperationDisplayValue{
		// 				Description: to.Ptr("Reads any PipelineTemplateDefinition"),
		// 				Operation: to.Ptr("Read PipelineTemplateDefinition"),
		// 				Provider: to.Ptr("Microsoft DevOps"),
		// 				Resource: to.Ptr("PipelineTemplateDefinitions"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 	}},
		// }
	}
}
