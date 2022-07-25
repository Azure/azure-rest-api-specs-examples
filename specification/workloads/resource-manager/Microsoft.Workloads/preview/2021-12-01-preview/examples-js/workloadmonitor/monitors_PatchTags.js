const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the Tags field of a SAP monitor for the specified subscription, resource group, and SAP monitor name.
 *
 * @summary Patches the Tags field of a SAP monitor for the specified subscription, resource group, and SAP monitor name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/monitors_PatchTags.json
 */
async function updateTagsFieldOfASapMonitor() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "mySapMonitor";
  const body = {
    identity: { type: "None" },
    tags: { testkey: "testvalue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.monitors.update(resourceGroupName, monitorName, body);
  console.log(result);
}

updateTagsFieldOfASapMonitor().catch(console.error);
