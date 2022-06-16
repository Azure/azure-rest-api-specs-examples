const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

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
