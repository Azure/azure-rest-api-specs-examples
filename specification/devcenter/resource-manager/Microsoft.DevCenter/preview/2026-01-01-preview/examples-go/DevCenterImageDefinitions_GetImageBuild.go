package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/DevCenterImageDefinitions_GetImageBuild.json
func ExampleCatalogImageDefinitionBuildClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCatalogImageDefinitionBuildClient().Get(ctx, "rg1", "DevDevCenter", "CentralCatalog", "DefaultDevImage", "0a28fc61-6f87-4611-8fe2-32df44ab93b7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.CatalogImageDefinitionBuildClientGetResponse{
	// 	ImageDefinitionBuild: armdevcenter.ImageDefinitionBuild{
	// 		Name: to.Ptr("0a28fc61-6f87-4611-8fe2-32df44ab93b7"),
	// 		Type: to.Ptr("Microsoft.DevCenter/centers/catalogs/imageDefinitions/builds"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/DevDevCenter/catalogs/CentralCatalog/imageDefinitions/DefaultDevImage/builds/0a28fc61-6f87-4611-8fe2-32df44ab93b7"),
	// 		Properties: &armdevcenter.ImageDefinitionBuildProperties{
	// 			ImageReference: &armdevcenter.ImageReference{
	// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/versions/1.0.0"),
	// 			},
	// 			Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			CreatedBy: to.Ptr("User1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("User1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
