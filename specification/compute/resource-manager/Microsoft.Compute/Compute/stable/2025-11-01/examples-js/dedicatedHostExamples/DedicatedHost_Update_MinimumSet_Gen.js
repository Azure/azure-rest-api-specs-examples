const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a dedicated host .
 *
 * @summary update a dedicated host .
 * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_Update_MinimumSet_Gen.json
 */
async function dedicatedHostUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHosts.update(
    "rgcompute",
    "aa",
    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
    {},
  );
  console.log(result);
}
