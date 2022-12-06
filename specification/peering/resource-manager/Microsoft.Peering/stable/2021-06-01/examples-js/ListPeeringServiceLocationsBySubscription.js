const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available locations for peering service.
 *
 * @summary Lists all of the available locations for peering service.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListPeeringServiceLocationsBySubscription.json
 */
async function listPeeringServiceLocations() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peeringServiceLocations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPeeringServiceLocations().catch(console.error);
