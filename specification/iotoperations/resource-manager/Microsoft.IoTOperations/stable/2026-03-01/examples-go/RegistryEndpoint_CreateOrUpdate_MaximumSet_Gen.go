package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2026-03-01/RegistryEndpoint_CreateOrUpdate_MaximumSet_Gen.json
func ExampleRegistryEndpointClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegistryEndpointClient().BeginCreateOrUpdate(ctx, "rgiotoperations", "resource-123", "resource-123", armiotoperations.RegistryEndpointResource{
		Properties: &armiotoperations.RegistryEndpointProperties{
			Host: to.Ptr("contoso.azurecr.io"),
			Authentication: &armiotoperations.RegistryEndpointAnonymousAuthentication{
				Method:            to.Ptr(armiotoperations.RegistryEndpointAuthenticationMethodAnonymous),
				AnonymousSettings: &armiotoperations.RegistryEndpointAnonymousSettings{},
			},
			CodeSigningCas: []armiotoperations.RegistryEndpointTrustedSigningKeyClassification{
				&armiotoperations.RegistryEndpointTrustedSigningKeySecret{
					Type:      to.Ptr(armiotoperations.RegistryEndpointTrustedSigningKeyTypeSecret),
					SecretRef: to.Ptr("my-secret"),
				},
				&armiotoperations.RegistryEndpointTrustedSigningKeyConfigMap{
					Type:         to.Ptr(armiotoperations.RegistryEndpointTrustedSigningKeyTypeConfigMap),
					ConfigMapRef: to.Ptr("my-configmap"),
				},
			},
		},
		ExtendedLocation: &armiotoperations.ExtendedLocation{
			Name: to.Ptr("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123"),
			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.RegistryEndpointClientCreateOrUpdateResponse{
	// 	RegistryEndpointResource: &armiotoperations.RegistryEndpointResource{
	// 		Properties: &armiotoperations.RegistryEndpointProperties{
	// 			Host: to.Ptr("contoso.azurecr.io"),
	// 			Authentication: &armiotoperations.RegistryEndpointAnonymousAuthentication{
	// 				Method: to.Ptr(armiotoperations.RegistryEndpointAuthenticationMethodAnonymous),
	// 				AnonymousSettings: &armiotoperations.RegistryEndpointAnonymousSettings{
	// 				},
	// 			},
	// 			CodeSigningCas: []armiotoperations.RegistryEndpointTrustedSigningKeyClassification{
	// 				&armiotoperations.RegistryEndpointTrustedSigningKeySecret{
	// 					Type: to.Ptr(armiotoperations.RegistryEndpointTrustedSigningKeyTypeSecret),
	// 					SecretRef: to.Ptr("my-secret"),
	// 				},
	// 				&armiotoperations.RegistryEndpointTrustedSigningKeyConfigMap{
	// 					Type: to.Ptr(armiotoperations.RegistryEndpointTrustedSigningKeyTypeConfigMap),
	// 					ConfigMapRef: to.Ptr("my-configmap"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123/registryEndpoints/resource-name123"),
	// 		Name: to.Ptr("resource-name123"),
	// 		Type: to.Ptr("Microsoft.IoTOperations/registryEndpoints"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("contosouser"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("contosouser"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}
