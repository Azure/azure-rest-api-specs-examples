package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/ImageDefinitions_GetImageBuildDetails.json
func ExampleProjectCatalogImageDefinitionBuildClient_GetBuildDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectCatalogImageDefinitionBuildClient().GetBuildDetails(ctx, "rg1", "DevProject", "CentralCatalog", "DefaultDevImage", "0a28fc61-6f87-4611-8fe2-32df44ab93b7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.ProjectCatalogImageDefinitionBuildClientGetBuildDetailsResponse{
	// 	ImageDefinitionBuildDetails: armdevcenter.ImageDefinitionBuildDetails{
	// 		ImageReference: &armdevcenter.ImageReference{
	// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/versions/1.0.0"),
	// 		},
	// 		Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 		TaskGroups: []*armdevcenter.ImageDefinitionBuildTaskGroup{
	// 			{
	// 				Name: to.Ptr("Provisioning"),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
	// 				Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 				Tasks: []*armdevcenter.ImageDefinitionBuildTask{
	// 					{
	// 						Name: to.Ptr("Provisioning"),
	// 						DisplayName: to.Ptr("Provisioning"),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
	// 						ID: to.Ptr("e72eb35f-f406-42ec-8b93-7d18a9f7b059"),
	// 						LogURI: to.Ptr(""),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
	// 						Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("Customizations"),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
	// 				Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 				Tasks: []*armdevcenter.ImageDefinitionBuildTask{
	// 					{
	// 						Name: to.Ptr("vs2022"),
	// 						DisplayName: to.Ptr("Install Visual Studio 2022"),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
	// 						ID: to.Ptr("e72eb35f-f406-42ec-8b93-7d18a9f7b059"),
	// 						LogURI: to.Ptr("https://8a40af38-3b4c-4672-a6a4-5e964b1870ed-contosodevcenter.centralus.devcenter.azure.com/projects/DevProject/catalogs/CentralCatalog/imageDefinitions/DefaultDevImage/builds/0a28fc61-6f87-4611-8fe2-32df44ab93b7/logs/e72eb35f-f406-42ec-8b93-7d18a9f7b059"),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
	// 						Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("Generalization"),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
	// 				Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 				Tasks: []*armdevcenter.ImageDefinitionBuildTask{
	// 					{
	// 						Name: to.Ptr("Generalization"),
	// 						DisplayName: to.Ptr("Generalization"),
	// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:12.952Z"); return t}()),
	// 						LogURI: to.Ptr(""),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:25:02.224Z"); return t}()),
	// 						Status: to.Ptr(armdevcenter.ImageDefinitionBuildStatusSucceeded),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
