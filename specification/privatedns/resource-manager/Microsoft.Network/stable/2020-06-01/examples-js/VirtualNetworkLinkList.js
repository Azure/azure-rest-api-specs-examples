const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function getPrivateDnsZoneVirtualNetworkLinks() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkLinks.list(resourceGroupName, privateZoneName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getPrivateDnsZoneVirtualNetworkLinks().catch(console.error);
