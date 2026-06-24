const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a dedicated host group.
 *
 * @summary delete a dedicated host group.
 * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Delete_MaximumSet_Gen.json
 */
async function dedicatedHostGroupDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.dedicatedHostGroups.delete("rgcompute", "a");
}
