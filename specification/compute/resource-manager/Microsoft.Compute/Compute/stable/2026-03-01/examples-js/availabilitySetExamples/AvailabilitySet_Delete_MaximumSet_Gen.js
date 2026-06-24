const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete an availability set.
 *
 * @summary delete an availability set.
 * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Delete_MaximumSet_Gen.json
 */
async function availabilitySetDeleteMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.availabilitySets.delete("rgcompute", "aaaaaaaaaaaaaaaaaaaa");
}
