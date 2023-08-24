package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/DaprComponents_List.json
func ExampleDaprComponentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDaprComponentsClient().NewListPager("examplerg", "myenvironment", nil)
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
		// page.DaprComponentsCollection = armappcontainers.DaprComponentsCollection{
		// 	Value: []*armappcontainers.DaprComponent{
		// 		{
		// 			Name: to.Ptr("reddog"),
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/daprcomponents"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprcomponents/reddog"),
		// 			Properties: &armappcontainers.DaprComponentProperties{
		// 				ComponentType: to.Ptr("state.azure.cosmosdb"),
		// 				IgnoreErrors: to.Ptr(false),
		// 				InitTimeout: to.Ptr("50s"),
		// 				Metadata: []*armappcontainers.DaprMetadata{
		// 					{
		// 						Name: to.Ptr("url"),
		// 						Value: to.Ptr("<COSMOS-URL>"),
		// 					},
		// 					{
		// 						Name: to.Ptr("database"),
		// 						Value: to.Ptr("itemsDB"),
		// 					},
		// 					{
		// 						Name: to.Ptr("collection"),
		// 						Value: to.Ptr("items"),
		// 					},
		// 					{
		// 						Name: to.Ptr("masterkey"),
		// 						SecretRef: to.Ptr("masterkey"),
		// 				}},
		// 				Scopes: []*string{
		// 					to.Ptr("container-app-1"),
		// 					to.Ptr("container-app-2")},
		// 					Secrets: []*armappcontainers.Secret{
		// 						{
		// 							Name: to.Ptr("masterkey"),
		// 					}},
		// 					Version: to.Ptr("v1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("vaultdog"),
		// 				Type: to.Ptr("Microsoft.App/managedEnvironments/daprcomponents"),
		// 				ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprcomponents/vaultdog"),
		// 				Properties: &armappcontainers.DaprComponentProperties{
		// 					ComponentType: to.Ptr("state.azure.cosmosdb"),
		// 					IgnoreErrors: to.Ptr(false),
		// 					InitTimeout: to.Ptr("50s"),
		// 					Metadata: []*armappcontainers.DaprMetadata{
		// 						{
		// 							Name: to.Ptr("url"),
		// 							Value: to.Ptr("<COSMOS-URL>"),
		// 						},
		// 						{
		// 							Name: to.Ptr("database"),
		// 							Value: to.Ptr("itemsDB"),
		// 						},
		// 						{
		// 							Name: to.Ptr("collection"),
		// 							Value: to.Ptr("items"),
		// 						},
		// 						{
		// 							Name: to.Ptr("masterkey"),
		// 							SecretRef: to.Ptr("masterkey"),
		// 					}},
		// 					Scopes: []*string{
		// 						to.Ptr("container-app-1"),
		// 						to.Ptr("container-app-2")},
		// 						SecretStoreComponent: to.Ptr("my-secret-store"),
		// 						Version: to.Ptr("v1"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("vaultdog"),
		// 					Type: to.Ptr("Microsoft.App/managedEnvironments/daprcomponents"),
		// 					ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/daprcomponents/vaultdog"),
		// 					Properties: &armappcontainers.DaprComponentProperties{
		// 						ComponentType: to.Ptr("state.azure.cosmosdb"),
		// 						IgnoreErrors: to.Ptr(false),
		// 						InitTimeout: to.Ptr("50s"),
		// 						Metadata: []*armappcontainers.DaprMetadata{
		// 							{
		// 								Name: to.Ptr("url"),
		// 								SecretRef: to.Ptr("cosmosdb/url"),
		// 							},
		// 							{
		// 								Name: to.Ptr("database"),
		// 								Value: to.Ptr("itemsDB"),
		// 							},
		// 							{
		// 								Name: to.Ptr("collection"),
		// 								Value: to.Ptr("items"),
		// 							},
		// 							{
		// 								Name: to.Ptr("masterkey"),
		// 								SecretRef: to.Ptr("cosmosdb/masterkey"),
		// 						}},
		// 						Scopes: []*string{
		// 							to.Ptr("container-app-1"),
		// 							to.Ptr("container-app-2")},
		// 							SecretStoreComponent: to.Ptr("my-secret-store"),
		// 							Version: to.Ptr("v1"),
		// 						},
		// 				}},
		// 			}
	}
}
