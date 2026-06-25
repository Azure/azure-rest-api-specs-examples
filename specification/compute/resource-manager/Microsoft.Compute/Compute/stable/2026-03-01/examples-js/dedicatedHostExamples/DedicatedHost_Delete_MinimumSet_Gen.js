const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a dedicated host.
 *
 * @summary delete a dedicated host.
 * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHost_Delete_MinimumSet_Gen.json
 */
async function dedicatedHostDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.dedicatedHosts.delete("rgcompute", "aaaaaaaaaaaaaaa", "aaaaa");
}
