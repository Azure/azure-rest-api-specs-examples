const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all availability sets in a subscription.
 *
 * @summary Lists all availability sets in a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/availabilitySetExamples/AvailabilitySet_ListBySubscription.json
 */
async function listAvailabilitySetsInASubscription() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const expand = "virtualMachines\\$ref";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilitySets.listBySubscription(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
