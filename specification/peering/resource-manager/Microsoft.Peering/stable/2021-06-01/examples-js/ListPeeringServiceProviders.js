const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
