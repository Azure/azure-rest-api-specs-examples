const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the legacy peerings under the given subscription matching the specified kind and location.
 *
 * @summary Lists all of the legacy peerings under the given subscription matching the specified kind and location.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListLegacyPeerings.json
 */
async function listLegacyPeerings() {
  const subscriptionId = "subId";
  const peeringLocation = "peeringLocation0";
  const kind = "Exchange";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.legacyPeerings.list(peeringLocation, kind)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listLegacyPeerings().catch(console.error);
