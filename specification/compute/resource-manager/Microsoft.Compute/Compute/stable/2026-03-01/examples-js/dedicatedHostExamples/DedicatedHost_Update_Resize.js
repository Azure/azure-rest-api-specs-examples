const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a dedicated host .
 *
 * @summary update a dedicated host .
 * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHost_Update_Resize.json
 */
async function dedicatedHostUpdateResize() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHosts.update(
    "rgcompute",
    "aaaaaaaaa",
    "aaaaaaaaaaaaaaaaaaaaa",
    { sku: { name: "DSv3-Type1" } },
  );
  console.log(result);
}
