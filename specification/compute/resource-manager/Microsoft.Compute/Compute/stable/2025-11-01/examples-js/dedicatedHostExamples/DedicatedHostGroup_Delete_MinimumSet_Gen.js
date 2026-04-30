const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a dedicated host group.
 *
 * @summary delete a dedicated host group.
 * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHostGroup_Delete_MinimumSet_Gen.json
 */
async function dedicatedHostGroupDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.dedicatedHostGroups.delete("rgcompute", "aaaa");
}
