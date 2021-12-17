Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.3.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRulesCreate.json
func ExampleDataCollectionRulesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRulesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<data-collection-rule-name>",
		&armmonitor.DataCollectionRulesCreateOptions{Body: &armmonitor.DataCollectionRuleResource{
			Location: to.StringPtr("<location>"),
			Properties: &armmonitor.DataCollectionRuleResourceProperties{
				DataCollectionRule: armmonitor.DataCollectionRule{
					DataFlows: []*armmonitor.DataFlow{
						{
							Destinations: []*string{
								to.StringPtr("centralWorkspace")},
							Streams: []*armmonitor.KnownDataFlowStreams{
								armmonitor.KnownDataFlowStreamsMicrosoftPerf.ToPtr(),
								armmonitor.KnownDataFlowStreamsMicrosoftSyslog.ToPtr(),
								armmonitor.KnownDataFlowStreamsMicrosoftWindowsEvent.ToPtr()},
						}},
					DataSources: &armmonitor.DataCollectionRuleDataSources{
						DataSourcesSpec: armmonitor.DataSourcesSpec{
							PerformanceCounters: []*armmonitor.PerfCounterDataSource{
								{
									Name: to.StringPtr("<name>"),
									CounterSpecifiers: []*string{
										to.StringPtr("\\Processor(_Total)\\% Processor Time"),
										to.StringPtr("\\Memory\\Committed Bytes"),
										to.StringPtr("\\LogicalDisk(_Total)\\Free Megabytes"),
										to.StringPtr("\\PhysicalDisk(_Total)\\Avg. Disk Queue Length")},
									SamplingFrequencyInSeconds: to.Int32Ptr(15),
									Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
										armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf.ToPtr()},
								},
								{
									Name: to.StringPtr("<name>"),
									CounterSpecifiers: []*string{
										to.StringPtr("\\Process(_Total)\\Thread Count")},
									SamplingFrequencyInSeconds: to.Int32Ptr(30),
									Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
										armmonitor.KnownPerfCounterDataSourceStreamsMicrosoftPerf.ToPtr()},
								}},
							Syslog: []*armmonitor.SyslogDataSource{
								{
									Name: to.StringPtr("<name>"),
									FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
										armmonitor.KnownSyslogDataSourceFacilityNamesCron.ToPtr()},
									LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
										armmonitor.KnownSyslogDataSourceLogLevelsDebug.ToPtr(),
										armmonitor.KnownSyslogDataSourceLogLevelsCritical.ToPtr(),
										armmonitor.KnownSyslogDataSourceLogLevelsEmergency.ToPtr()},
									Streams: []*armmonitor.KnownSyslogDataSourceStreams{
										armmonitor.KnownSyslogDataSourceStreamsMicrosoftSyslog.ToPtr()},
								},
								{
									Name: to.StringPtr("<name>"),
									FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
										armmonitor.KnownSyslogDataSourceFacilityNamesSyslog.ToPtr()},
									LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
										armmonitor.KnownSyslogDataSourceLogLevelsAlert.ToPtr(),
										armmonitor.KnownSyslogDataSourceLogLevelsCritical.ToPtr(),
										armmonitor.KnownSyslogDataSourceLogLevelsEmergency.ToPtr()},
									Streams: []*armmonitor.KnownSyslogDataSourceStreams{
										armmonitor.KnownSyslogDataSourceStreamsMicrosoftSyslog.ToPtr()},
								}},
							WindowsEventLogs: []*armmonitor.WindowsEventLogDataSource{
								{
									Name: to.StringPtr("<name>"),
									Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
										armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent.ToPtr()},
									XPathQueries: []*string{
										to.StringPtr("Security!")},
								},
								{
									Name: to.StringPtr("<name>"),
									Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
										armmonitor.KnownWindowsEventLogDataSourceStreamsMicrosoftWindowsEvent.ToPtr()},
									XPathQueries: []*string{
										to.StringPtr("System![System[(Level = 1 or Level = 2 or Level = 3)]]"),
										to.StringPtr("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]")},
								}},
						},
					},
					Destinations: &armmonitor.DataCollectionRuleDestinations{
						DestinationsSpec: armmonitor.DestinationsSpec{
							LogAnalytics: []*armmonitor.LogAnalyticsDestination{
								{
									Name:                to.StringPtr("<name>"),
									WorkspaceResourceID: to.StringPtr("<workspace-resource-id>"),
								}},
						},
					},
				},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataCollectionRuleResource.ID: %s\n", *res.ID)
}
```
