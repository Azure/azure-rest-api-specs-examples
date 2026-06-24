const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update an dedicated host group.
 *
 * @summary update an dedicated host group.
 * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Update_MinimumSet_Gen.json
 */
async function dedicatedHostGroupUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.update("rgcompute", "aaaaaaaaaaa", {});
  console.log(result);
}
