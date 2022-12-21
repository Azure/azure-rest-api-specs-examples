const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the virtual machines under the specified subscription for the specified location.
 *
 * @summary Gets all the virtual machines under the specified subscription for the specified location.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/ListVirtualMachinesInASubscriptionByLocation.json
 */
async function listsAllTheVirtualMachinesUnderTheSpecifiedSubscriptionForTheSpecifiedLocation() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachines.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
