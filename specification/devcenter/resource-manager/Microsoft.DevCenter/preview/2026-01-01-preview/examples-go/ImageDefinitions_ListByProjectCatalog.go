package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/ImageDefinitions_ListByProjectCatalog.json
func ExampleProjectCatalogImageDefinitionsClient_NewListByProjectCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectCatalogImageDefinitionsClient().NewListByProjectCatalogPager("rg1", "ContosoProject", "TeamCatalog", nil)
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
		// page = armdevcenter.ProjectCatalogImageDefinitionsClientListByProjectCatalogResponse{
		// 	ImageDefinitionListResult: armdevcenter.ImageDefinitionListResult{
		// 		Value: []*armdevcenter.ImageDefinition{
		// 			{
		// 				Name: to.Ptr("WebDevBox"),
		// 				Type: to.Ptr("Microsoft.DevCenter/projects/catalogs/imagedefinitions"),
		// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/ContosoProject/catalogs/TeamCatalog/imagedefinitions/WebDevBox"),
		// 				Properties: &armdevcenter.ImageDefinitionProperties{
		// 					ActiveImageReference: &armdevcenter.ImageReference{
		// 						ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/versions/1.0.0"),
		// 					},
		// 					AutoImageBuild: to.Ptr(armdevcenter.AutoImageBuildStatusEnabled),
		// 					ImageReference: &armdevcenter.ImageReference{
		// 						ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/versions/1.0.0"),
		// 					},
		// 					ImageValidationErrorDetails: &armdevcenter.ImageValidationErrorDetails{
		// 						Code: to.Ptr("400"),
		// 						Message: to.Ptr("Validation failed"),
		// 					},
		// 					ImageValidationStatus: to.Ptr(armdevcenter.ImageValidationStatusFailed),
		// 					LatestBuild: &armdevcenter.LatestImageBuild{
		// 						Name: to.Ptr("0a28fc61-6f87-4611-8fe2-32df44ab93b7"),
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
		// 						Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusRunning),
		// 					},
		// 					ValidationStatus: to.Ptr(armdevcenter.CatalogResourceValidationStatusSucceeded),
		// 				},
		// 				SystemData: &armdevcenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user1"),
		// 					LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
