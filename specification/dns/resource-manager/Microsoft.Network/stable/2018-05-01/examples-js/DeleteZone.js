const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone.
 *
 * @summary Deletes a DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeleteZone.json
 */
async function deleteZone() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.zones.beginDeleteAndWait(resourceGroupName, zoneName);
  console.log(result);
}

deleteZone().catch(console.error);
