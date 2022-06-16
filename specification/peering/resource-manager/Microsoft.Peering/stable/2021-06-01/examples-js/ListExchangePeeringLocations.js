const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function listExchangePeeringLocations() {
  const subscriptionId = "subId";
  const kind = "Exchange";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peeringLocations.list(kind)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExchangePeeringLocations().catch(console.error);
