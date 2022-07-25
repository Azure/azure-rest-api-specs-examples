const { HanaManagementClient } = require("@azure/arm-hanaonazure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches the Tags field of a SAP monitor for the specified subscription, resource group, and monitor name.
 *
 * @summary Patches the Tags field of a SAP monitor for the specified subscription, resource group, and monitor name.
 * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/SapMonitors_PatchTags.json
 */
async function updateTagsFieldOfASapMonitor() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const sapMonitorName = "mySapMonitor";
  const tagsParameter = { tags: { testkey: "testvalue" } };
  const credential = new DefaultAzureCredential();
  const client = new HanaManagementClient(credential, subscriptionId);
  const result = await client.sapMonitors.update(resourceGroupName, sapMonitorName, tagsParameter);
  console.log(result);
}

updateTagsFieldOfASapMonitor().catch(console.error);
