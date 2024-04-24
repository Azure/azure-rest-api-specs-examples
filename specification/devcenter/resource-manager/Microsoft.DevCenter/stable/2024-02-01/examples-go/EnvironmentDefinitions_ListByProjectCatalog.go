package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/EnvironmentDefinitions_ListByProjectCatalog.json
func ExampleEnvironmentDefinitionsClient_NewListByProjectCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentDefinitionsClient().NewListByProjectCatalogPager("rg1", "DevProject", "myCatalog", nil)
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
		// page.EnvironmentDefinitionListResult = armdevcenter.EnvironmentDefinitionListResult{
		// 	Value: []*armdevcenter.EnvironmentDefinition{
		// 		{
		// 			Name: to.Ptr("myEnvironmentDefinition"),
		// 			Type: to.Ptr("Microsoft.DevCenter/projects/catalogs/environmentDefinitions"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/catalogs/myCatalog/environmentDefinitions/myEnvironmentDefinition"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user1"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdevcenter.EnvironmentDefinitionProperties{
		// 				Description: to.Ptr("My sample environment definition."),
		// 				Parameters: []*armdevcenter.EnvironmentDefinitionParameter{
		// 					{
		// 						Name: to.Ptr("Function App Runtime"),
		// 						Type: to.Ptr(armdevcenter.ParameterTypeString),
		// 						ID: to.Ptr("functionAppRuntime"),
		// 						Required: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("Storage Account Type"),
		// 						Type: to.Ptr(armdevcenter.ParameterTypeString),
		// 						ID: to.Ptr("storageAccountType"),
		// 						Required: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("HTTPS only"),
		// 						Type: to.Ptr(armdevcenter.ParameterTypeBoolean),
		// 						ID: to.Ptr("httpsOnly"),
		// 						ReadOnly: to.Ptr(true),
		// 						Required: to.Ptr(true),
		// 				}},
		// 				TemplatePath: to.Ptr("azuredeploy.json"),
		// 				ValidationStatus: to.Ptr(armdevcenter.CatalogResourceValidationStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
