const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about an availability set.
 *
 * @summary Retrieves information about an availability set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/availabilitySetExamples/AvailabilitySets_Get_MaximumSet_Gen.json
 */
async function availabilitySetsGetMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const availabilitySetName = "aaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.availabilitySets.get(resourceGroupName, availabilitySetName);
  console.log(result);
}

availabilitySetsGetMaximumSetGen().catch(console.error);
