const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about a dedicated host group.
 *
 * @summary retrieves information about a dedicated host group.
 * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
 */
async function createAnUltraSSDEnabledDedicatedHostGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscriptionId}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.get("myResourceGroup", "myDedicatedHostGroup");
  console.log(result);
}
