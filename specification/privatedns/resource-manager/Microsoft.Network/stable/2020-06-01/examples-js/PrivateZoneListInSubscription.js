const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function getPrivateDnsZoneBySubscription() {
  const subscriptionId = "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateZones.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getPrivateDnsZoneBySubscription().catch(console.error);
