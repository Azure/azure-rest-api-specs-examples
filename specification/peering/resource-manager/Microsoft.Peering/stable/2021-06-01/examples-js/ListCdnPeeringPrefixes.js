const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the advertised prefixes for the specified peering location
 *
 * @summary Lists all of the advertised prefixes for the specified peering location
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListCdnPeeringPrefixes.json
 */
async function listAllTheCdnPeeringPrefixesAdvertisedAtAParticularPeeringLocation() {
  const subscriptionId = "subId";
  const peeringLocation = "peeringLocation0";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cdnPeeringPrefixes.list(peeringLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllTheCdnPeeringPrefixesAdvertisedAtAParticularPeeringLocation().catch(console.error);
