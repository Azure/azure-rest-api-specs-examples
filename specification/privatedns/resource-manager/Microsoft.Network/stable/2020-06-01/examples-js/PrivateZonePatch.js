const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Private DNS zone. Does not modify virtual network links or DNS records within the zone.
 *
 * @summary Updates a Private DNS zone. Does not modify virtual network links or DNS records within the zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZonePatch.json
 */
async function patchPrivateDnsZone() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const parameters = { tags: { key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.privateZones.beginUpdateAndWait(
    resourceGroupName,
    privateZoneName,
    parameters
  );
  console.log(result);
}

patchPrivateDnsZone().catch(console.error);
