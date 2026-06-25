const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about an availability set.
 *
 * @summary retrieves information about an availability set.
 * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Get_MaximumSet_Gen.json
 */
async function availabilitySetGetMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.get("rgcompute", "aaaaaaaaaaaa");
  console.log(result);
}
