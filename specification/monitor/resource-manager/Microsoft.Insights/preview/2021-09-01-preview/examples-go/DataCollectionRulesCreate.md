Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.7.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRulesCreate.json
func ExampleDataCollectionRulesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewDataCollectionRulesClient("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"myResourceGroup",
		"myCollectionRule",
		&armmonitor.DataCollectionRulesClientCreateOptions{Body: &armmonitor.DataCollectionRuleResource{
			Location: to.Ptr("eastus"),
			Properties: &armmonitor.DataCollectionRuleResourceProperties{
				DataFlows: []*armmonitor.DataFlow{
					{
						Destinations: []*string{
							to.Ptr("centralWorkspace")},
						Streams: []*armmonitor.KnownDataFlowStreams{
							to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftPerf),
							to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftSyslog),
							to.Ptr(armmonitor.KnownDataFlowStreamsMicrosoftWindowsEvent)},
					}},
				DataSources: &armmonitor.DataCollectionRuleDataSources{
					PerformanceCounters: []*armmonitor.PerfCounterDataSource{
						{
							Name: to.Ptr("cloudTeamCoreCounters"),
							CounterSpecifiers: []*string{
								to.Ptr("\\Processor(_Total)\\% Processor Time"),
								to.Ptr("\\Memory\\Committed Bytes"),
								to.Ptr("\\LogicalDisk(_Total)\\Free Megabytes"),
								to.Ptr("\\PhysicalDisk(_Total)\\Avg. Disk Queue Length")},
							SamplingFrequencyInSeconds: to.Ptr[int32](15),
							Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
								to.Ptr(armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf)},
						},
						{
							Name: to.Ptr("appTeamExtraCounters"),
							CounterSpecifiers: []*string{
								to.Ptr("\\Process(_Total)\\Thread Count")},
							SamplingFrequencyInSeconds: to.Ptr[int32](30),
							Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
								to.Ptr(armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf)},
						}},
					Syslog: []*armmonitor.SyslogDataSource{
						{
							Name: to.Ptr("cronSyslog"),
							FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
								to.Ptr(armmonitor.KnownSyslogDataSourceFacilityNamesCron)},
							LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
								to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsDebug),
								to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsCritical),
								to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsEmergency)},
							Streams: []*armmonitor.KnownSyslogDataSourceStreams{
								to.Ptr(armmonitor.KnownSyslogDataSourceStreamsMicrosoftSyslog)},
						},
						{
							Name: to.Ptr("syslogBase"),
							FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
								to.Ptr(armmonitor.KnownSyslogDataSourceFacilityNamesSyslog)},
							LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
								to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsAlert),
								to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsCritical),
								to.Ptr(armmonitor.KnownSyslogDataSourceLogLevelsEmergency)},
							Streams: []*armmonitor.KnownSyslogDataSourceStreams{
								to.Ptr(armmonitor.KnownSyslogDataSourceStreamsMicrosoftSyslog)},
						}},
					WindowsEventLogs: []*armmonitor.WindowsEventLogDataSource{
						{
							Name: to.Ptr("cloudSecurityTeamEvents"),
							Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
								to.Ptr(armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent)},
							XPathQueries: []*string{
								to.Ptr("Security!")},
						},
						{
							Name: to.Ptr("appTeam1AppEvents"),
							Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
								to.Ptr(armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent)},
							XPathQueries: []*string{
								to.Ptr("System![System[(Level = 1 or Level = 2 or Level = 3)]]"),
								to.Ptr("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]")},
						}},
				},
				Destinations: &armmonitor.DataCollectionRuleDestinations{
					LogAnalytics: []*armmonitor.LogAnalyticsDestination{
						{
							Name:                to.Ptr("centralWorkspace"),
							WorkspaceResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace"),
						}},
				},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
