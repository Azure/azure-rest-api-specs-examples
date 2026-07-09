package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/CustomizationTasks_GetErrorDetails.json
func ExampleCustomizationTasksClient_GetErrorDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomizationTasksClient().GetErrorDetails(ctx, "rg1", "Contoso", "CentralCatalog", "SampleTask", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.CustomizationTasksClientGetErrorDetailsResponse{
	// 	CatalogResourceValidationErrorDetails: armdevcenter.CatalogResourceValidationErrorDetails{
	// 		Errors: []*armdevcenter.CatalogErrorDetails{
	// 			{
	// 				Code: to.Ptr("ParameterValueInvalid"),
	// 				Message: to.Ptr("Expected parameter value for 'InstanceCount' to be integer but found the string 'test'."),
	// 			},
	// 		},
	// 	},
	// }
}
