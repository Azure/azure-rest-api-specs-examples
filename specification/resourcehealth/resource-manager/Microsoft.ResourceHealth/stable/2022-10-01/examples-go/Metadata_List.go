package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b74978708bb95475562412d4654c00fbcedd9f89/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/Metadata_List.json
func ExampleMetadataClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMetadataClient().NewListPager(nil)
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
		// page.MetadataEntityListResult = armresourcehealth.MetadataEntityListResult{
		// 	Value: []*armresourcehealth.MetadataEntity{
		// 		{
		// 			Name: to.Ptr("status"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 			ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/status"),
		// 			Properties: &armresourcehealth.MetadataEntityProperties{
		// 				ApplicableScenarios: []*armresourcehealth.Scenario{
		// 					to.Ptr(armresourcehealth.ScenarioAlerts)},
		// 					DisplayName: to.Ptr("Status"),
		// 					SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 						{
		// 							DisplayName: to.Ptr("Active"),
		// 							ID: to.Ptr("Active"),
		// 						},
		// 						{
		// 							DisplayName: to.Ptr("Resolved"),
		// 							ID: to.Ptr("Resolved"),
		// 						},
		// 						{
		// 							DisplayName: to.Ptr("In Progress"),
		// 							ID: to.Ptr("In Progress"),
		// 						},
		// 						{
		// 							DisplayName: to.Ptr("Updated"),
		// 							ID: to.Ptr("Updated"),
		// 					}},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("level"),
		// 				Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 				ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/level"),
		// 				Properties: &armresourcehealth.MetadataEntityProperties{
		// 					ApplicableScenarios: []*armresourcehealth.Scenario{
		// 						to.Ptr(armresourcehealth.ScenarioAlerts)},
		// 						DisplayName: to.Ptr("Level"),
		// 						SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 							{
		// 								DisplayName: to.Ptr("Informational"),
		// 								ID: to.Ptr("Informational"),
		// 							},
		// 							{
		// 								DisplayName: to.Ptr("Critical"),
		// 								ID: to.Ptr("Critical"),
		// 						}},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("cause"),
		// 					Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 					ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/cause"),
		// 					Properties: &armresourcehealth.MetadataEntityProperties{
		// 						ApplicableScenarios: []*armresourcehealth.Scenario{
		// 							to.Ptr(armresourcehealth.ScenarioAlerts)},
		// 							DisplayName: to.Ptr("Cause"),
		// 							SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 								{
		// 									DisplayName: to.Ptr("Platform Initiated"),
		// 									ID: to.Ptr("PlatformInitiated"),
		// 								},
		// 								{
		// 									DisplayName: to.Ptr("User Initiated"),
		// 									ID: to.Ptr("UserInitiated"),
		// 								},
		// 								{
		// 									DisplayName: to.Ptr("Unknown"),
		// 									ID: to.Ptr("Unknown"),
		// 							}},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("previousHealthStatus"),
		// 						Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 						ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/previousHealthStatus"),
		// 						Properties: &armresourcehealth.MetadataEntityProperties{
		// 							ApplicableScenarios: []*armresourcehealth.Scenario{
		// 								to.Ptr(armresourcehealth.ScenarioAlerts)},
		// 								DisplayName: to.Ptr("Previous Health Status"),
		// 								SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 									{
		// 										DisplayName: to.Ptr("Available"),
		// 										ID: to.Ptr("Available"),
		// 									},
		// 									{
		// 										DisplayName: to.Ptr("Unavailable"),
		// 										ID: to.Ptr("Unavailable"),
		// 									},
		// 									{
		// 										DisplayName: to.Ptr("Degraded"),
		// 										ID: to.Ptr("Degraded"),
		// 									},
		// 									{
		// 										DisplayName: to.Ptr("Unknown"),
		// 										ID: to.Ptr("Unknown"),
		// 								}},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("currentHealthStatus"),
		// 							Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 							ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/currentHealthStatus"),
		// 							Properties: &armresourcehealth.MetadataEntityProperties{
		// 								ApplicableScenarios: []*armresourcehealth.Scenario{
		// 									to.Ptr(armresourcehealth.ScenarioAlerts)},
		// 									DisplayName: to.Ptr("Current Health Status"),
		// 									SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 										{
		// 											DisplayName: to.Ptr("Available"),
		// 											ID: to.Ptr("Available"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Unavailable"),
		// 											ID: to.Ptr("Unavailable"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Degraded"),
		// 											ID: to.Ptr("Degraded"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Unknown"),
		// 											ID: to.Ptr("Unknown"),
		// 									}},
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("supportedResourceType"),
		// 								Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 								ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/supportedResourceType"),
		// 								Properties: &armresourcehealth.MetadataEntityProperties{
		// 									DisplayName: to.Ptr("Supported Resource Type"),
		// 									SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 										{
		// 											DisplayName: to.Ptr("Analysis Services"),
		// 											ID: to.Ptr("Microsoft.AnalysisServices/servers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("API Management service"),
		// 											ID: to.Ptr("Microsoft.ApiManagement/service"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.AppPlatform/Spring"),
		// 											ID: to.Ptr("Microsoft.AppPlatform/Spring"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Data Explorer Cluster"),
		// 											ID: to.Ptr("Microsoft.Kusto/Clusters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Data Explorer Data Connection"),
		// 											ID: to.Ptr("Microsoft.Kusto/clusters/databases/dataconnections"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("microsoft.fluidrelay/fluidrelayservers"),
		// 											ID: to.Ptr("microsoft.fluidrelay/fluidrelayservers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("AKS cluster"),
		// 											ID: to.Ptr("microsoft.ContainerService/managedClusters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Dashboard/grafana"),
		// 											ID: to.Ptr("Microsoft.Dashboard/grafana"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Purview"),
		// 											ID: to.Ptr("Microsoft.Purview/accounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.SignalRService/SignalR"),
		// 											ID: to.Ptr("Microsoft.SignalRService/SignalR"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.SignalRService/WebPubSub"),
		// 											ID: to.Ptr("Microsoft.SignalRService/WebPubSub"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Batch account"),
		// 											ID: to.Ptr("Microsoft.Batch/batchAccounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Redis cache"),
		// 											ID: to.Ptr("Microsoft.Cache/Redis"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Cdn profile"),
		// 											ID: to.Ptr("Microsoft.Cdn/profiles"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Cloud services (classic)"),
		// 											ID: to.Ptr("Microsoft.ClassicCompute/DomainNames"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Virtual machine (classic)"),
		// 											ID: to.Ptr("Microsoft.ClassicCompute/virtualMachines"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Cognitive Services account"),
		// 											ID: to.Ptr("Microsoft.CognitiveServices/accounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Dedicated Host"),
		// 											ID: to.Ptr("Microsoft.Compute/hostGroups/hosts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Virtual machine"),
		// 											ID: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Virtual machine scale sets"),
		// 											ID: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("microsoft.connectedvmwarevsphere/vcenters"),
		// 											ID: to.Ptr("microsoft.connectedvmwarevsphere/vcenters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Data factory"),
		// 											ID: to.Ptr("Microsoft.DataFactory/factories"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Data Lake Analytics account"),
		// 											ID: to.Ptr("Microsoft.DataLakeAnalytics/accounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Data Lake Store account"),
		// 											ID: to.Ptr("Microsoft.DataLakeStore/accounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Database Migration Service"),
		// 											ID: to.Ptr("Microsoft.DataMigration/services"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 											ID: to.Ptr("Microsoft.DataProtection/BackupVaults"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Data Share account"),
		// 											ID: to.Ptr("Microsoft.DataShare/accounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Database for MariaDB servers"),
		// 											ID: to.Ptr("Microsoft.DBforMariaDB/servers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Database for MySQL flexible servers"),
		// 											ID: to.Ptr("Microsoft.DBforMySQL/flexibleServers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Database for MySQL servers"),
		// 											ID: to.Ptr("Microsoft.DBforMySQL/servers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Database for PostgreSQL Flexible Server"),
		// 											ID: to.Ptr("Microsoft.DBforPostgreSQL/flexibleservers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Database for PostgreSQL servers"),
		// 											ID: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 											ID: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("IoT Hub"),
		// 											ID: to.Ptr("Microsoft.Devices/IotHubs"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Digital Twins"),
		// 											ID: to.Ptr("Microsoft.DigitalTwins/DigitalTwinsInstances"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Cosmos DB"),
		// 											ID: to.Ptr("Microsoft.Documentdb/databaseaccounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Event Hub"),
		// 											ID: to.Ptr("Microsoft.EventHub/namespaces"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
		// 											ID: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("HDInsight clusters"),
		// 											ID: to.Ptr("Microsoft.HdInsight/clusters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Arc enabled servers"),
		// 											ID: to.Ptr("Microsoft.HybridCompute/machines"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.HybridNetwork/devices"),
		// 											ID: to.Ptr("Microsoft.HybridNetwork/devices"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
		// 											ID: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("IoT Central Application"),
		// 											ID: to.Ptr("Microsoft.IoTCentral/IoTApps"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Key vault"),
		// 											ID: to.Ptr("Microsoft.KeyVault/vaults"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Azure Arc enabled Kubernetes"),
		// 											ID: to.Ptr("Microsoft.Kubernetes/connectedClusters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Machine Learning Studio web services"),
		// 											ID: to.Ptr("Microsoft.MachineLearning/webServices"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Media service"),
		// 											ID: to.Ptr("Microsoft.media/mediaservices"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Application Gateway"),
		// 											ID: to.Ptr("Microsoft.Network/applicationGateways"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Network/bastionHosts"),
		// 											ID: to.Ptr("Microsoft.Network/bastionHosts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("VPN Connection"),
		// 											ID: to.Ptr("Microsoft.Network/Connections"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
		// 											ID: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Network/dnsResolvers"),
		// 											ID: to.Ptr("Microsoft.Network/dnsResolvers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Network/dnsResolvers/outboundEndpoints"),
		// 											ID: to.Ptr("Microsoft.Network/dnsResolvers/outboundEndpoints"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Network/dnszones"),
		// 											ID: to.Ptr("Microsoft.Network/dnszones"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("ExpressRoute circuit"),
		// 											ID: to.Ptr("Microsoft.Network/expressRouteCircuits"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Express RouteGateways"),
		// 											ID: to.Ptr("Microsoft.Network/ExpressRouteGateways"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Front Door"),
		// 											ID: to.Ptr("Microsoft.Network/frontdoors"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Load balancer"),
		// 											ID: to.Ptr("Microsoft.Network/LoadBalancers"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Network/NATGateways"),
		// 											ID: to.Ptr("Microsoft.Network/NATGateways"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Traffic Manager profile"),
		// 											ID: to.Ptr("Microsoft.Network/trafficmanagerprofiles"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("VPN Gateway"),
		// 											ID: to.Ptr("Microsoft.Network/virtualNetworkGateways"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Notification Hub namespace"),
		// 											ID: to.Ptr("Microsoft.NotificationHubs/Namespaces"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Log Analytics workspace"),
		// 											ID: to.Ptr("Microsoft.OperationalInsights/workspaces"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Power BI Embedded"),
		// 											ID: to.Ptr("Microsoft.PowerBIDedicated/capacities"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.RecoveryServices/vaults"),
		// 											ID: to.Ptr("Microsoft.RecoveryServices/vaults"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 											ID: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Search service"),
		// 											ID: to.Ptr("Microsoft.Search/searchServices"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Service Bus"),
		// 											ID: to.Ptr("Microsoft.ServiceBus/namespaces"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Service Fabric cluster"),
		// 											ID: to.Ptr("Microsoft.ServiceFabric/clusters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.ServiceFabric/managedClusters"),
		// 											ID: to.Ptr("Microsoft.ServiceFabric/managedClusters"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Managed database"),
		// 											ID: to.Ptr("Microsoft.Sql/managedInstances/databases"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Managed instance"),
		// 											ID: to.Ptr("Microsoft.Sql/managedInstances"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("SQL database"),
		// 											ID: to.Ptr("Microsoft.Sql/servers/databases"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Storage account"),
		// 											ID: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.StorageCache/caches"),
		// 											ID: to.Ptr("Microsoft.StorageCache/caches"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Stream Analytics job"),
		// 											ID: to.Ptr("Microsoft.StreamAnalytics/streamingjobs"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Microsoft.Synapse/workspaces"),
		// 											ID: to.Ptr("Microsoft.Synapse/workspaces"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("App Service plan"),
		// 											ID: to.Ptr("Microsoft.Web/serverFarms"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Website"),
		// 											ID: to.Ptr("Microsoft.Web/sites"),
		// 									}},
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("supportedEventTypes"),
		// 								Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 								ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/supportedEventTypes"),
		// 								Properties: &armresourcehealth.MetadataEntityProperties{
		// 									DisplayName: to.Ptr("Supported Service Health Event Types"),
		// 									SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 										{
		// 											DisplayName: to.Ptr("Planned Maintenance"),
		// 											ID: to.Ptr("Maintenance"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Health Advisories"),
		// 											ID: to.Ptr("Informational"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Security Advisories"),
		// 											ID: to.Ptr("Security"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Health Advisories"),
		// 											ID: to.Ptr("ActionRequired"),
		// 										},
		// 										{
		// 											DisplayName: to.Ptr("Service Issue"),
		// 											ID: to.Ptr("Incident"),
		// 									}},
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("supportedServices"),
		// 								Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
		// 								ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/supportedServices"),
		// 								Properties: &armresourcehealth.MetadataEntityProperties{
		// 									DisplayName: to.Ptr("Supported Service Health Services"),
		// 									SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
		// 										{
		// 											DisplayName: to.Ptr("API Management"),
		// 											ID: to.Ptr("API Management"),
		// 											ResourceTypes: []*string{
		// 												to.Ptr("microsoft.apimanagement/service")},
		// 											},
		// 											{
		// 												DisplayName: to.Ptr("Action Groups"),
		// 												ID: to.Ptr("Action Groups"),
		// 												ResourceTypes: []*string{
		// 													to.Ptr("microsoft.insights/actiongroups"),
		// 													to.Ptr("microsoft.alertsmanagement/alerts")},
		// 												},
		// 												{
		// 													DisplayName: to.Ptr("Activity Logs & Alerts"),
		// 													ID: to.Ptr("Activity Logs & Alerts"),
		// 													ResourceTypes: []*string{
		// 														to.Ptr("microsoft.insights/alertrules"),
		// 														to.Ptr("microsoft.insights/activitylogalerts"),
		// 														to.Ptr("microsoft.alertsmanagement/alerts")},
		// 													},
		// 													{
		// 														DisplayName: to.Ptr("Advisor"),
		// 														ID: to.Ptr("Advisor"),
		// 														ResourceTypes: []*string{
		// 														},
		// 													},
		// 													{
		// 														DisplayName: to.Ptr("Alerts"),
		// 														ID: to.Ptr("Alerts"),
		// 														ResourceTypes: []*string{
		// 															to.Ptr("microsoft.insights/alertrules"),
		// 															to.Ptr("microsoft.insights/activitylogalerts"),
		// 															to.Ptr("microsoft.insights/metricalerts"),
		// 															to.Ptr("microsoft.insights/scheduledqueryrules"),
		// 															to.Ptr("microsoft.insights/actiongroups"),
		// 															to.Ptr("microsoft.alertsmanagement/alerts"),
		// 															to.Ptr("microsoft.alertsmanagement/actionrules"),
		// 															to.Ptr("microsoft.alertsmanagement/smartdetectoralertrules")},
		// 														},
		// 														{
		// 															DisplayName: to.Ptr("Alerts & Metrics"),
		// 															ID: to.Ptr("Alerts & Metrics"),
		// 															ResourceTypes: []*string{
		// 																to.Ptr("microsoft.insights/alertrules"),
		// 																to.Ptr("microsoft.insights/metricalerts"),
		// 																to.Ptr("microsoft.alertsmanagement/alerts")},
		// 															},
		// 															{
		// 																DisplayName: to.Ptr("App Service"),
		// 																ID: to.Ptr("App Service"),
		// 																ResourceTypes: []*string{
		// 																	to.Ptr("microsoft.web/sites")},
		// 																},
		// 																{
		// 																	DisplayName: to.Ptr("App Service (Linux)"),
		// 																	ID: to.Ptr("App Service (Linux)"),
		// 																	ResourceTypes: []*string{
		// 																		to.Ptr("microsoft.web/sites")},
		// 																	},
		// 																	{
		// 																		DisplayName: to.Ptr("App Service (Linux) \\ Web App for Containers"),
		// 																		ID: to.Ptr("App Service (Linux) \\ Web App for Containers"),
		// 																		ResourceTypes: []*string{
		// 																			to.Ptr("microsoft.web/sites")},
		// 																		},
		// 																		{
		// 																			DisplayName: to.Ptr("App Service (Linux) \\ Web Apps"),
		// 																			ID: to.Ptr("App Service (Linux) \\ Web Apps"),
		// 																			ResourceTypes: []*string{
		// 																				to.Ptr("microsoft.web/sites")},
		// 																			},
		// 																			{
		// 																				DisplayName: to.Ptr("App Service \\ Web Apps"),
		// 																				ID: to.Ptr("App Service \\ Web Apps"),
		// 																				ResourceTypes: []*string{
		// 																					to.Ptr("microsoft.web/sites")},
		// 																				},
		// 																				{
		// 																					DisplayName: to.Ptr("Application Gateway"),
		// 																					ID: to.Ptr("Application Gateway"),
		// 																					ResourceTypes: []*string{
		// 																						to.Ptr("microsoft.network/applicationgateways")},
		// 																					},
		// 																					{
		// 																						DisplayName: to.Ptr("Application Insights"),
		// 																						ID: to.Ptr("Application Insights"),
		// 																						ResourceTypes: []*string{
		// 																							to.Ptr("microsoft.insights/components")},
		// 																						},
		// 																						{
		// 																							DisplayName: to.Ptr("AutoScale"),
		// 																							ID: to.Ptr("AutoScale"),
		// 																							ResourceTypes: []*string{
		// 																								to.Ptr("microsoft.insights/autoscalesettings")},
		// 																							},
		// 																							{
		// 																								DisplayName: to.Ptr("Automation"),
		// 																								ID: to.Ptr("Automation"),
		// 																								ResourceTypes: []*string{
		// 																									to.Ptr("microsoft.automation/automationaccounts"),
		// 																									to.Ptr("microsoft.automation/automationaccounts/runbooks"),
		// 																									to.Ptr("microsoft.automation/automationaccounts/jobs")},
		// 																								},
		// 																								{
		// 																									DisplayName: to.Ptr("Azure Active Directory"),
		// 																									ID: to.Ptr("Azure Active Directory"),
		// 																									ResourceTypes: []*string{
		// 																									},
		// 																								},
		// 																								{
		// 																									DisplayName: to.Ptr("Azure Active Directory B2C"),
		// 																									ID: to.Ptr("Azure Active Directory B2C"),
		// 																									ResourceTypes: []*string{
		// 																									},
		// 																								},
		// 																								{
		// 																									DisplayName: to.Ptr("Azure Active Directory Domain Services"),
		// 																									ID: to.Ptr("Azure Active Directory Domain Services"),
		// 																									ResourceTypes: []*string{
		// 																										to.Ptr("microsoft.aad/domainservices")},
		// 																									},
		// 																									{
		// 																										DisplayName: to.Ptr("Azure Active Directory \\ Enterprise State Roaming"),
		// 																										ID: to.Ptr("Azure Active Directory \\ Enterprise State Roaming"),
		// 																										ResourceTypes: []*string{
		// 																										},
		// 																									},
		// 																									{
		// 																										DisplayName: to.Ptr("Azure Analysis Services"),
		// 																										ID: to.Ptr("Azure Analysis Services"),
		// 																										ResourceTypes: []*string{
		// 																											to.Ptr("microsoft.analysisservices/servers")},
		// 																										},
		// 																										{
		// 																											DisplayName: to.Ptr("Azure API for FHIR"),
		// 																											ID: to.Ptr("Azure API for FHIR"),
		// 																											ResourceTypes: []*string{
		// 																												to.Ptr("Microsoft.HealthcareApis/services")},
		// 																											},
		// 																											{
		// 																												DisplayName: to.Ptr("Azure App Configuration"),
		// 																												ID: to.Ptr("Azure App Configuration"),
		// 																												ResourceTypes: []*string{
		// 																												},
		// 																											},
		// 																											{
		// 																												DisplayName: to.Ptr("Azure Applied AI Services"),
		// 																												ID: to.Ptr("Azure Applied AI Services"),
		// 																												ResourceTypes: []*string{
		// 																												},
		// 																											},
		// 																											{
		// 																												DisplayName: to.Ptr("Azure Arc enabled Kubernetes"),
		// 																												ID: to.Ptr("Azure Arc enabled Kubernetes"),
		// 																												ResourceTypes: []*string{
		// 																													to.Ptr("microsoft.kubernetes/connectedclusters")},
		// 																												},
		// 																												{
		// 																													DisplayName: to.Ptr("Azure Arc enabled data services"),
		// 																													ID: to.Ptr("Azure Arc enabled data services"),
		// 																													ResourceTypes: []*string{
		// 																													},
		// 																												},
		// 																												{
		// 																													DisplayName: to.Ptr("Azure Arc enabled servers"),
		// 																													ID: to.Ptr("Azure Arc enabled servers"),
		// 																													ResourceTypes: []*string{
		// 																														to.Ptr("microsoft.hybridcompute/machines")},
		// 																													},
		// 																													{
		// 																														DisplayName: to.Ptr("Azure Bastion"),
		// 																														ID: to.Ptr("Azure Bastion"),
		// 																														ResourceTypes: []*string{
		// 																															to.Ptr("microsoft.network/bastionhosts")},
		// 																														},
		// 																														{
		// 																															DisplayName: to.Ptr("Azure Blueprints"),
		// 																															ID: to.Ptr("Azure Blueprints"),
		// 																															ResourceTypes: []*string{
		// 																															},
		// 																														},
		// 																														{
		// 																															DisplayName: to.Ptr("Azure Bot Service"),
		// 																															ID: to.Ptr("Azure Bot Service"),
		// 																															ResourceTypes: []*string{
		// 																																to.Ptr("microsoft.botservice/botservices")},
		// 																															},
		// 																															{
		// 																																DisplayName: to.Ptr("Azure Chaos Studio"),
		// 																																ID: to.Ptr("Azure Chaos Studio"),
		// 																																ResourceTypes: []*string{
		// 																																},
		// 																															},
		// 																															{
		// 																																DisplayName: to.Ptr("Azure Communication Services"),
		// 																																ID: to.Ptr("Azure Communication Services"),
		// 																																ResourceTypes: []*string{
		// 																																	to.Ptr("microsoft.communication/communicationservices")},
		// 																																},
		// 																																{
		// 																																	DisplayName: to.Ptr("Azure Container Apps"),
		// 																																	ID: to.Ptr("Azure Container Apps"),
		// 																																	ResourceTypes: []*string{
		// 																																	},
		// 																																},
		// 																																{
		// 																																	DisplayName: to.Ptr("Azure Container Registry"),
		// 																																	ID: to.Ptr("Azure Container Registry"),
		// 																																	ResourceTypes: []*string{
		// 																																		to.Ptr("microsoft.containerregistry/registries")},
		// 																																	},
		// 																																	{
		// 																																		DisplayName: to.Ptr("Azure Container Service"),
		// 																																		ID: to.Ptr("Azure Container Service"),
		// 																																		ResourceTypes: []*string{
		// 																																			to.Ptr("microsoft.containerservice/managedclusters")},
		// 																																		},
		// 																																		{
		// 																																			DisplayName: to.Ptr("Azure Cosmos DB"),
		// 																																			ID: to.Ptr("Azure Cosmos DB"),
		// 																																			ResourceTypes: []*string{
		// 																																				to.Ptr("microsoft.documentdb/databaseaccounts")},
		// 																																			},
		// 																																			{
		// 																																				DisplayName: to.Ptr("Azure Cosmos DB for PostgreSQL"),
		// 																																				ID: to.Ptr("Azure Cosmos DB for PostgreSQL"),
		// 																																				ResourceTypes: []*string{
		// 																																					to.Ptr("microsoft.dbforpostgresql/servergroupsv2")},
		// 																																				},
		// 																																				{
		// 																																					DisplayName: to.Ptr("Azure DDoS Protection"),
		// 																																					ID: to.Ptr("Azure DDoS Protection"),
		// 																																					ResourceTypes: []*string{
		// 																																					},
		// 																																				},
		// 																																				{
		// 																																					DisplayName: to.Ptr("Azure DNS"),
		// 																																					ID: to.Ptr("Azure DNS"),
		// 																																					ResourceTypes: []*string{
		// 																																						to.Ptr("microsoft.network/dnszones")},
		// 																																					},
		// 																																					{
		// 																																						DisplayName: to.Ptr("Azure Data Explorer"),
		// 																																						ID: to.Ptr("Azure Data Explorer"),
		// 																																						ResourceTypes: []*string{
		// 																																							to.Ptr("microsoft.kusto/clusters"),
		// 																																							to.Ptr("microsoft.kusto/clusters/databases/dataconnections")},
		// 																																						},
		// 																																						{
		// 																																							DisplayName: to.Ptr("Azure Data Lake Storage Gen1"),
		// 																																							ID: to.Ptr("Azure Data Lake Storage Gen1"),
		// 																																							ResourceTypes: []*string{
		// 																																								to.Ptr("microsoft.datalakestore/accounts")},
		// 																																							},
		// 																																							{
		// 																																								DisplayName: to.Ptr("Azure Data Lake Storage Gen2"),
		// 																																								ID: to.Ptr("Azure Data Lake Storage Gen2"),
		// 																																								ResourceTypes: []*string{
		// 																																									to.Ptr("microsoft.datalakestore/accounts")},
		// 																																								},
		// 																																								{
		// 																																									DisplayName: to.Ptr("Azure Data Share"),
		// 																																									ID: to.Ptr("Azure Data Share"),
		// 																																									ResourceTypes: []*string{
		// 																																										to.Ptr("microsoft.datashare/accounts"),
		// 																																										to.Ptr("microsoft.datashare/shares"),
		// 																																										to.Ptr("microsoft.datashare/sharesubscriptions"),
		// 																																										to.Ptr("microsoft.datashare/datasets"),
		// 																																										to.Ptr("microsoft.datashare/datasetmappings"),
		// 																																										to.Ptr("microsoft.datashare/invitations")},
		// 																																									},
		// 																																									{
		// 																																										DisplayName: to.Ptr("Azure Database for MySQL"),
		// 																																										ID: to.Ptr("Azure Database for MySQL"),
		// 																																										ResourceTypes: []*string{
		// 																																											to.Ptr("microsoft.dbformysql/servers")},
		// 																																										},
		// 																																										{
		// 																																											DisplayName: to.Ptr("Azure Database for MySQL flexible servers"),
		// 																																											ID: to.Ptr("Azure Database for MySQL flexible servers"),
		// 																																											ResourceTypes: []*string{
		// 																																												to.Ptr("microsoft.dbformysql/flexibleservers")},
		// 																																											},
		// 																																											{
		// 																																												DisplayName: to.Ptr("Azure Database for PostgreSQL"),
		// 																																												ID: to.Ptr("Azure Database for PostgreSQL"),
		// 																																												ResourceTypes: []*string{
		// 																																													to.Ptr("microsoft.dbforpostgresql/servers")},
		// 																																												},
		// 																																												{
		// 																																													DisplayName: to.Ptr("Azure Database for PostgreSQL flexible servers"),
		// 																																													ID: to.Ptr("Azure Database for PostgreSQL flexible servers"),
		// 																																													ResourceTypes: []*string{
		// 																																														to.Ptr("microsoft.dbforpostgresql/flexibleservers")},
		// 																																													},
		// 																																													{
		// 																																														DisplayName: to.Ptr("Azure Database Migration Service"),
		// 																																														ID: to.Ptr("Azure Database Migration Service"),
		// 																																														ResourceTypes: []*string{
		// 																																															to.Ptr("Microsoft.DataMigration/services")},
		// 																																														},
		// 																																														{
		// 																																															DisplayName: to.Ptr("Azure Databricks"),
		// 																																															ID: to.Ptr("Azure Databricks"),
		// 																																															ResourceTypes: []*string{
		// 																																																to.Ptr("microsoft.databricks/workspaces")},
		// 																																															},
		// 																																															{
		// 																																																DisplayName: to.Ptr("Azure Dedicated Host"),
		// 																																																ID: to.Ptr("Azure Dedicated Host"),
		// 																																																ResourceTypes: []*string{
		// 																																																	to.Ptr("microsoft.compute/hostgroups/hosts")},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevOps"),
		// 																																																	ID: to.Ptr("Azure DevOps"),
		// 																																																	ResourceTypes: []*string{
		// 																																																	},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevOps \\ Artifacts"),
		// 																																																	ID: to.Ptr("Azure DevOps \\ Artifacts"),
		// 																																																	ResourceTypes: []*string{
		// 																																																	},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevOps \\ Boards"),
		// 																																																	ID: to.Ptr("Azure DevOps \\ Boards"),
		// 																																																	ResourceTypes: []*string{
		// 																																																	},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevOps \\ Pipelines"),
		// 																																																	ID: to.Ptr("Azure DevOps \\ Pipelines"),
		// 																																																	ResourceTypes: []*string{
		// 																																																	},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevOps \\ Repos"),
		// 																																																	ID: to.Ptr("Azure DevOps \\ Repos"),
		// 																																																	ResourceTypes: []*string{
		// 																																																	},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevOps \\ Test Plans"),
		// 																																																	ID: to.Ptr("Azure DevOps \\ Test Plans"),
		// 																																																	ResourceTypes: []*string{
		// 																																																	},
		// 																																																},
		// 																																																{
		// 																																																	DisplayName: to.Ptr("Azure DevTest Labs"),
		// 																																																	ID: to.Ptr("Azure DevTest Labs"),
		// 																																																	ResourceTypes: []*string{
		// 																																																		to.Ptr("microsoft.devtestlab/labs")},
		// 																																																	},
		// 																																																	{
		// 																																																		DisplayName: to.Ptr("Azure Digital Twins"),
		// 																																																		ID: to.Ptr("Azure Digital Twins"),
		// 																																																		ResourceTypes: []*string{
		// 																																																			to.Ptr("microsoft.digitaltwins/digitaltwinsinstances")},
		// 																																																		},
		// 																																																		{
		// 																																																			DisplayName: to.Ptr("Azure FarmBeats"),
		// 																																																			ID: to.Ptr("Azure FarmBeats"),
		// 																																																			ResourceTypes: []*string{
		// 																																																			},
		// 																																																		},
		// 																																																		{
		// 																																																			DisplayName: to.Ptr("Azure File Sync"),
		// 																																																			ID: to.Ptr("Azure File Sync"),
		// 																																																			ResourceTypes: []*string{
		// 																																																			},
		// 																																																		},
		// 																																																		{
		// 																																																			DisplayName: to.Ptr("Azure Firewall"),
		// 																																																			ID: to.Ptr("Azure Firewall"),
		// 																																																			ResourceTypes: []*string{
		// 																																																				to.Ptr("microsoft.network/azurefirewalls")},
		// 																																																			},
		// 																																																			{
		// 																																																				DisplayName: to.Ptr("Azure Firewall Manager"),
		// 																																																				ID: to.Ptr("Azure Firewall Manager"),
		// 																																																				ResourceTypes: []*string{
		// 																																																					to.Ptr("microsoft.network/firewallpolicies")},
		// 																																																				},
		// 																																																				{
		// 																																																					DisplayName: to.Ptr("Azure Fluid Relay"),
		// 																																																					ID: to.Ptr("Azure Fluid Relay"),
		// 																																																					ResourceTypes: []*string{
		// 																																																						to.Ptr("microsoft.fluidrelay/fluidrelayservers")},
		// 																																																					},
		// 																																																					{
		// 																																																						DisplayName: to.Ptr("Azure Frontdoor"),
		// 																																																						ID: to.Ptr("Azure Frontdoor"),
		// 																																																						ResourceTypes: []*string{
		// 																																																							to.Ptr("microsoft.network/frontdoors")},
		// 																																																						},
		// 																																																						{
		// 																																																							DisplayName: to.Ptr("Azure Health Data Services"),
		// 																																																							ID: to.Ptr("Azure Health Data Services"),
		// 																																																							ResourceTypes: []*string{
		// 																																																							},
		// 																																																						},
		// 																																																						{
		// 																																																							DisplayName: to.Ptr("Azure HPC Cache"),
		// 																																																							ID: to.Ptr("Azure HPC Cache"),
		// 																																																							ResourceTypes: []*string{
		// 																																																								to.Ptr("microsoft.storagecache/caches")},
		// 																																																							},
		// 																																																							{
		// 																																																								DisplayName: to.Ptr("Azure Information Protection"),
		// 																																																								ID: to.Ptr("Azure Information Protection"),
		// 																																																								ResourceTypes: []*string{
		// 																																																								},
		// 																																																							},
		// 																																																							{
		// 																																																								DisplayName: to.Ptr("Azure Immersive Reader"),
		// 																																																								ID: to.Ptr("Azure Immersive Reader"),
		// 																																																								ResourceTypes: []*string{
		// 																																																								},
		// 																																																							},
		// 																																																							{
		// 																																																								DisplayName: to.Ptr("Azure IoT Hub"),
		// 																																																								ID: to.Ptr("Azure IoT Hub"),
		// 																																																								ResourceTypes: []*string{
		// 																																																									to.Ptr("microsoft.devices/iothubs")},
		// 																																																								},
		// 																																																								{
		// 																																																									DisplayName: to.Ptr("Azure Key Vault Managed HSM"),
		// 																																																									ID: to.Ptr("Azure Key Vault Managed HSM"),
		// 																																																									ResourceTypes: []*string{
		// 																																																									},
		// 																																																								},
		// 																																																								{
		// 																																																									DisplayName: to.Ptr("Azure Kubernetes Service (AKS)"),
		// 																																																									ID: to.Ptr("Azure Kubernetes Service (AKS)"),
		// 																																																									ResourceTypes: []*string{
		// 																																																										to.Ptr("microsoft.containerservice/managedclusters")},
		// 																																																									},
		// 																																																									{
		// 																																																										DisplayName: to.Ptr("Azure Kubernetes Service On Azure Stack HCI"),
		// 																																																										ID: to.Ptr("Azure Kubernetes Service On Azure Stack HCI"),
		// 																																																										ResourceTypes: []*string{
		// 																																																										},
		// 																																																									},
		// 																																																									{
		// 																																																										DisplayName: to.Ptr("Azure Lab Services"),
		// 																																																										ID: to.Ptr("Azure Lab Services"),
		// 																																																										ResourceTypes: []*string{
		// 																																																										},
		// 																																																									},
		// 																																																									{
		// 																																																										DisplayName: to.Ptr("Azure Load Testing"),
		// 																																																										ID: to.Ptr("Azure Load Testing"),
		// 																																																										ResourceTypes: []*string{
		// 																																																										},
		// 																																																									},
		// 																																																									{
		// 																																																										DisplayName: to.Ptr("Azure Managed Applications"),
		// 																																																										ID: to.Ptr("Azure Managed Applications"),
		// 																																																										ResourceTypes: []*string{
		// 																																																										},
		// 																																																									},
		// 																																																									{
		// 																																																										DisplayName: to.Ptr("Azure Managed Instance for Apache Cassandra"),
		// 																																																										ID: to.Ptr("Azure Managed Instance for Apache Cassandra"),
		// 																																																										ResourceTypes: []*string{
		// 																																																										},
		// 																																																									},
		// 																																																									{
		// 																																																										DisplayName: to.Ptr("Azure Maps"),
		// 																																																										ID: to.Ptr("Azure Maps"),
		// 																																																										ResourceTypes: []*string{
		// 																																																											to.Ptr("microsoft.maps/accounts")},
		// 																																																										},
		// 																																																										{
		// 																																																											DisplayName: to.Ptr("Azure Migrate"),
		// 																																																											ID: to.Ptr("Azure Migrate"),
		// 																																																											ResourceTypes: []*string{
		// 																																																												to.Ptr("microsoft.migrate/migrateprojects")},
		// 																																																											},
		// 																																																											{
		// 																																																												DisplayName: to.Ptr("Azure Monitor"),
		// 																																																												ID: to.Ptr("Azure Monitor"),
		// 																																																												ResourceTypes: []*string{
		// 																																																													to.Ptr("microsoft.alertsmanagement/alerts"),
		// 																																																													to.Ptr("microsoft.insights/scheduledqueryrules")},
		// 																																																												},
		// 																																																												{
		// 																																																													DisplayName: to.Ptr("Azure Netapp Files"),
		// 																																																													ID: to.Ptr("Azure Netapp Files"),
		// 																																																													ResourceTypes: []*string{
		// 																																																														to.Ptr("microsoft.netapp/netappaccounts")},
		// 																																																													},
		// 																																																													{
		// 																																																														DisplayName: to.Ptr("Azure Network Function Manager - Device"),
		// 																																																														ID: to.Ptr("Azure Network Function Manager - Device"),
		// 																																																														ResourceTypes: []*string{
		// 																																																															to.Ptr("microsoft.hybridnetwork/devices")},
		// 																																																														},
		// 																																																														{
		// 																																																															DisplayName: to.Ptr("Azure Network Function Manager - Network Function"),
		// 																																																															ID: to.Ptr("Azure Network Function Manager - Network Function"),
		// 																																																															ResourceTypes: []*string{
		// 																																																																to.Ptr("microsoft.hybridnetwork/networkfunctions")},
		// 																																																															},
		// 																																																															{
		// 																																																																DisplayName: to.Ptr("Azure OpenAI Service"),
		// 																																																																ID: to.Ptr("Azure OpenAI Service"),
		// 																																																																ResourceTypes: []*string{
		// 																																																																},
		// 																																																															},
		// 																																																															{
		// 																																																																DisplayName: to.Ptr("Azure Open Datasets"),
		// 																																																																ID: to.Ptr("Azure Open Datasets"),
		// 																																																																ResourceTypes: []*string{
		// 																																																																},
		// 																																																															},
		// 																																																															{
		// 																																																																DisplayName: to.Ptr("Azure Orbital"),
		// 																																																																ID: to.Ptr("Azure Orbital"),
		// 																																																																ResourceTypes: []*string{
		// 																																																																},
		// 																																																															},
		// 																																																															{
		// 																																																																DisplayName: to.Ptr("Azure Peering Service"),
		// 																																																																ID: to.Ptr("Azure Peering Service"),
		// 																																																																ResourceTypes: []*string{
		// 																																																																	to.Ptr("microsoft.peering/peeringservices")},
		// 																																																																},
		// 																																																																{
		// 																																																																	DisplayName: to.Ptr("Azure Policy"),
		// 																																																																	ID: to.Ptr("Azure Policy"),
		// 																																																																	ResourceTypes: []*string{
		// 																																																																	},
		// 																																																																},
		// 																																																																{
		// 																																																																	DisplayName: to.Ptr("Azure Private 5G Core"),
		// 																																																																	ID: to.Ptr("Azure Private 5G Core"),
		// 																																																																	ResourceTypes: []*string{
		// 																																																																	},
		// 																																																																},
		// 																																																																{
		// 																																																																	DisplayName: to.Ptr("Azure Private Link"),
		// 																																																																	ID: to.Ptr("Azure Private Link"),
		// 																																																																	ResourceTypes: []*string{
		// 																																																																		to.Ptr("microsoft.network/privateendpoints"),
		// 																																																																		to.Ptr("microsoft.network/privatelinkservices")},
		// 																																																																	},
		// 																																																																	{
		// 																																																																		DisplayName: to.Ptr("Azure Purview"),
		// 																																																																		ID: to.Ptr("Azure Purview"),
		// 																																																																		ResourceTypes: []*string{
		// 																																																																			to.Ptr("microsoft.purview/accounts")},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			DisplayName: to.Ptr("Azure Quantum"),
		// 																																																																			ID: to.Ptr("Azure Quantum"),
		// 																																																																			ResourceTypes: []*string{
		// 																																																																			},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			DisplayName: to.Ptr("Azure Red Hat OpenShift"),
		// 																																																																			ID: to.Ptr("Azure Red Hat OpenShift"),
		// 																																																																			ResourceTypes: []*string{
		// 																																																																			},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			DisplayName: to.Ptr("Azure Reservations"),
		// 																																																																			ID: to.Ptr("Azure Reservations"),
		// 																																																																			ResourceTypes: []*string{
		// 																																																																			},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			DisplayName: to.Ptr("Azure Resource Manager"),
		// 																																																																			ID: to.Ptr("Azure Resource Manager"),
		// 																																																																			ResourceTypes: []*string{
		// 																																																																			},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			DisplayName: to.Ptr("Azure Resource Mover"),
		// 																																																																			ID: to.Ptr("Azure Resource Mover"),
		// 																																																																			ResourceTypes: []*string{
		// 																																																																				to.Ptr("microsoft.migrate/movecollections")},
		// 																																																																			},
		// 																																																																			{
		// 																																																																				DisplayName: to.Ptr("Azure Search"),
		// 																																																																				ID: to.Ptr("Azure Search"),
		// 																																																																				ResourceTypes: []*string{
		// 																																																																					to.Ptr("microsoft.search/searchservices")},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DisplayName: to.Ptr("Azure Sentinel"),
		// 																																																																					ID: to.Ptr("Azure Sentinel"),
		// 																																																																					ResourceTypes: []*string{
		// 																																																																					},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DisplayName: to.Ptr("Azure Metrics Advisor"),
		// 																																																																					ID: to.Ptr("Azure Metrics Advisor"),
		// 																																																																					ResourceTypes: []*string{
		// 																																																																					},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DisplayName: to.Ptr("Azure SignalR Service"),
		// 																																																																					ID: to.Ptr("Azure SignalR Service"),
		// 																																																																					ResourceTypes: []*string{
		// 																																																																						to.Ptr("microsoft.signalrservice/signalr")},
		// 																																																																					},
		// 																																																																					{
		// 																																																																						DisplayName: to.Ptr("Azure Spatial Anchors"),
		// 																																																																						ID: to.Ptr("Azure Spatial Anchors"),
		// 																																																																						ResourceTypes: []*string{
		// 																																																																							to.Ptr("microsoft.mixedreality/spatialanchorsaccounts")},
		// 																																																																						},
		// 																																																																						{
		// 																																																																							DisplayName: to.Ptr("Azure Sphere"),
		// 																																																																							ID: to.Ptr("Azure Sphere"),
		// 																																																																							ResourceTypes: []*string{
		// 																																																																							},
		// 																																																																						},
		// 																																																																						{
		// 																																																																							DisplayName: to.Ptr("Azure Spring Cloud"),
		// 																																																																							ID: to.Ptr("Azure Spring Cloud"),
		// 																																																																							ResourceTypes: []*string{
		// 																																																																								to.Ptr("microsoft.appplatform/spring")},
		// 																																																																							},
		// 																																																																							{
		// 																																																																								DisplayName: to.Ptr("Azure SQL Migration extension for Azure Data Studio powered by Azure Database Migration Service"),
		// 																																																																								ID: to.Ptr("Azure SQL Migration extension for Azure Data Studio powered by Azure Database Migration Service"),
		// 																																																																								ResourceTypes: []*string{
		// 																																																																									to.Ptr("microsoft.datamigration/sqlmigrationservices")},
		// 																																																																								},
		// 																																																																								{
		// 																																																																									DisplayName: to.Ptr("Azure Stack"),
		// 																																																																									ID: to.Ptr("Azure Stack"),
		// 																																																																									ResourceTypes: []*string{
		// 																																																																									},
		// 																																																																								},
		// 																																																																								{
		// 																																																																									DisplayName: to.Ptr("Azure Stack Edge"),
		// 																																																																									ID: to.Ptr("Azure Stack Edge"),
		// 																																																																									ResourceTypes: []*string{
		// 																																																																									},
		// 																																																																								},
		// 																																																																								{
		// 																																																																									DisplayName: to.Ptr("Azure Stack HCI"),
		// 																																																																									ID: to.Ptr("Azure Stack HCI"),
		// 																																																																									ResourceTypes: []*string{
		// 																																																																										to.Ptr("microsoft.azurestackhci/clusters")},
		// 																																																																									},
		// 																																																																									{
		// 																																																																										DisplayName: to.Ptr("Azure Static Web Apps"),
		// 																																																																										ID: to.Ptr("Azure Static Web Apps"),
		// 																																																																										ResourceTypes: []*string{
		// 																																																																											to.Ptr("microsoft.web/staticsites")},
		// 																																																																										},
		// 																																																																										{
		// 																																																																											DisplayName: to.Ptr("Azure Synapse Analytics"),
		// 																																																																											ID: to.Ptr("Azure Synapse Analytics"),
		// 																																																																											ResourceTypes: []*string{
		// 																																																																												to.Ptr("microsoft.sql/servers"),
		// 																																																																												to.Ptr("microsoft.synapse/workspaces")},
		// 																																																																											},
		// 																																																																											{
		// 																																																																												DisplayName: to.Ptr("Azure Video Analyzer"),
		// 																																																																												ID: to.Ptr("Azure Video Analyzer"),
		// 																																																																												ResourceTypes: []*string{
		// 																																																																												},
		// 																																																																											},
		// 																																																																											{
		// 																																																																												DisplayName: to.Ptr("Azure Video Analyzer For Media"),
		// 																																																																												ID: to.Ptr("Azure Video Analyzer For Media"),
		// 																																																																												ResourceTypes: []*string{
		// 																																																																												},
		// 																																																																											},
		// 																																																																											{
		// 																																																																												DisplayName: to.Ptr("Azure VMware Solution"),
		// 																																																																												ID: to.Ptr("Azure VMware Solution"),
		// 																																																																												ResourceTypes: []*string{
		// 																																																																													to.Ptr("microsoft.avs/privateclouds"),
		// 																																																																													to.Ptr("microsoft.avs/privateclouds/clusters")},
		// 																																																																												},
		// 																																																																												{
		// 																																																																													DisplayName: to.Ptr("Azure VMware Solution by CloudSimple"),
		// 																																																																													ID: to.Ptr("Azure VMware Solution by CloudSimple"),
		// 																																																																													ResourceTypes: []*string{
		// 																																																																													},
		// 																																																																												},
		// 																																																																												{
		// 																																																																													DisplayName: to.Ptr("Azure Web PubSub"),
		// 																																																																													ID: to.Ptr("Azure Web PubSub"),
		// 																																																																													ResourceTypes: []*string{
		// 																																																																													},
		// 																																																																												},
		// 																																																																												{
		// 																																																																													DisplayName: to.Ptr("Recovery Services vault"),
		// 																																																																													ID: to.Ptr("Recovery Services vault"),
		// 																																																																													ResourceTypes: []*string{
		// 																																																																														to.Ptr("microsoft.recoveryservices/vaults")},
		// 																																																																													},
		// 																																																																													{
		// 																																																																														DisplayName: to.Ptr("Backup vault"),
		// 																																																																														ID: to.Ptr("Backup vault"),
		// 																																																																														ResourceTypes: []*string{
		// 																																																																															to.Ptr("microsoft.dataprotection/backupvaults")},
		// 																																																																														},
		// 																																																																														{
		// 																																																																															DisplayName: to.Ptr("Batch"),
		// 																																																																															ID: to.Ptr("Batch"),
		// 																																																																															ResourceTypes: []*string{
		// 																																																																																to.Ptr("microsoft.batch/batchaccounts")},
		// 																																																																															},
		// 																																																																															{
		// 																																																																																DisplayName: to.Ptr("BareMetal Infrastructure"),
		// 																																																																																ID: to.Ptr("BareMetal Infrastructure"),
		// 																																																																																ResourceTypes: []*string{
		// 																																																																																	to.Ptr("microsoft.baremetalinfrastructure/baremetalinstances")},
		// 																																																																																},
		// 																																																																																{
		// 																																																																																	DisplayName: to.Ptr("CDN"),
		// 																																																																																	ID: to.Ptr("CDN"),
		// 																																																																																	ResourceTypes: []*string{
		// 																																																																																		to.Ptr("microsoft.cdn/profiles")},
		// 																																																																																	},
		// 																																																																																	{
		// 																																																																																		DisplayName: to.Ptr("Change Analysis"),
		// 																																																																																		ID: to.Ptr("Change Analysis"),
		// 																																																																																		ResourceTypes: []*string{
		// 																																																																																		},
		// 																																																																																	},
		// 																																																																																	{
		// 																																																																																		DisplayName: to.Ptr("Cloud Services"),
		// 																																																																																		ID: to.Ptr("Cloud Services"),
		// 																																																																																		ResourceTypes: []*string{
		// 																																																																																			to.Ptr("microsoft.classiccompute/domainnames"),
		// 																																																																																			to.Ptr("microsoft.compute/domainnames")},
		// 																																																																																		},
		// 																																																																																		{
		// 																																																																																			DisplayName: to.Ptr("Cloud Shell"),
		// 																																																																																			ID: to.Ptr("Cloud Shell"),
		// 																																																																																			ResourceTypes: []*string{
		// 																																																																																			},
		// 																																																																																		},
		// 																																																																																		{
		// 																																																																																			DisplayName: to.Ptr("Cognitive Services \\ Anomaly Detector\""),
		// 																																																																																			ID: to.Ptr("Cognitive Services \\ Anomaly Detector"),
		// 																																																																																			ResourceTypes: []*string{
		// 																																																																																			},
		// 																																																																																		},
		// 																																																																																		{
		// 																																																																																			DisplayName: to.Ptr("Cognitive Services \\ Azure Metrics Advisor"),
		// 																																																																																			ID: to.Ptr("Cognitive Services \\ Azure Metrics Advisor"),
		// 																																																																																			ResourceTypes: []*string{
		// 																																																																																			},
		// 																																																																																		},
		// 																																																																																		{
		// 																																																																																			DisplayName: to.Ptr("Cognitive Services \\ Bing Autosuggest API"),
		// 																																																																																			ID: to.Ptr("Cognitive Services \\ Bing Autosuggest API"),
		// 																																																																																			ResourceTypes: []*string{
		// 																																																																																				to.Ptr("microsoft.cognitiveservices/accounts")},
		// 																																																																																			},
		// 																																																																																			{
		// 																																																																																				DisplayName: to.Ptr("Cognitive Services \\ Bing Search API"),
		// 																																																																																				ID: to.Ptr("Cognitive Services \\ Bing Search API"),
		// 																																																																																				ResourceTypes: []*string{
		// 																																																																																					to.Ptr("microsoft.cognitiveservices/accounts")},
		// 																																																																																				},
		// 																																																																																				{
		// 																																																																																					DisplayName: to.Ptr("Cognitive Services \\ Bing Speech API"),
		// 																																																																																					ID: to.Ptr("Cognitive Services \\ Bing Speech API"),
		// 																																																																																					ResourceTypes: []*string{
		// 																																																																																					},
		// 																																																																																				},
		// 																																																																																				{
		// 																																																																																					DisplayName: to.Ptr("Cognitive Services \\ Bing Spell Check API"),
		// 																																																																																					ID: to.Ptr("Cognitive Services \\ Bing Spell Check API"),
		// 																																																																																					ResourceTypes: []*string{
		// 																																																																																						to.Ptr("microsoft.cognitiveservices/accounts")},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Bing Video Search API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Bing Video Search API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Bing Visual Search API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Bing Visual Search API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Bing Web Search API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Bing Web Search API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Computer Vision API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Computer Vision API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Content Moderator API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Content Moderator API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Custom Vision API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Custom Vision API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Face API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Face API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Form Recognizer API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Form Recognizer API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Azure Form Recognizer"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Azure Form Recognizer"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Image Search API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Image Search API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Language Understanding (LUIS)"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Language Understanding (LUIS)"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ News Search API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ News Search API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Personalizer API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Personalizer API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ QnA Maker API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ QnA Maker API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Speech Services API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Speech Services API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Speaker Recognition API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Speaker Recognition API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																						},
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						DisplayName: to.Ptr("Cognitive Services \\ Text Analytics API"),
		// 																																																																																						ID: to.Ptr("Cognitive Services \\ Text Analytics API"),
		// 																																																																																						ResourceTypes: []*string{
		// 																																																																																							to.Ptr("microsoft.cognitiveservices/accounts")},
		// 																																																																																						},
		// 																																																																																						{
		// 																																																																																							DisplayName: to.Ptr("Cognitive Services \\ Translator Speech API"),
		// 																																																																																							ID: to.Ptr("Cognitive Services \\ Translator Speech API"),
		// 																																																																																							ResourceTypes: []*string{
		// 																																																																																								to.Ptr("microsoft.cognitiveservices/accounts")},
		// 																																																																																							},
		// 																																																																																							{
		// 																																																																																								DisplayName: to.Ptr("Cognitive Services \\ Translator Text API"),
		// 																																																																																								ID: to.Ptr("Cognitive Services \\ Translator Text API"),
		// 																																																																																								ResourceTypes: []*string{
		// 																																																																																									to.Ptr("microsoft.cognitiveservices/accounts")},
		// 																																																																																								},
		// 																																																																																								{
		// 																																																																																									DisplayName: to.Ptr("Container Instances"),
		// 																																																																																									ID: to.Ptr("Container Instances"),
		// 																																																																																									ResourceTypes: []*string{
		// 																																																																																										to.Ptr("microsoft.containerinstance/containergroups")},
		// 																																																																																									},
		// 																																																																																									{
		// 																																																																																										DisplayName: to.Ptr("Container Registry"),
		// 																																																																																										ID: to.Ptr("Container Registry"),
		// 																																																																																										ResourceTypes: []*string{
		// 																																																																																											to.Ptr("microsoft.containerregistry/registries")},
		// 																																																																																										},
		// 																																																																																										{
		// 																																																																																											DisplayName: to.Ptr("Cost Management"),
		// 																																																																																											ID: to.Ptr("Cost Management"),
		// 																																																																																											ResourceTypes: []*string{
		// 																																																																																											},
		// 																																																																																										},
		// 																																																																																										{
		// 																																																																																											DisplayName: to.Ptr("Data Catalog"),
		// 																																																																																											ID: to.Ptr("Data Catalog"),
		// 																																																																																											ResourceTypes: []*string{
		// 																																																																																												to.Ptr("microsoft.datacatalog/catalogs")},
		// 																																																																																											},
		// 																																																																																											{
		// 																																																																																												DisplayName: to.Ptr("Data Factory"),
		// 																																																																																												ID: to.Ptr("Data Factory"),
		// 																																																																																												ResourceTypes: []*string{
		// 																																																																																													to.Ptr("microsoft.datafactory/datafactories")},
		// 																																																																																												},
		// 																																																																																												{
		// 																																																																																													DisplayName: to.Ptr("Data Factory V1"),
		// 																																																																																													ID: to.Ptr("Data Factory V1"),
		// 																																																																																													ResourceTypes: []*string{
		// 																																																																																														to.Ptr("microsoft.datafactory/datafactories")},
		// 																																																																																													},
		// 																																																																																													{
		// 																																																																																														DisplayName: to.Ptr("Data Factory V2"),
		// 																																																																																														ID: to.Ptr("Data Factory V2"),
		// 																																																																																														ResourceTypes: []*string{
		// 																																																																																															to.Ptr("microsoft.datafactory/factories")},
		// 																																																																																														},
		// 																																																																																														{
		// 																																																																																															DisplayName: to.Ptr("Data Factory V2 \\ SSIS Integration Runtime"),
		// 																																																																																															ID: to.Ptr("Data Factory V2 \\ SSIS Integration Runtime"),
		// 																																																																																															ResourceTypes: []*string{
		// 																																																																																																to.Ptr("microsoft.datafactory/factories")},
		// 																																																																																															},
		// 																																																																																															{
		// 																																																																																																DisplayName: to.Ptr("Data Factory \\ Azure Integration Runtime"),
		// 																																																																																																ID: to.Ptr("Data Factory \\ Azure Integration Runtime"),
		// 																																																																																																ResourceTypes: []*string{
		// 																																																																																																	to.Ptr("microsoft.datafactory/datafactories")},
		// 																																																																																																},
		// 																																																																																																{
		// 																																																																																																	DisplayName: to.Ptr("Data Lake Analytics"),
		// 																																																																																																	ID: to.Ptr("Data Lake Analytics"),
		// 																																																																																																	ResourceTypes: []*string{
		// 																																																																																																		to.Ptr("microsoft.datalakeanalytics/accounts"),
		// 																																																																																																		to.Ptr("microsoft.datalakeanalytics/accounts/storageaccounts")},
		// 																																																																																																	},
		// 																																																																																																	{
		// 																																																																																																		DisplayName: to.Ptr("Diagnostic Logs"),
		// 																																																																																																		ID: to.Ptr("Diagnostic Logs"),
		// 																																																																																																		ResourceTypes: []*string{
		// 																																																																																																		},
		// 																																																																																																	},
		// 																																																																																																	{
		// 																																																																																																		DisplayName: to.Ptr("Event Grid"),
		// 																																																																																																		ID: to.Ptr("Event Grid"),
		// 																																																																																																		ResourceTypes: []*string{
		// 																																																																																																			to.Ptr("microsoft.eventgrid/topics")},
		// 																																																																																																		},
		// 																																																																																																		{
		// 																																																																																																			DisplayName: to.Ptr("Event Hubs"),
		// 																																																																																																			ID: to.Ptr("Event Hubs"),
		// 																																																																																																			ResourceTypes: []*string{
		// 																																																																																																				to.Ptr("microsoft.eventhub/namespaces")},
		// 																																																																																																			},
		// 																																																																																																			{
		// 																																																																																																				DisplayName: to.Ptr("ExpressRoute"),
		// 																																																																																																				ID: to.Ptr("ExpressRoute"),
		// 																																																																																																				ResourceTypes: []*string{
		// 																																																																																																					to.Ptr("microsoft.network/expressroutecircuits")},
		// 																																																																																																				},
		// 																																																																																																				{
		// 																																																																																																					DisplayName: to.Ptr("ExpressRoute \\ ExpressRoute Circuits"),
		// 																																																																																																					ID: to.Ptr("ExpressRoute \\ ExpressRoute Circuits"),
		// 																																																																																																					ResourceTypes: []*string{
		// 																																																																																																						to.Ptr("microsoft.network/expressroutecircuits")},
		// 																																																																																																					},
		// 																																																																																																					{
		// 																																																																																																						DisplayName: to.Ptr("ExpressRoute \\ ExpressRoute Gateways"),
		// 																																																																																																						ID: to.Ptr("ExpressRoute \\ ExpressRoute Gateways"),
		// 																																																																																																						ResourceTypes: []*string{
		// 																																																																																																						},
		// 																																																																																																					},
		// 																																																																																																					{
		// 																																																																																																						DisplayName: to.Ptr("Functions"),
		// 																																																																																																						ID: to.Ptr("Functions"),
		// 																																																																																																						ResourceTypes: []*string{
		// 																																																																																																							to.Ptr("microsoft.web/sites")},
		// 																																																																																																						},
		// 																																																																																																						{
		// 																																																																																																							DisplayName: to.Ptr("HDInsight"),
		// 																																																																																																							ID: to.Ptr("HDInsight"),
		// 																																																																																																							ResourceTypes: []*string{
		// 																																																																																																								to.Ptr("microsoft.hdinsight/clusters")},
		// 																																																																																																							},
		// 																																																																																																							{
		// 																																																																																																								DisplayName: to.Ptr("Health Bot"),
		// 																																																																																																								ID: to.Ptr("Health Bot"),
		// 																																																																																																								ResourceTypes: []*string{
		// 																																																																																																								},
		// 																																																																																																							},
		// 																																																																																																							{
		// 																																																																																																								DisplayName: to.Ptr("HockeyApp"),
		// 																																																																																																								ID: to.Ptr("HockeyApp"),
		// 																																																																																																								ResourceTypes: []*string{
		// 																																																																																																								},
		// 																																																																																																							},
		// 																																																																																																							{
		// 																																																																																																								DisplayName: to.Ptr("Internet Analyzer"),
		// 																																																																																																								ID: to.Ptr("Internet Analyzer"),
		// 																																																																																																								ResourceTypes: []*string{
		// 																																																																																																									to.Ptr("microsoft.network/networkexperiment")},
		// 																																																																																																								},
		// 																																																																																																								{
		// 																																																																																																									DisplayName: to.Ptr("IoT Central"),
		// 																																																																																																									ID: to.Ptr("IoT Central"),
		// 																																																																																																									ResourceTypes: []*string{
		// 																																																																																																										to.Ptr("microsoft.iotcentral/iotapps")},
		// 																																																																																																									},
		// 																																																																																																									{
		// 																																																																																																										DisplayName: to.Ptr("IoT Hub \\ IoT Hub Device Provisioning Service"),
		// 																																																																																																										ID: to.Ptr("IoT Hub \\ IoT Hub Device Provisioning Service"),
		// 																																																																																																										ResourceTypes: []*string{
		// 																																																																																																										},
		// 																																																																																																									},
		// 																																																																																																									{
		// 																																																																																																										DisplayName: to.Ptr("Key Vault"),
		// 																																																																																																										ID: to.Ptr("Key Vault"),
		// 																																																																																																										ResourceTypes: []*string{
		// 																																																																																																											to.Ptr("microsoft.keyvault/vaults")},
		// 																																																																																																										},
		// 																																																																																																										{
		// 																																																																																																											DisplayName: to.Ptr("Load Balancer"),
		// 																																																																																																											ID: to.Ptr("Load Balancer"),
		// 																																																																																																											ResourceTypes: []*string{
		// 																																																																																																												to.Ptr("microsoft.network/loadbalancers")},
		// 																																																																																																											},
		// 																																																																																																											{
		// 																																																																																																												DisplayName: to.Ptr("Log Analytics"),
		// 																																																																																																												ID: to.Ptr("Log Analytics"),
		// 																																																																																																												ResourceTypes: []*string{
		// 																																																																																																													to.Ptr("microsoft.operationalinsights/workspaces")},
		// 																																																																																																												},
		// 																																																																																																												{
		// 																																																																																																													DisplayName: to.Ptr("Logic Apps"),
		// 																																																																																																													ID: to.Ptr("Logic Apps"),
		// 																																																																																																													ResourceTypes: []*string{
		// 																																																																																																														to.Ptr("microsoft.logic/workflows")},
		// 																																																																																																													},
		// 																																																																																																													{
		// 																																																																																																														DisplayName: to.Ptr("Machine Learning Services"),
		// 																																																																																																														ID: to.Ptr("Machine Learning Services"),
		// 																																																																																																														ResourceTypes: []*string{
		// 																																																																																																															to.Ptr("microsoft.machinelearning/workspaces"),
		// 																																																																																																															to.Ptr("microsoft.machinelearning/commitmentplans"),
		// 																																																																																																															to.Ptr("microsoft.machinelearning/webservices")},
		// 																																																																																																														},
		// 																																																																																																														{
		// 																																																																																																															DisplayName: to.Ptr("Machine Learning Services \\ Machine Learning Experimentation Service"),
		// 																																																																																																															ID: to.Ptr("Machine Learning Services \\ Machine Learning Experimentation Service"),
		// 																																																																																																															ResourceTypes: []*string{
		// 																																																																																																															},
		// 																																																																																																														},
		// 																																																																																																														{
		// 																																																																																																															DisplayName: to.Ptr("Machine Learning Services \\ Machine Learning Model Management"),
		// 																																																																																																															ID: to.Ptr("Machine Learning Services \\ Machine Learning Model Management"),
		// 																																																																																																															ResourceTypes: []*string{
		// 																																																																																																															},
		// 																																																																																																														},
		// 																																																																																																														{
		// 																																																																																																															DisplayName: to.Ptr("Machine Learning Studio"),
		// 																																																																																																															ID: to.Ptr("Machine Learning Studio"),
		// 																																																																																																															ResourceTypes: []*string{
		// 																																																																																																															},
		// 																																																																																																														},
		// 																																																																																																														{
		// 																																																																																																															DisplayName: to.Ptr("MariaDB"),
		// 																																																																																																															ID: to.Ptr("MariaDB"),
		// 																																																																																																															ResourceTypes: []*string{
		// 																																																																																																																to.Ptr("microsoft.dbformariadb/servers")},
		// 																																																																																																															},
		// 																																																																																																															{
		// 																																																																																																																DisplayName: to.Ptr("Marketplace"),
		// 																																																																																																																ID: to.Ptr("Marketplace"),
		// 																																																																																																																ResourceTypes: []*string{
		// 																																																																																																																},
		// 																																																																																																															},
		// 																																																																																																															{
		// 																																																																																																																DisplayName: to.Ptr("Media Services \\ Encoding"),
		// 																																																																																																																ID: to.Ptr("Media Services \\ Encoding"),
		// 																																																																																																																ResourceTypes: []*string{
		// 																																																																																																																	to.Ptr("microsoft.media/mediaservices")},
		// 																																																																																																																},
		// 																																																																																																																{
		// 																																																																																																																	DisplayName: to.Ptr("Media Services \\ Streaming"),
		// 																																																																																																																	ID: to.Ptr("Media Services \\ Streaming"),
		// 																																																																																																																	ResourceTypes: []*string{
		// 																																																																																																																		to.Ptr("microsoft.media/mediaservices")},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Azure Attestation"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Azure Attestation"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Azure portal"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Azure portal"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Azure portal \\ Marketplace"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Azure portal \\ Marketplace"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Defender for Cloud"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Defender for Cloud"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Defender for IoT"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Defender for IoT"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Entra Permissions Management"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Entra Permissions Management"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Microsoft Genomics"),
		// 																																																																																																																		ID: to.Ptr("Microsoft Genomics"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Mobile Engagement"),
		// 																																																																																																																		ID: to.Ptr("Mobile Engagement"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Multi-Factor Authentication"),
		// 																																																																																																																		ID: to.Ptr("Multi-Factor Authentication"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Network Infrastructure"),
		// 																																																																																																																		ID: to.Ptr("Network Infrastructure"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Network Watcher"),
		// 																																																																																																																		ID: to.Ptr("Network Watcher"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																		},
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		DisplayName: to.Ptr("Notification Hubs"),
		// 																																																																																																																		ID: to.Ptr("Notification Hubs"),
		// 																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																			to.Ptr("microsoft.notificationhubs/namespaces")},
		// 																																																																																																																		},
		// 																																																																																																																		{
		// 																																																																																																																			DisplayName: to.Ptr("Power BI Embedded"),
		// 																																																																																																																			ID: to.Ptr("Power BI Embedded"),
		// 																																																																																																																			ResourceTypes: []*string{
		// 																																																																																																																				to.Ptr("microsoft.powerbi/workspacecollections"),
		// 																																																																																																																				to.Ptr("microsoft.powerbidedicated/capacities")},
		// 																																																																																																																			},
		// 																																																																																																																			{
		// 																																																																																																																				DisplayName: to.Ptr("Redis Cache"),
		// 																																																																																																																				ID: to.Ptr("Redis Cache"),
		// 																																																																																																																				ResourceTypes: []*string{
		// 																																																																																																																					to.Ptr("microsoft.cache/redis")},
		// 																																																																																																																				},
		// 																																																																																																																				{
		// 																																																																																																																					DisplayName: to.Ptr("Remote Rendering"),
		// 																																																																																																																					ID: to.Ptr("Remote Rendering"),
		// 																																																																																																																					ResourceTypes: []*string{
		// 																																																																																																																						to.Ptr("microsoft.mixedreality/remoterenderingaccounts")},
		// 																																																																																																																					},
		// 																																																																																																																					{
		// 																																																																																																																						DisplayName: to.Ptr("SAP HANA on Azure Large Instances"),
		// 																																																																																																																						ID: to.Ptr("SAP HANA on Azure Large Instances"),
		// 																																																																																																																						ResourceTypes: []*string{
		// 																																																																																																																							to.Ptr("microsoft.hanaonazure/hanainstances")},
		// 																																																																																																																						},
		// 																																																																																																																						{
		// 																																																																																																																							DisplayName: to.Ptr("SQL Data Warehouse"),
		// 																																																																																																																							ID: to.Ptr("SQL Data Warehouse"),
		// 																																																																																																																							ResourceTypes: []*string{
		// 																																																																																																																								to.Ptr("microsoft.sql/servers/databases")},
		// 																																																																																																																							},
		// 																																																																																																																							{
		// 																																																																																																																								DisplayName: to.Ptr("SQL Database"),
		// 																																																																																																																								ID: to.Ptr("SQL Database"),
		// 																																																																																																																								ResourceTypes: []*string{
		// 																																																																																																																									to.Ptr("microsoft.sql/servers"),
		// 																																																																																																																									to.Ptr("microsoft.sql/servers/databases"),
		// 																																																																																																																									to.Ptr("microsoft.sql/servers/elasticpools")},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DisplayName: to.Ptr("SQL Managed Instance"),
		// 																																																																																																																									ID: to.Ptr("SQL Managed Instance"),
		// 																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																										to.Ptr("microsoft.sql/managedinstances/databases"),
		// 																																																																																																																										to.Ptr("microsoft.sql/managedinstances")},
		// 																																																																																																																									},
		// 																																																																																																																									{
		// 																																																																																																																										DisplayName: to.Ptr("SQL Server on Azure Virtual Machines"),
		// 																																																																																																																										ID: to.Ptr("SQL Server on Azure Virtual Machines"),
		// 																																																																																																																										ResourceTypes: []*string{
		// 																																																																																																																											to.Ptr("microsoft.sqlvirtualmachine/sqlvirtualmachines")},
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											DisplayName: to.Ptr("SQL Server Stretch Database"),
		// 																																																																																																																											ID: to.Ptr("SQL Server Stretch Database"),
		// 																																																																																																																											ResourceTypes: []*string{
		// 																																																																																																																											},
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											DisplayName: to.Ptr("Scheduler"),
		// 																																																																																																																											ID: to.Ptr("Scheduler"),
		// 																																																																																																																											ResourceTypes: []*string{
		// 																																																																																																																												to.Ptr("microsoft.scheduler/jobcollections")},
		// 																																																																																																																											},
		// 																																																																																																																											{
		// 																																																																																																																												DisplayName: to.Ptr("Security Center"),
		// 																																																																																																																												ID: to.Ptr("Security Center"),
		// 																																																																																																																												ResourceTypes: []*string{
		// 																																																																																																																												},
		// 																																																																																																																											},
		// 																																																																																																																											{
		// 																																																																																																																												DisplayName: to.Ptr("Service Bus"),
		// 																																																																																																																												ID: to.Ptr("Service Bus"),
		// 																																																																																																																												ResourceTypes: []*string{
		// 																																																																																																																													to.Ptr("microsoft.servicebus/namespaces")},
		// 																																																																																																																												},
		// 																																																																																																																												{
		// 																																																																																																																													DisplayName: to.Ptr("Service Fabric"),
		// 																																																																																																																													ID: to.Ptr("Service Fabric"),
		// 																																																																																																																													ResourceTypes: []*string{
		// 																																																																																																																														to.Ptr("microsoft.servicefabric/clusters")},
		// 																																																																																																																													},
		// 																																																																																																																													{
		// 																																																																																																																														DisplayName: to.Ptr("Site Recovery"),
		// 																																																																																																																														ID: to.Ptr("Site Recovery"),
		// 																																																																																																																														ResourceTypes: []*string{
		// 																																																																																																																															to.Ptr("microsoft.recoveryservices/vaults")},
		// 																																																																																																																														},
		// 																																																																																																																														{
		// 																																																																																																																															DisplayName: to.Ptr("StorSimple"),
		// 																																																																																																																															ID: to.Ptr("StorSimple"),
		// 																																																																																																																															ResourceTypes: []*string{
		// 																																																																																																																																to.Ptr("microsoft.storsimple/managers"),
		// 																																																																																																																																to.Ptr("microsoft.storsimplebvtd2/managers")},
		// 																																																																																																																															},
		// 																																																																																																																															{
		// 																																																																																																																																DisplayName: to.Ptr("Storage"),
		// 																																																																																																																																ID: to.Ptr("Storage"),
		// 																																																																																																																																ResourceTypes: []*string{
		// 																																																																																																																																	to.Ptr("microsoft.classicstorage/storageaccounts"),
		// 																																																																																																																																	to.Ptr("microsoft.storage/storageaccounts")},
		// 																																																																																																																																},
		// 																																																																																																																																{
		// 																																																																																																																																	DisplayName: to.Ptr("Stream Analytics"),
		// 																																																																																																																																	ID: to.Ptr("Stream Analytics"),
		// 																																																																																																																																	ResourceTypes: []*string{
		// 																																																																																																																																		to.Ptr("microsoft.streamanalytics/streamingjobs")},
		// 																																																																																																																																	},
		// 																																																																																																																																	{
		// 																																																																																																																																		DisplayName: to.Ptr("Subscription Management"),
		// 																																																																																																																																		ID: to.Ptr("Subscription Management"),
		// 																																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																																		},
		// 																																																																																																																																	},
		// 																																																																																																																																	{
		// 																																																																																																																																		DisplayName: to.Ptr("Time Series Insights"),
		// 																																																																																																																																		ID: to.Ptr("Time Series Insights"),
		// 																																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																																		},
		// 																																																																																																																																	},
		// 																																																																																																																																	{
		// 																																																																																																																																		DisplayName: to.Ptr("Traffic Manager"),
		// 																																																																																																																																		ID: to.Ptr("Traffic Manager"),
		// 																																																																																																																																		ResourceTypes: []*string{
		// 																																																																																																																																			to.Ptr("microsoft.network/trafficmanagerprofiles")},
		// 																																																																																																																																		},
		// 																																																																																																																																		{
		// 																																																																																																																																			DisplayName: to.Ptr("VPN Gateway"),
		// 																																																																																																																																			ID: to.Ptr("VPN Gateway"),
		// 																																																																																																																																			ResourceTypes: []*string{
		// 																																																																																																																																				to.Ptr("microsoft.network/virtualnetworkgateways"),
		// 																																																																																																																																				to.Ptr("microsoft.network/connections")},
		// 																																																																																																																																			},
		// 																																																																																																																																			{
		// 																																																																																																																																				DisplayName: to.Ptr("VPN Gateway \\ Virtual WAN"),
		// 																																																																																																																																				ID: to.Ptr("VPN Gateway \\ Virtual WAN"),
		// 																																																																																																																																				ResourceTypes: []*string{
		// 																																																																																																																																				},
		// 																																																																																																																																			},
		// 																																																																																																																																			{
		// 																																																																																																																																				DisplayName: to.Ptr("Virtual Machine Scale Sets"),
		// 																																																																																																																																				ID: to.Ptr("Virtual Machine Scale Sets"),
		// 																																																																																																																																				ResourceTypes: []*string{
		// 																																																																																																																																					to.Ptr("microsoft.compute/virtualmachinescalesets")},
		// 																																																																																																																																				},
		// 																																																																																																																																				{
		// 																																																																																																																																					DisplayName: to.Ptr("Virtual Machines"),
		// 																																																																																																																																					ID: to.Ptr("Virtual Machines"),
		// 																																																																																																																																					ResourceTypes: []*string{
		// 																																																																																																																																						to.Ptr("microsoft.classiccompute/virtualmachines"),
		// 																																																																																																																																						to.Ptr("microsoft.compute/virtualmachines")},
		// 																																																																																																																																					},
		// 																																																																																																																																					{
		// 																																																																																																																																						DisplayName: to.Ptr("Virtual Network"),
		// 																																																																																																																																						ID: to.Ptr("Virtual Network"),
		// 																																																																																																																																						ResourceTypes: []*string{
		// 																																																																																																																																							to.Ptr("microsoft.network/virtualnetworks"),
		// 																																																																																																																																							to.Ptr("microsoft.classicnetwork/virtualnetworks")},
		// 																																																																																																																																						},
		// 																																																																																																																																						{
		// 																																																																																																																																							DisplayName: to.Ptr("Windows 10 IoT Core Services"),
		// 																																																																																																																																							ID: to.Ptr("Windows 10 IoT Core Services"),
		// 																																																																																																																																							ResourceTypes: []*string{
		// 																																																																																																																																							},
		// 																																																																																																																																						},
		// 																																																																																																																																						{
		// 																																																																																																																																							DisplayName: to.Ptr("Windows Virtual Desktop"),
		// 																																																																																																																																							ID: to.Ptr("Windows Virtual Desktop"),
		// 																																																																																																																																							ResourceTypes: []*string{
		// 																																																																																																																																								to.Ptr("microsoft.desktopvirtualization/hostpools")},
		// 																																																																																																																																							},
		// 																																																																																																																																							{
		// 																																																																																																																																								DisplayName: to.Ptr("Azure Automanage"),
		// 																																																																																																																																								ID: to.Ptr("Azure Automanage"),
		// 																																																																																																																																								ResourceTypes: []*string{
		// 																																																																																																																																									to.Ptr("microsoft.automanage/configurationprofiles"),
		// 																																																																																																																																									to.Ptr("microsoft.compute/virtualmachines/providers/configurationprofileassignments"),
		// 																																																																																																																																									to.Ptr("microsoft.hybridcompute/machines/providers/configurationprofileassignments"),
		// 																																																																																																																																									to.Ptr("microsoft.Azurestackhci/clusters/providers/configurationprofileassignments")},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Microsoft Graph"),
		// 																																																																																																																																									ID: to.Ptr("Microsoft Graph"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Azure confidential ledger"),
		// 																																																																																																																																									ID: to.Ptr("Azure confidential ledger"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Azure Managed Grafana"),
		// 																																																																																																																																									ID: to.Ptr("Azure Managed Grafana"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Universal Print"),
		// 																																																																																																																																									ID: to.Ptr("Universal Print"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Azure Applied AI Services \\ Azure Video Indexer"),
		// 																																																																																																																																									ID: to.Ptr("Azure Applied AI Services \\ Azure Video Indexer"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Azure Route Server"),
		// 																																																																																																																																									ID: to.Ptr("Azure Route Server"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Virtual Network NAT"),
		// 																																																																																																																																									ID: to.Ptr("Virtual Network NAT"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									DisplayName: to.Ptr("Microsoft Defender External Attack Surface Management"),
		// 																																																																																																																																									ID: to.Ptr("Microsoft Defender External Attack Surface Management"),
		// 																																																																																																																																									ResourceTypes: []*string{
		// 																																																																																																																																									},
		// 																																																																																																																																							}},
		// 																																																																																																																																						},
		// 																																																																																																																																				}},
		// 																																																																																																																																			}
	}
}
