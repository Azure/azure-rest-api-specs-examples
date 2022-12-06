const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available peering service locations for the specified kind of peering.
 *
 * @summary Lists all of the available peering service locations for the specified kind of peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListPeeringServiceProviders.json
 */
async function listPeeringServiceProviders() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peeringServiceProviders.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPeeringServiceProviders().catch(console.error);
