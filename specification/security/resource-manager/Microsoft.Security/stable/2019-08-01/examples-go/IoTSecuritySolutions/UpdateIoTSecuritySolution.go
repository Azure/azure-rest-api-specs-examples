package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIotSecuritySolutionClient().Update(ctx, "myRg", "default", armsecurity.UpdateIotSecuritySolutionData{
		Tags: map[string]*string{
			"foo": to.Ptr("bar"),
		},
		Properties: &armsecurity.UpdateIoTSecuritySolutionProperties{
			RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
				{
					RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTOpenPorts),
					Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
				},
				{
					RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTSharedCredentials),
					Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
				}},
			UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
				Query: to.Ptr("where type != \"microsoft.devices/iothubs\" | where name contains \"v2\""),
				QuerySubscriptions: []*string{
					to.Ptr("075423e9-7d33-4166-8bdf-3920b04e3735")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IoTSecuritySolutionModel = armsecurity.IoTSecuritySolutionModel{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Security/IoTSecuritySolutions"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/MyGroup/providers/Microsoft.Security/Locations/eastus/IoTSecuritySolutions/default"),
	// 	Tags: map[string]*string{
	// 		"foo": to.Ptr("bar"),
	// 	},
	// 	Location: to.Ptr("East Us"),
	// 	Properties: &armsecurity.IoTSecuritySolutionProperties{
	// 		AutoDiscoveredResources: []*string{
	// 			to.Ptr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735"),
	// 			to.Ptr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")},
	// 			DisabledDataSources: []*armsecurity.DataSource{
	// 			},
	// 			DisplayName: to.Ptr("Solution Default"),
	// 			Export: []*armsecurity.ExportData{
	// 				to.Ptr(armsecurity.ExportDataRawEvents)},
	// 				IotHubs: []*string{
	// 					to.Ptr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")},
	// 					RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
	// 						{
	// 							Name: to.Ptr("Service Principal Not Used with ACR"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTAcrauthentication),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Agent sending underutilized messages"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTAgentSendsUnutilizedMessages),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatus("TurnedOn")),
	// 						},
	// 						{
	// 							Name: to.Ptr("Operating system (OS) baseline validation failure"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTBaseline),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Edge Hub memory can be optimized"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTEdgeHubMemOptimize),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("No Logging Configured for Edge Module"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTEdgeLoggingOptions),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Module Settings Inconsistent in SecurityGroup"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTInconsistentModuleSettings),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Install the Azure Security of Things Agent"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTInstallAgent),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Default IP Filter Policy should be Deny"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTIpfilterDenyAll),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("IP Filter rule includes large IP range"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTIpfilterPermissiveRule),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Open Ports On Device"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTOpenPorts),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Permissive firewall policy in one of the chains was found"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTPermissiveFirewallPolicy),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Permissive firewall rule in the input chain was found"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTPermissiveInputFirewallRules),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Permissive firewall rule in the output chain was found"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTPermissiveOutputFirewallRules),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("High level permissions configured in Edge model twin for Edge module"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTPrivilegedDockerOptions),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("Same Authentication Credentials used by multiple devices"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTSharedCredentials),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
	// 						},
	// 						{
	// 							Name: to.Ptr("TLS cipher suite upgrade"),
	// 							RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTVulnerableTLSCipherSuite),
	// 							Status: to.Ptr(armsecurity.RecommendationConfigStatusEnabled),
	// 					}},
	// 					Status: to.Ptr(armsecurity.SecuritySolutionStatusEnabled),
	// 					UnmaskedIPLoggingStatus: to.Ptr(armsecurity.UnmaskedIPLoggingStatusEnabled),
	// 					UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
	// 						Query: to.Ptr("where type != \"microsoft.devices/iothubs\" | where name contains \"v2\""),
	// 						QuerySubscriptions: []*string{
	// 							to.Ptr("075423e9-7d33-4166-8bdf-3920b04e3735")},
	// 						},
	// 						Workspace: to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1"),
	// 					},
	// 					SystemData: &armsecurity.SystemData{
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.092Z"); return t}()),
	// 						CreatedBy: to.Ptr("string"),
	// 						CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-27T21:53:29.092Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("string"),
	// 						LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 					},
	// 				}
}
