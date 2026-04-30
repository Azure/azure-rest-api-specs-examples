const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a dedicated host.
 *
 * @summary delete a dedicated host.
 * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_Delete_MaximumSet_Gen.json
 */
async function dedicatedHostDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.dedicatedHosts.delete("rgcompute", "aaaaaa", "aaaaaaaaaaaaaaa");
}
