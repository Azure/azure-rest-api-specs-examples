const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the Tags field of a SAP monitor for the specified subscription, resource group, and SAP monitor name.
 *
 * @summary Patches the Tags field of a SAP monitor for the specified subscription, resource group, and SAP monitor name.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_PatchTags.json
 */
async function updateTagsFieldOfASapMonitor() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "myResourceGroup";
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
