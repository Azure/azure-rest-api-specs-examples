const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group, and resource name.
 *
 * @summary Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_List.json
 */
async function getPropertiesOfASapMonitor() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "mySapMonitor";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapLandscapeMonitorOperations.list(resourceGroupName, monitorName);
  console.log(result);
}
