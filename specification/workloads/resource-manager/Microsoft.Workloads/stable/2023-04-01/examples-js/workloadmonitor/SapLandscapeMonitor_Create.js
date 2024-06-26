const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a SAP Landscape Monitor Dashboard for the specified subscription, resource group, and resource name.
 *
 * @summary Creates a SAP Landscape Monitor Dashboard for the specified subscription, resource group, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_Create.json
 */
async function createForSapLandscapeMonitorDashboard() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "mySapMonitor";
  const sapLandscapeMonitorParameter = {
    grouping: {
      landscape: [{ name: "Prod", topSid: ["SID1", "SID2"] }],
      sapApplication: [{ name: "ERP1", topSid: ["SID1", "SID2"] }],
    },
    topMetricsThresholds: [{ name: "Instance Availability", green: 90, red: 50, yellow: 75 }],
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapLandscapeMonitorOperations.create(
    resourceGroupName,
    monitorName,
    sapLandscapeMonitorParameter
  );
  console.log(result);
}
