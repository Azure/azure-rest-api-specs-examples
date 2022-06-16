const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
