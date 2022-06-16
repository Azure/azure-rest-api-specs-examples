const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Private DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone. Private DNS zone cannot be deleted unless all virtual network links to it are removed.
 *
 * @summary Deletes a Private DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone. Private DNS zone cannot be deleted unless all virtual network links to it are removed.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZoneDelete.json
 */
async function deletePrivateDnsZone() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.privateZones.beginDeleteAndWait(resourceGroupName, privateZoneName);
  console.log(result);
}

deletePrivateDnsZone().catch(console.error);
