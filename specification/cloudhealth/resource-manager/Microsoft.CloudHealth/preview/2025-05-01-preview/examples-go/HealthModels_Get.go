package armcloudhealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cloudhealth/armcloudhealth"
)

// Generated from example definition: 2025-05-01-preview/HealthModels_Get.json
func ExampleHealthModelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcloudhealth.NewClientFactory("4980D7D5-4E07-47AD-AD34-E76C6BC9F061", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHealthModelsClient().Get(ctx, "rgopenapi", "myHealthModel", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcloudhealth.HealthModelsClientGetResponse{
	// 	HealthModel: &armcloudhealth.HealthModel{
	// 		Properties: &armcloudhealth.HealthModelProperties{
	// 			DataplaneEndpoint: to.Ptr("https://mymodel123.healthmodels.azure.com"),
	// 			ProvisioningState: to.Ptr(armcloudhealth.HealthModelProvisioningStateSucceeded),
	// 			Discovery: &armcloudhealth.ModelDiscoverySettings{
	// 				Scope: to.Ptr("/providers/Microsoft.Management/serviceGroups/myServiceGroup"),
	// 				AddRecommendedSignals: to.Ptr(armcloudhealth.DiscoveryRuleRecommendedSignalsBehaviorEnabled),
	// 			},
	// 		},
	// 		Identity: &armcloudhealth.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("b3f9c5a0-7c5b-4a5a-8b7a-3b5fddc1b0a1"),
	// 			TenantID: to.Ptr("b3f9c5a0-7c5b-4a5a-8b7a-3b5fddc1b0a0"),
	// 			Type: to.Ptr(armcloudhealth.ManagedServiceIdentityType("SystemAssigned, UserAssigned")),
	// 			UserAssignedIdentities: map[string]*armcloudhealth.UserAssignedIdentity{
	// 				"/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ua1": &armcloudhealth.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("b3f9c5a0-7c5b-4a5a-8b7a-3b5fddc1b0a0"),
	// 					ClientID: to.Ptr("b3f9c5a0-7c5b-4a5a-8b7a-3b5fddc1b0a0"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key2961": to.Ptr("hbljozzkqrpcthsjtfkyozpwyx"),
	// 		},
	// 		Location: to.Ptr("eastus2"),
	// 		ID: to.Ptr("/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.CloudHealth/healthmodels/myHealthModel"),
	// 		Name: to.Ptr("myHealthModel"),
	// 		Type: to.Ptr("Microsoft.CloudHealth/healthmodels"),
	// 		SystemData: &armcloudhealth.SystemData{
	// 			CreatedBy: to.Ptr("cbhzxxlvkmufetjjjwtk"),
	// 			CreatedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.327Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("arz"),
	// 			LastModifiedByType: to.Ptr(armcloudhealth.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-18T14:04:09.328Z"); return t}()),
	// 		},
	// 	},
	// }
}
