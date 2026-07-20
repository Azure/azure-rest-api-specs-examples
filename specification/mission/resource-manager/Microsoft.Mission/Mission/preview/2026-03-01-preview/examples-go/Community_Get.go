package armenclave_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/enclave/armenclave"
)

// Generated from example definition: 2026-03-01-preview/Community_Get.json
func ExampleCommunityClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armenclave.NewClientFactory("CA1CB369-DD26-4DB2-9D43-9AFEF0F22093", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityClient().Get(ctx, "rgopenapi", "TestMyCommunity", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armenclave.CommunityClientGetResponse{
	// 	CommunityResource: armenclave.CommunityResource{
	// 		Identity: &armenclave.ManagedServiceIdentity{
	// 			Type: to.Ptr(armenclave.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("1a2e532b-9900-414c-8600-cfc6126628d7"),
	// 			TenantID: to.Ptr("f686d426-8d16-42db-81b7-ab578e110ccd"),
	// 			UserAssignedIdentities: map[string]*armenclave.UserAssignedIdentity{
	// 				"/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armenclave.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("f4aa4954-a564-4933-a7e1-502019d807c2"),
	// 					ClientID: to.Ptr("b82bf757-ee7f-4632-9df1-5e52a720fdd2"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armenclave.CommunityProperties{
	// 			AddressSpace: to.Ptr(""),
	// 			DNSServers: []*string{
	// 				to.Ptr("azure.net"),
	// 			},
	// 			ProvisioningState: to.Ptr(armenclave.ProvisioningStateSucceeded),
	// 			ManagedOnBehalfOfConfiguration: &armenclave.ManagedOnBehalfOfConfiguration{
	// 				MoboBrokerResources: []*armenclave.MoboBrokerResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Resources/moboBrokers/bnthrkwfkfeorrzvtdxbfz"),
	// 					},
	// 				},
	// 			},
	// 			ResourceCollection: []*string{
	// 				to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg"),
	// 			},
	// 			GovernedServiceList: []*armenclave.GovernedServiceItem{
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierAKS),
	// 					ServiceName: to.Ptr("AKS"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("d300338e-65d1-4be3-b18e-fb4ce5715a8f"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierAppService),
	// 					ServiceName: to.Ptr("App Service"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("528d78c5-246c-4f26-ade6-d30798705411"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierContainerRegistry),
	// 					ServiceName: to.Ptr("Container Registry"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("b3fe25eb-cdc6-475f-96a5-04ac270f630d"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierCosmosDB),
	// 					ServiceName: to.Ptr("CosmosDB"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("6bd484ca-ae8d-46cf-9b33-e1feef84bfba"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierKeyVault),
	// 					ServiceName: to.Ptr("Key Vault"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("4f4dba0f-a5ee-494b-8df7-f9727dea6f37"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierMicrosoftSQL),
	// 					ServiceName: to.Ptr("Microsoft SQL"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("0fbe78a5-1722-4f1b-83a5-89c14151fa60"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierMonitoring),
	// 					ServiceName: to.Ptr("Monitoring"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOption("Not Applicable")),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("0a9ea1cb-7925-47fc-b0fe-8bb0a8190423"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierPostgreSQL),
	// 					ServiceName: to.Ptr("PostgreSQL"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("5eaa16b4-81f2-4354-aef3-2d77288e396e"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierServiceBus),
	// 					ServiceName: to.Ptr("Service Bus"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("8fcdb3f1-1369-426d-9917-81edfee903ab"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierStorage),
	// 					ServiceName: to.Ptr("Storage"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 						to.Ptr("ca122c06-05f6-4423-9018-ccb523168eb2"),
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierAzureFirewalls),
	// 					ServiceName: to.Ptr("Azure Firewalls"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierInsights),
	// 					ServiceName: to.Ptr("Insights"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierLogic),
	// 					ServiceName: to.Ptr("Logic"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierPrivateDNSZones),
	// 					ServiceName: to.Ptr("Private DNS Zones"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 					},
	// 				},
	// 				{
	// 					ServiceID: to.Ptr(armenclave.ServiceIdentifierDataConnectors),
	// 					ServiceName: to.Ptr("Data Connectors"),
	// 					Option: to.Ptr(armenclave.GovernedServiceItemOptionAllow),
	// 					Enforcement: to.Ptr(armenclave.GovernedServiceItemEnforcementEnabled),
	// 					PolicyAction: to.Ptr(armenclave.GovernedServiceItemPolicyActionNone),
	// 					Initiatives: []*string{
	// 					},
	// 				},
	// 			},
	// 			CommunityRoleAssignments: []*armenclave.RoleAssignmentItem{
	// 				{
	// 					RoleDefinitionID: to.Ptr("b24988ac-6180-42a0-ab88-20f7382dd24c"),
	// 					Principals: []*armenclave.Principal{
	// 						{
	// 							ID: to.Ptr("01234567-89ab-ef01-2345-0123456789ab"),
	// 							Type: to.Ptr(armenclave.PrincipalTypeGroup),
	// 						},
	// 						{
	// 							ID: to.Ptr("355a6bb0-abc0-4cba-000d-12a345b678c0"),
	// 							Type: to.Ptr(armenclave.PrincipalTypeUser),
	// 						},
	// 					},
	// 					Condition: to.Ptr("@RoleDefinition.Name StringNotEquals 'Owner'"),
	// 				},
	// 				{
	// 					RoleDefinitionID: to.Ptr("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"),
	// 					Principals: []*armenclave.Principal{
	// 						{
	// 							ID: to.Ptr("355a6bb0-abc0-4cba-000d-12a345b678c9"),
	// 							Type: to.Ptr(armenclave.PrincipalTypeUser),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			MaintenanceModeConfiguration: &armenclave.MaintenanceModeConfigurationModel{
	// 				Mode: to.Ptr(armenclave.MaintenanceModeConfigurationModelModeOff),
	// 				Principals: []*armenclave.Principal{
	// 					{
	// 						ID: to.Ptr("355a6bb0-abc0-4cba-000d-12a345b678c9"),
	// 						Type: to.Ptr(armenclave.PrincipalTypeUser),
	// 					},
	// 				},
	// 				Justification: to.Ptr(armenclave.MaintenanceModeConfigurationModelJustificationOff),
	// 			},
	// 			FirewallSKU: to.Ptr(armenclave.FirewallSKUStandard),
	// 			DedicatedHubList: []*armenclave.DedicatedHubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/rgopenapi/providers/Microsoft.Mission/communities/TestMyCommunity/dedicatedHubs/TestDedicatedHub"),
	// 					Name: to.Ptr("TestDedicatedHub"),
	// 					Type: to.Ptr("Microsoft.Mission/communities/dedicatedHubs"),
	// 					Location: to.Ptr("westcentralus"),
	// 					Properties: &armenclave.DedicatedHubProperties{
	// 						ProvisioningState: to.Ptr(armenclave.ProvisioningStateSucceeded),
	// 						VHubResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Network/virtualHubs/TestVirtualHub"),
	// 						FirewallResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Network/azureFirewalls/TestFirewall"),
	// 						FirewallPolicyResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestResourceGroup/providers/Microsoft.Network/firewallPolicies/TestFirewallPolicy"),
	// 						Designation: to.Ptr(armenclave.DesignationPooled),
	// 					},
	// 					Tags: map[string]*string{
	// 						"environment": to.Ptr("test"),
	// 					},
	// 					SystemData: &armenclave.SystemData{
	// 						CreatedBy: to.Ptr("myAlias"),
	// 						CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("myAlias"),
	// 						LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			MonitoringSettings: &armenclave.MonitoringSettingsModel{
	// 				DiagnosticDestinations: []*armenclave.MonitoringDestination{
	// 					{
	// 						DestinationType: to.Ptr(armenclave.MonitoringDestinationTypeCommunityWorkspace),
	// 					},
	// 					{
	// 						DestinationType: to.Ptr(armenclave.MonitoringDestinationTypeEnclaveWorkspace),
	// 						DiagnosticSettingsName: to.Ptr("customName"),
	// 					},
	// 					{
	// 						DestinationType: to.Ptr(armenclave.MonitoringDestinationTypeCustomWorkspace),
	// 						CustomWorkspaceResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/rgopenapi/providers/Microsoft.OperationalInsights/workspaces/CustomWorkspace"),
	// 						DiagnosticSettingsName: to.Ptr("customName"),
	// 					},
	// 					{
	// 						DestinationType: to.Ptr(armenclave.MonitoringDestinationTypeCustomWorkspace),
	// 						CustomWorkspaceResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/rgopenapi/providers/Microsoft.OperationalInsights/workspaces/CustomWorkspace"),
	// 						DiagnosticSettingsName: to.Ptr("customName"),
	// 					},
	// 				},
	// 				FlowLogDestination: &armenclave.MonitoringDestination{
	// 					DestinationType: to.Ptr(armenclave.MonitoringDestinationTypeCustomWorkspace),
	// 					CustomWorkspaceResourceID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/rgopenapi/providers/Microsoft.OperationalInsights/workspaces/CustomWorkspace"),
	// 					DiagnosticSettingsName: to.Ptr("customName"),
	// 				},
	// 			},
	// 			AddressSpaces: []*string{
	// 				to.Ptr("10.0.0.0/16"),
	// 				to.Ptr("10.1.0.0/16"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"sampletag": to.Ptr("samplevalue"),
	// 		},
	// 		Location: to.Ptr("westcentralus"),
	// 		ID: to.Ptr("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity"),
	// 		Name: to.Ptr("TestMyCommunity"),
	// 		Type: to.Ptr("Microsoft.Mission/communities"),
	// 		SystemData: &armenclave.SystemData{
	// 			CreatedBy: to.Ptr("myAlias"),
	// 			CreatedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("myAlias"),
	// 			LastModifiedByType: to.Ptr(armenclave.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-17T20:43:17.760Z"); return t}()),
	// 		},
	// 	},
	// }
}
