const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update an availability set.
 *
 * @summary update an availability set.
 * x-ms-original-file: 2025-11-01/availabilitySetExamples/AvailabilitySet_Update_MinimumSet_Gen.json
 */
async function availabilitySetUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.update("rgcompute", "aaaaaaaaaaaaaaaaaaaa", {});
  console.log(result);
}
