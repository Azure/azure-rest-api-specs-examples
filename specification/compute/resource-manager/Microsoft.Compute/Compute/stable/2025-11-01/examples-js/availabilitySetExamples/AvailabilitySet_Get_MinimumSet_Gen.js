const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieves information about an availability set.
 *
 * @summary retrieves information about an availability set.
 * x-ms-original-file: 2025-11-01/availabilitySetExamples/AvailabilitySet_Get_MinimumSet_Gen.json
 */
async function availabilitySetGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.get("rgcompute", "aaaaaaaaaaaaaaaaaaaa");
  console.log(result);
}
