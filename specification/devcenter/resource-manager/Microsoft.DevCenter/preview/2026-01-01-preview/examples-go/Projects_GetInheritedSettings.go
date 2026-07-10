package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Projects_GetInheritedSettings.json
func ExampleProjectsClient_GetInheritedSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().GetInheritedSettings(ctx, "rg1", "Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.ProjectsClientGetInheritedSettingsResponse{
	// 	InheritedSettingsForProject: armdevcenter.InheritedSettingsForProject{
	// 		ProjectCatalogSettings: &armdevcenter.InheritedProjectCatalogSettings{
	// 			CatalogItemSyncEnableStatus: to.Ptr(armdevcenter.CatalogItemSyncEnableStatusEnabled),
	// 			ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			DefaultValues: []*armdevcenter.DefaultValue{
	// 				{
	// 					Name: to.Ptr("environmentDefinitionCatalogItemSync"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Name: to.Ptr("imageDefinitionCatalogItemSync"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 			},
	// 		},
	// 		NetworkSettings: &armdevcenter.ProjectNetworkSettings{
	// 			MicrosoftHostedNetworkEnableStatus: to.Ptr(armdevcenter.MicrosoftHostedNetworkEnableStatusEnabled),
	// 		},
	// 		AzureAiServicesSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusAutoDeploy),
	// 		},
	// 		DevBoxScheduleDeleteSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 			ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
	// 			DefaultValues: []*armdevcenter.DefaultValue{
	// 				{
	// 					Name: to.Ptr("inactiveThreshold"),
	// 					Value: to.Ptr("P30D"),
	// 				},
	// 				{
	// 					Name: to.Ptr("gracePeriod"),
	// 					Value: to.Ptr("P3D"),
	// 				},
	// 				{
	// 					Name: to.Ptr("cancelOnConnectEnableStatus"),
	// 					Value: to.Ptr("false"),
	// 				},
	// 			},
	// 		},
	// 		DevBoxLimitsSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			DefaultValues: []*armdevcenter.DefaultValue{
	// 				{
	// 					Name: to.Ptr("maxDevBoxesPerUser"),
	// 					Value: to.Ptr("10"),
	// 				},
	// 			},
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
	// 		},
	// 		DisplayNameSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusDisabled),
	// 		},
	// 		DevBoxTunnelSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusDisabled),
	// 		},
	// 		ServerlessGpuSessionsSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			ValuesModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 			DefaultValues: []*armdevcenter.DefaultValue{
	// 				{
	// 					Name: to.Ptr("maxConcurrentSessionsPerProject"),
	// 					Value: to.Ptr("3"),
	// 				},
	// 			},
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusAutoDeploy),
	// 		},
	// 		UserCustomizationsSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableModifiable),
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusEnabled),
	// 		},
	// 		WorkspaceStorageSettings: &armdevcenter.FeatureState{
	// 			StatusModifiable: to.Ptr(armdevcenter.FeatureStateModifiableNotModifiable),
	// 			DefaultStatus: to.Ptr(armdevcenter.FeatureStatusDisabled),
	// 		},
	// 	},
	// }
}
