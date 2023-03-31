package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_GenerateOperations.json
func ExampleProviderRegistrationsClient_GenerateOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderRegistrationsClient().GenerateOperations(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationsDefinitionArray = []*armproviderhub.OperationsDefinition{
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Read"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Read employees"),
	// 			Operation: to.Ptr("Gets/List employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(false),
	// 	},
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Write"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Writes employees"),
	// 			Operation: to.Ptr("Create/update employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(false),
	// 	},
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Delete"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Deletes employees"),
	// 			Operation: to.Ptr("Deletes employee resource"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(false),
	// 		Origin: to.Ptr(armproviderhub.OperationsDefinitionOriginUser),
	// 	},
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Action"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Writes employees"),
	// 			Operation: to.Ptr("Create/update employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(true),
	// 		Origin: to.Ptr(armproviderhub.OperationsDefinitionOriginSystem),
	// }}
}
