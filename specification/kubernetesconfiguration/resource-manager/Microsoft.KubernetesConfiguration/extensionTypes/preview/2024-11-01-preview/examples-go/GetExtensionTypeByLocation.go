package armextensiontypes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensiontypes"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensionTypes/preview/2024-11-01-preview/examples/GetExtensionTypeByLocation.json
func ExampleClient_LocationGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensiontypes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().LocationGet(ctx, "westus2", "extensionType1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtensionType = armextensiontypes.ExtensionType{
	// 	Name: to.Ptr("bbbb"),
	// 	Type: to.Ptr("cccc"),
	// 	ID: to.Ptr("aaaa"),
	// 	Properties: &armextensiontypes.ExtensionTypeProperties{
	// 		Description: to.Ptr("My extension type!"),
	// 		IsManagedIdentityRequired: to.Ptr(true),
	// 		IsSystemExtension: to.Ptr(false),
	// 		PlanInfo: &armextensiontypes.ExtensionTypePropertiesPlanInfo{
	// 			OfferID: to.Ptr("myOfferId"),
	// 			PlanID: to.Ptr("myPlanId"),
	// 			PublisherID: to.Ptr("myPublisherId"),
	// 		},
	// 		Publisher: to.Ptr("Microsoft"),
	// 		SupportedClusterTypes: []*string{
	// 			to.Ptr("connectedClusters")},
	// 			SupportedScopes: &armextensiontypes.ExtensionTypePropertiesSupportedScopes{
	// 				ClusterScopeSettings: &armextensiontypes.ClusterScopeSettings{
	// 					Name: to.Ptr("bbbbbbbbbbbbb"),
	// 					Type: to.Ptr("ccccccccccccccccccccccccc"),
	// 					ID: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 					Properties: &armextensiontypes.ClusterScopeSettingsProperties{
	// 						AllowMultipleInstances: to.Ptr(true),
	// 						DefaultReleaseNamespace: to.Ptr("kube-system"),
	// 					},
	// 				},
	// 				DefaultScope: to.Ptr("cluster"),
	// 			},
	// 		},
	// 	}
}
