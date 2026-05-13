const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all of the available peering locations for the specified kind of peering.
 *
 * @summary lists all of the available peering locations for the specified kind of peering.
 * x-ms-original-file: 2025-05-01/ListExchangePeeringLocations.json
 */
async function listExchangePeeringLocations() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subId";
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.peeringLocations.list("Exchange")) {
    resArray.push(item);
  }

  console.log(resArray);
}
