package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/ConfigurationServices_List.json
func ExampleConfigurationServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationServicesClient().NewListPager("myResourceGroup", "myservice", nil)
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
		// page.ConfigurationServiceResourceCollection = armappplatform.ConfigurationServiceResourceCollection{
		// 	Value: []*armappplatform.ConfigurationServiceResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/configurationServices"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/configurationServices/default"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.ConfigurationServiceProperties{
		// 				Instances: []*armappplatform.ConfigurationServiceInstance{
		// 					{
		// 						Name: to.Ptr("instance1"),
		// 						Status: to.Ptr("Running"),
		// 					},
		// 					{
		// 						Name: to.Ptr("instance2"),
		// 						Status: to.Ptr("Running"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armappplatform.ConfigurationServiceProvisioningStateSucceeded),
		// 				ResourceRequests: &armappplatform.ConfigurationServiceResourceRequests{
		// 					CPU: to.Ptr("1"),
		// 					InstanceCount: to.Ptr[int32](2),
		// 					Memory: to.Ptr("1G"),
		// 				},
		// 				Settings: &armappplatform.ConfigurationServiceSettings{
		// 					GitProperty: &armappplatform.ConfigurationServiceGitProperty{
		// 						Repositories: []*armappplatform.ConfigurationServiceGitRepository{
		// 							{
		// 								Name: to.Ptr("fake"),
		// 								Label: to.Ptr("master"),
		// 								Patterns: []*string{
		// 									to.Ptr("app1"),
		// 									to.Ptr("app2/dev")},
		// 									URI: to.Ptr("https://github.com/fake-user/fake-repository.git"),
		// 							}},
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
