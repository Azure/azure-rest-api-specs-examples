const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all proximity placement groups in a subscription.
 *
 * @summary Lists all proximity placement groups in a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListProximityPlacementGroupsInASubscription.json
 */
async function createAProximityPlacementGroup() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.proximityPlacementGroups.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

createAProximityPlacementGroup().catch(console.error);
