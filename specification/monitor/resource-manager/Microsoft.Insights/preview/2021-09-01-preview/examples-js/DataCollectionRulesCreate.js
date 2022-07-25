const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a data collection rule.
 *
 * @summary Creates or updates a data collection rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRulesCreate.json
 */
async function createOrUpdateDataCollectionRule() {
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = "myResourceGroup";
  const dataCollectionRuleName = "myCollectionRule";
  const body = {
    dataFlows: [
      {
        destinations: ["centralWorkspace"],
        streams: ["Microsoft-Perf", "Microsoft-Syslog", "Microsoft-WindowsEvent"],
      },
    ],
    dataSources: {
      performanceCounters: [
        {
          name: "cloudTeamCoreCounters",
          counterSpecifiers: [
            "Processor(_Total)% Processor Time",
            "MemoryCommitted Bytes",
            "LogicalDisk(_Total)Free Megabytes",
            "PhysicalDisk(_Total)Avg. Disk Queue Length",
          ],
          samplingFrequencyInSeconds: 15,
          streams: ["Microsoft-Perf"],
        },
        {
          name: "appTeamExtraCounters",
          counterSpecifiers: ["Process(_Total)Thread Count"],
          samplingFrequencyInSeconds: 30,
          streams: ["Microsoft-Perf"],
        },
      ],
      syslog: [
        {
          name: "cronSyslog",
          facilityNames: ["cron"],
          logLevels: ["Debug", "Critical", "Emergency"],
          streams: ["Microsoft-Syslog"],
        },
        {
          name: "syslogBase",
          facilityNames: ["syslog"],
          logLevels: ["Alert", "Critical", "Emergency"],
          streams: ["Microsoft-Syslog"],
        },
      ],
      windowsEventLogs: [
        {
          name: "cloudSecurityTeamEvents",
          streams: ["Microsoft-WindowsEvent"],
          xPathQueries: ["Security!"],
        },
        {
          name: "appTeam1AppEvents",
          streams: ["Microsoft-WindowsEvent"],
          xPathQueries: [
            "System![System[(Level = 1 or Level = 2 or Level = 3)]]",
            "Application!*[System[(Level = 1 or Level = 2 or Level = 3)]]",
          ],
        },
      ],
    },
    destinations: {
      logAnalytics: [
        {
          name: "centralWorkspace",
          workspaceResourceId:
            "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace",
        },
      ],
    },
    location: "eastus",
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionRules.create(
    resourceGroupName,
    dataCollectionRuleName,
    options
  );
  console.log(result);
}

createOrUpdateDataCollectionRule().catch(console.error);
