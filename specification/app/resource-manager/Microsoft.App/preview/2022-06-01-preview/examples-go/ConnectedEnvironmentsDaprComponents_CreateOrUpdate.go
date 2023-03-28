package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/212686c8383679e034b19143e13cbeb5a40ab454/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsDaprComponents_CreateOrUpdate.json
func ExampleConnectedEnvironmentsDaprComponentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsDaprComponentsClient().CreateOrUpdate(ctx, "examplerg", "myenvironment", "reddog", armappcontainers.DaprComponent{
		Properties: &armappcontainers.DaprComponentProperties{
			ComponentType: to.Ptr("state.azure.cosmosdb"),
			IgnoreErrors:  to.Ptr(false),
			InitTimeout:   to.Ptr("50s"),
			Metadata: []*armappcontainers.DaprMetadata{
				{
					Name:  to.Ptr("url"),
					Value: to.Ptr("<COSMOS-URL>"),
				},
				{
					Name:  to.Ptr("database"),
					Value: to.Ptr("itemsDB"),
				},
				{
					Name:  to.Ptr("collection"),
					Value: to.Ptr("items"),
				},
				{
					Name:      to.Ptr("masterkey"),
					SecretRef: to.Ptr("masterkey"),
				}},
			Scopes: []*string{
				to.Ptr("container-app-1"),
				to.Ptr("container-app-2")},
			Secrets: []*armappcontainers.Secret{
				{
					Name:  to.Ptr("masterkey"),
					Value: to.Ptr("keyvalue"),
				}},
			Version: to.Ptr("v1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DaprComponent = armappcontainers.DaprComponent{
	// 	Name: to.Ptr("reddog"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments/daprcomponents"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/jlaw-demo1/daprcomponents/reddog"),
	// 	Properties: &armappcontainers.DaprComponentProperties{
	// 		ComponentType: to.Ptr("state.azure.cosmosdb"),
	// 		IgnoreErrors: to.Ptr(false),
	// 		InitTimeout: to.Ptr("50s"),
	// 		Metadata: []*armappcontainers.DaprMetadata{
	// 			{
	// 				Name: to.Ptr("url"),
	// 				Value: to.Ptr("<COSMOS-URL>"),
	// 			},
	// 			{
	// 				Name: to.Ptr("database"),
	// 				Value: to.Ptr("itemsDB"),
	// 			},
	// 			{
	// 				Name: to.Ptr("collection"),
	// 				Value: to.Ptr("items"),
	// 			},
	// 			{
	// 				Name: to.Ptr("masterkey"),
	// 				SecretRef: to.Ptr("masterkey"),
	// 		}},
	// 		Scopes: []*string{
	// 			to.Ptr("container-app-1"),
	// 			to.Ptr("container-app-2")},
	// 			Secrets: []*armappcontainers.Secret{
	// 				{
	// 					Name: to.Ptr("masterkey"),
	// 			}},
	// 			Version: to.Ptr("v1"),
	// 		},
	// 	}
}
