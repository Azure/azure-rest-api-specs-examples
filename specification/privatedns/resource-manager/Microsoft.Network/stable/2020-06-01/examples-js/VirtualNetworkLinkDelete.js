const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletePrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginDeleteAndWait(
    resourceGroupName,
    privateZoneName,
    virtualNetworkLinkName
  );
  console.log(result);
}

deletePrivateDnsZoneVirtualNetworkLink().catch(console.error);
