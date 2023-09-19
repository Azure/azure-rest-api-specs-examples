const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
 *
 * @summary Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/availabilitySetExamples/AvailabilitySet_ListAvailableSizes_MinimumSet_Gen.json
 */
async function availabilitySetListAvailableSizesMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const availabilitySetName = "aa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilitySets.listAvailableSizes(
    resourceGroupName,
    availabilitySetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
