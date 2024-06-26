package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/DataCollectionRulesListByResourceGroup.json
func ExampleDataCollectionRulesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataCollectionRulesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.DataCollectionRuleResourceListResult = armmonitor.DataCollectionRuleResourceListResult{
		// 	Value: []*armmonitor.DataCollectionRuleResource{
		// 		{
		// 			Name: to.Ptr("myCollectionRule"),
		// 			Type: to.Ptr("Microsoft.Insights/dataCollectionRules"),
		// 			Etag: to.Ptr("070057da-0000-0000-0000-5ba70d6c0000"),
		// 			ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armmonitor.DataCollectionRuleResourceProperties{
		// 				DataFlows: []*armmonitor.DataFlow{
		// 					{
		// 						Destinations: []*string{
		// 							to.Ptr("centralWorkspace")},
		// 							Streams: []*armmonitor.KnownDataFlowStreams{
		// 								to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftPerf),
		// 								to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftSyslog),
		// 								to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftWindowsEvent)},
		// 						}},
		// 						DataSources: &armmonitor.DataCollectionRuleDataSources{
		// 							PerformanceCounters: []*armmonitor.PerfCounterDataSource{
		// 								{
		// 									Name: to.Ptr("cloudTeamCoreCounters"),
		// 									CounterSpecifiers: []*string{
		// 										to.Ptr("\\Processor(_Total)\\% Processor Time"),
		// 										to.Ptr("\\Memory\\Committed Bytes"),
		// 										to.Ptr("\\LogicalDisk(_Total)\\Free Megabytes"),
		// 										to.Ptr("\\PhysicalDisk(_Total)\\Avg. Disk Queue Length")},
		// 										SamplingFrequencyInSeconds: to.Ptr[int32](15),
		// 										Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
		// 											to.Ptr(armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf)},
		// 										},
		// 										{
		// 											Name: to.Ptr("appTeamExtraCounters"),
		// 											CounterSpecifiers: []*string{
		// 												to.Ptr("\\Process(_Total)\\Thread Count")},
		// 												SamplingFrequencyInSeconds: to.Ptr[int32](30),
		// 												Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
		// 													to.Ptr(armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf)},
		// 											}},
		// 											Syslog: []*armmonitor.SyslogDataSource{
		// 												{
		// 													Name: to.Ptr("cronSyslog"),
		// 													FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
		// 														to.Ptr(armmonitor.KnownSyslogDataSourceFacilityNamesCron)},
		// 														LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
		// 															to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsDebug),
		// 															to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsCritical),
		// 															to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsEmergency)},
		// 															Streams: []*armmonitor.KnownSyslogDataSourceStreams{
		// 																to.Ptr(armmonitor.KnownSyslogDataSourceStreamsMicrosoftSyslog)},
		// 															},
		// 															{
		// 																Name: to.Ptr("syslogBase"),
		// 																FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
		// 																	to.Ptr(armmonitor.KnownSyslogDataSourceFacilityNamesSyslog)},
		// 																	LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
		// 																		to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsAlert),
		// 																		to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsCritical),
		// 																		to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsEmergency)},
		// 																		Streams: []*armmonitor.KnownSyslogDataSourceStreams{
		// 																			to.Ptr(armmonitor.KnownSyslogDataSourceStreamsMicrosoftSyslog)},
		// 																	}},
		// 																	WindowsEventLogs: []*armmonitor.WindowsEventLogDataSource{
		// 																		{
		// 																			Name: to.Ptr("cloudSecurityTeamEvents"),
		// 																			Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
		// 																				to.Ptr(armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent)},
		// 																				XPathQueries: []*string{
		// 																					to.Ptr("Security!")},
		// 																				},
		// 																				{
		// 																					Name: to.Ptr("appTeam1AppEvents"),
		// 																					Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
		// 																						to.Ptr(armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent)},
		// 																						XPathQueries: []*string{
		// 																							to.Ptr("System![System[(Level = 1 or Level = 2 or Level = 3)]]"),
		// 																							to.Ptr("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]")},
		// 																					}},
		// 																				},
		// 																				Destinations: &armmonitor.DataCollectionRuleDestinations{
		// 																					LogAnalytics: []*armmonitor.LogAnalyticsDestination{
		// 																						{
		// 																							Name: to.Ptr("centralWorkspace"),
		// 																							WorkspaceID: to.Ptr("9ba8bc53-bd36-4156-8667-e983e7ae0e4f"),
		// 																							WorkspaceResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace"),
		// 																					}},
		// 																				},
		// 																				ImmutableID: to.Ptr("dcr-b74e0d383fc9415abaa584ec41adece3"),
		// 																			},
		// 																			SystemData: &armmonitor.DataCollectionRuleResourceSystemData{
		// 																				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
		// 																				CreatedBy: to.Ptr("user1"),
		// 																				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 																				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
		// 																				LastModifiedBy: to.Ptr("user2"),
		// 																				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 																			},
		// 																			Tags: map[string]*string{
		// 																				"tag1": to.Ptr("A"),
		// 																				"tag2": to.Ptr("B"),
		// 																			},
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("herCollectionRule"),
		// 																			Type: to.Ptr("Microsoft.Insights/dataCollectionRules"),
		// 																			Etag: to.Ptr("3afa167b-3255-432b-b66d-e74a348468af"),
		// 																			ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/herCollectionRule"),
		// 																			Location: to.Ptr("westus"),
		// 																			Properties: &armmonitor.DataCollectionRuleResourceProperties{
		// 																				DataFlows: []*armmonitor.DataFlow{
		// 																					{
		// 																						Destinations: []*string{
		// 																							to.Ptr("herWorkspace")},
		// 																							Streams: []*armmonitor.KnownDataFlowStreams{
		// 																								to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftPerf),
		// 																								to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftWindowsEvent)},
		// 																						}},
		// 																						DataSources: &armmonitor.DataCollectionRuleDataSources{
		// 																							PerformanceCounters: []*armmonitor.PerfCounterDataSource{
		// 																								{
		// 																									Name: to.Ptr("herPerfCounters"),
		// 																									CounterSpecifiers: []*string{
		// 																										to.Ptr("\\Processor(_Total)\\% Processor Time"),
		// 																										to.Ptr("\\LogicalDisk(_Total)\\Free Megabytes"),
		// 																										to.Ptr("\\PhysicalDisk(_Total)\\Avg. Disk Queue Length")},
		// 																										SamplingFrequencyInSeconds: to.Ptr[int32](30),
		// 																										Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
		// 																											to.Ptr(armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf)},
		// 																									}},
		// 																									WindowsEventLogs: []*armmonitor.WindowsEventLogDataSource{
		// 																										{
		// 																											Name: to.Ptr("herWindowsEventLogs"),
		// 																											Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
		// 																												to.Ptr(armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent)},
		// 																												XPathQueries: []*string{
		// 																													to.Ptr("Security!*"),
		// 																													to.Ptr("System!*[System[(Level = 1 or Level = 2 or Level = 3)]]"),
		// 																													to.Ptr("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]")},
		// 																											}},
		// 																										},
		// 																										Destinations: &armmonitor.DataCollectionRuleDestinations{
		// 																											LogAnalytics: []*armmonitor.LogAnalyticsDestination{
		// 																												{
		// 																													Name: to.Ptr("herWorkspace"),
		// 																													WorkspaceID: to.Ptr("c49b982a-c4f7-40c4-bd13-d6588f1b4015"),
		// 																													WorkspaceResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/herResourceGroup/providers/Microsoft.OperationalInsights/workspaces/herTeamWorkspace"),
		// 																											}},
		// 																										},
		// 																										ImmutableID: to.Ptr("dcr-ca1e6d9dad844b2c94e7961c42af917a"),
		// 																										ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleProvisioningStateSucceeded),
		// 																									},
		// 																									SystemData: &armmonitor.DataCollectionRuleResourceSystemData{
		// 																										CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T12:34:56.123Z"); return t}()),
		// 																										CreatedBy: to.Ptr("user1"),
		// 																										CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 																										LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T12:34:56.123Z"); return t}()),
		// 																										LastModifiedBy: to.Ptr("user2"),
		// 																										LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 																									},
		// 																									Tags: map[string]*string{
		// 																										"tag1": to.Ptr("C"),
		// 																									},
		// 																							}},
		// 																						}
	}
}
