const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function patchPrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const parameters = {
    registrationEnabled: true,
    tags: { key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginUpdateAndWait(
    resourceGroupName,
    privateZoneName,
    virtualNetworkLinkName,
    parameters
  );
  console.log(result);
}

patchPrivateDnsZoneVirtualNetworkLink().catch(console.error);
