package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/ProjectPolicies_ListByDevCenter.json
func ExampleProjectPoliciesClient_NewListByDevCenterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58ffff1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectPoliciesClient().NewListByDevCenterPager("rg1", "Contoso", nil)
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
		// page = armdevcenter.ProjectPoliciesClientListByDevCenterResponse{
		// 	ProjectPolicyListResult: armdevcenter.ProjectPolicyListResult{
		// 		Value: []*armdevcenter.ProjectPolicy{
		// 			{
		// 				Name: to.Ptr("DevOnlyResources"),
		// 				Type: to.Ptr("Microsoft.DevCenter/devcenters/projectpolicies"),
		// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff1/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/projectPolicies/DevOnlyResources"),
		// 				Properties: &armdevcenter.ProjectPolicyProperties{
		// 					ResourcePolicies: []*armdevcenter.ResourcePolicy{
		// 						{
		// 							Resources: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff1/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/attachednetworks/network-westus3"),
		// 						},
		// 						{
		// 							ResourceType: to.Ptr(armdevcenter.ResourceTypeImages),
		// 							Action: to.Ptr(armdevcenter.PolicyActionDeny),
		// 						},
		// 						{
		// 							ResourceType: to.Ptr(armdevcenter.ResourceTypeSKUs),
		// 							Action: to.Ptr(armdevcenter.PolicyActionAllow),
		// 						},
		// 					},
		// 					Scopes: []*string{
		// 						to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff1/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject"),
		// 					},
		// 					ConfigurationPolicies: &armdevcenter.ConfigurationPolicies{
		// 						DevBoxLimitsFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 							ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 							DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
		// 							DefaultValues: []*armdevcenter.DefaultValue{
		// 								{
		// 									Name: to.Ptr("maxDevBoxesPerUser"),
		// 									Value: to.Ptr("10"),
		// 								},
		// 							},
		// 						},
		// 						AzureAiServicesFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						DevBoxScheduleDeleteFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 							ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						DisplayNameFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						DevBoxTunnelFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						ProjectCatalogFeatureStatus: &armdevcenter.FeatureState{
		// 							ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						ServerlessGpuSessionsFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 							ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						UserCustomizationsFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 						WorkspaceStorageFeatureStatus: &armdevcenter.FeatureState{
		// 							StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armdevcenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					CreatedBy: to.Ptr("User1"),
		// 					CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("User1"),
		// 					LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
