const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function listPeeringServicesInASubscription() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peeringServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPeeringServicesInASubscription().catch(console.error);
