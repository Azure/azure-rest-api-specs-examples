const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available peering locations for the specified kind of peering.
 *
 * @summary Lists all of the available peering locations for the specified kind of peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListDirectPeeringLocations.json
 */
async function listDirectPeeringLocations() {
  const subscriptionId = "subId";
  const kind = "Direct";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peeringLocations.list(kind)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDirectPeeringLocations().catch(console.error);
