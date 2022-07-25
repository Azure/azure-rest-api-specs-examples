const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets properties of a SAP monitor for the specified subscription, resource group, and resource name.
 *
 * @summary Gets properties of a SAP monitor for the specified subscription, resource group, and resource name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_Get.json
 */
async function getPropertiesOfASapMonitor() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.monitors.get(resourceGroupName, monitorName);
  console.log(result);
}

getPropertiesOfASapMonitor().catch(console.error);
