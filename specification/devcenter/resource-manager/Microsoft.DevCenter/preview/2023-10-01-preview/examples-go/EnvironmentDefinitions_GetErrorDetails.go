package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/EnvironmentDefinitions_GetErrorDetails.json
func ExampleEnvironmentDefinitionsClient_GetErrorDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentDefinitionsClient().GetErrorDetails(ctx, "rg1", "Contoso", "myCatalog", "myEnvironmentDefinition", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CatalogResourceValidationErrorDetails = armdevcenter.CatalogResourceValidationErrorDetails{
	// 	Errors: []*armdevcenter.CatalogErrorDetails{
	// 		{
	// 			Code: to.Ptr("ParameterValueInvalid"),
	// 			Message: to.Ptr("Expected parameter value for 'InstanceCount' to be integer but found the string 'test'."),
	// 	}},
	// }
}
