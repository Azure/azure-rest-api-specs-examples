const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a DNS zone. Does not modify DNS records within the zone.
 *
 * @summary Updates a DNS zone. Does not modify DNS records within the zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/PatchZone.json
 */
async function patchZone() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const parameters = { tags: { key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.zones.update(resourceGroupName, zoneName, parameters);
  console.log(result);
}

patchZone().catch(console.error);
