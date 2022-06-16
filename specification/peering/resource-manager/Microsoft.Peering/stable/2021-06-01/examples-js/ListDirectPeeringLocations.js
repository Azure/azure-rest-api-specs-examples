const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
