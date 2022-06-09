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
		&armmonitor.DataCollectionRulesClientCreateOptions{Body: &armmonitor.DataCollectionRuleResource{
			Location: to.StringPtr("<location>"),
			Properties: &armmonitor.DataCollectionRuleResourceProperties{
				DataFlows: []*armmonitor.DataFlow{
					{
						Destinations: []*string{
							to.StringPtr("centralWorkspace")},
						Streams: []*armmonitor.KnownDataFlowStreams{
							armmonitor.KnownDataFlowStreams("Microsoft-Perf").ToPtr(),
							armmonitor.KnownDataFlowStreams("Microsoft-Syslog").ToPtr(),
							armmonitor.KnownDataFlowStreams("Microsoft-WindowsEvent").ToPtr()},
					}},
				DataSources: &armmonitor.DataCollectionRuleDataSources{
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
								armmonitor.KnownPerfCounterDataSourceStreams("Microsoft-Perf").ToPtr()},
						},
						{
							Name: to.StringPtr("<name>"),
							CounterSpecifiers: []*string{
								to.StringPtr("\\Process(_Total)\\Thread Count")},
							SamplingFrequencyInSeconds: to.Int32Ptr(30),
							Streams: []*armmonitor.KnownPerfCounterDataSourceStreams{
								armmonitor.KnownPerfCounterDataSourceStreams("Microsoft-Perf").ToPtr()},
						}},
					Syslog: []*armmonitor.SyslogDataSource{
						{
							Name: to.StringPtr("<name>"),
							FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
								armmonitor.KnownSyslogDataSourceFacilityNames("cron").ToPtr()},
							LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
								armmonitor.KnownSyslogDataSourceLogLevels("Debug").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Critical").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Emergency").ToPtr()},
							Streams: []*armmonitor.KnownSyslogDataSourceStreams{
								armmonitor.KnownSyslogDataSourceStreams("Microsoft-Syslog").ToPtr()},
						},
						{
							Name: to.StringPtr("<name>"),
							FacilityNames: []*armmonitor.KnownSyslogDataSourceFacilityNames{
								armmonitor.KnownSyslogDataSourceFacilityNames("syslog").ToPtr()},
							LogLevels: []*armmonitor.KnownSyslogDataSourceLogLevels{
								armmonitor.KnownSyslogDataSourceLogLevels("Alert").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Critical").ToPtr(),
								armmonitor.KnownSyslogDataSourceLogLevels("Emergency").ToPtr()},
							Streams: []*armmonitor.KnownSyslogDataSourceStreams{
								armmonitor.KnownSyslogDataSourceStreams("Microsoft-Syslog").ToPtr()},
						}},
					WindowsEventLogs: []*armmonitor.WindowsEventLogDataSource{
						{
							Name: to.StringPtr("<name>"),
							Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
								armmonitor.KnownWindowsEventLogDataSourceStreams("Microsoft-WindowsEvent").ToPtr()},
							XPathQueries: []*string{
								to.StringPtr("Security!")},
						},
						{
							Name: to.StringPtr("<name>"),
							Streams: []*armmonitor.KnownWindowsEventLogDataSourceStreams{
								armmonitor.KnownWindowsEventLogDataSourceStreams("Microsoft-WindowsEvent").ToPtr()},
							XPathQueries: []*string{
								to.StringPtr("System![System[(Level = 1 or Level = 2 or Level = 3)]]"),
								to.StringPtr("Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]")},
						}},
				},
				Destinations: &armmonitor.DataCollectionRuleDestinations{
					LogAnalytics: []*armmonitor.LogAnalyticsDestination{
						{
							Name:                to.StringPtr("<name>"),
							WorkspaceResourceID: to.StringPtr("<workspace-resource-id>"),
						}},
				},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataCollectionRulesClientCreateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.
