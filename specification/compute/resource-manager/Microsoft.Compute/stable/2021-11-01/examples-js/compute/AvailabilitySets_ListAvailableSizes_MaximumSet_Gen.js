const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
 *
 * @summary Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/AvailabilitySets_ListAvailableSizes_MaximumSet_Gen.json
 */
async function availabilitySetsListAvailableSizesMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const availabilitySetName = "aaaaaaaaaaaaaaaaaaaa";
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

availabilitySetsListAvailableSizesMaximumSetGen().catch(console.error);
