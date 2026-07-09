package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/ProjectPolicies_Get.json
func ExampleProjectPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58ffff1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectPoliciesClient().Get(ctx, "rg1", "Contoso", "DevOnlyResources", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.ProjectPoliciesClientGetResponse{
	// 	ProjectPolicy: armdevcenter.ProjectPolicy{
	// 		Name: to.Ptr("DevOnlyResources"),
	// 		Type: to.Ptr("Microsoft.DevCenter/devcenters/projectpolicies"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff1/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/projectPolicies/DevOnlyResources"),
	// 		Properties: &armdevcenter.ProjectPolicyProperties{
	// 			ResourcePolicies: []*armdevcenter.ResourcePolicy{
	// 				{
	// 					Resources: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff1/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/attachednetworks/network-westus3"),
	// 				},
	// 				{
	// 					ResourceType: to.Ptr(armdevcenter.ResourceTypeImages),
	// 					Action: to.Ptr(armdevcenter.PolicyActionDeny),
	// 				},
	// 				{
	// 					ResourceType: to.Ptr(armdevcenter.ResourceTypeSKUs),
	// 					Action: to.Ptr(armdevcenter.PolicyActionAllow),
	// 				},
	// 			},
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff1/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject"),
	// 			},
	// 			ConfigurationPolicies: &armdevcenter.ConfigurationPolicies{
	// 				AzureAiServicesFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusAutoDeploy),
	// 				},
	// 				DevBoxScheduleDeleteFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 					ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
	// 					DefaultValues: []*armdevcenter.DefaultValue{
	// 						{
	// 							Name: to.Ptr("inactiveThreshold"),
	// 							Value: to.Ptr("P30D"),
	// 						},
	// 						{
	// 							Name: to.Ptr("gracePeriod"),
	// 							Value: to.Ptr("P3D"),
	// 						},
	// 						{
	// 							Name: to.Ptr("cancelOnConnectEnableStatus"),
	// 							Value: to.Ptr("false"),
	// 						},
	// 					},
	// 				},
	// 				DevBoxLimitsFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
	// 					DefaultValues: []*armdevcenter.DefaultValue{
	// 						{
	// 							Name: to.Ptr("maxDevBoxesPerUser"),
	// 							Value: to.Ptr("10"),
	// 						},
	// 					},
	// 				},
	// 				DisplayNameFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusDisabled),
	// 				},
	// 				DevBoxTunnelFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusDisabled),
	// 				},
	// 				ProjectCatalogFeatureStatus: &armdevcenter.FeatureState{
	// 					ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					DefaultValues: []*armdevcenter.DefaultValue{
	// 						{
	// 							Name: to.Ptr("environmentDefinitionCatalogItemSync"),
	// 							Value: to.Ptr("Enabled"),
	// 						},
	// 						{
	// 							Name: to.Ptr("imageDefinitionCatalogItemSync"),
	// 							Value: to.Ptr("Enabled"),
	// 						},
	// 					},
	// 				},
	// 				ServerlessGpuSessionsFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusAutoDeploy),
	// 					DefaultValues: []*armdevcenter.DefaultValue{
	// 						{
	// 							Name: to.Ptr("maxConcurrentSessionsPerProject"),
	// 							Value: to.Ptr("3"),
	// 						},
	// 					},
	// 				},
	// 				UserCustomizationsFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
	// 				},
	// 				WorkspaceStorageFeatureStatus: &armdevcenter.FeatureState{
	// 					StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 					DefaultStatus: to.Ptr(armdevcenter.FeatureStatusAutoDeploy),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
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
