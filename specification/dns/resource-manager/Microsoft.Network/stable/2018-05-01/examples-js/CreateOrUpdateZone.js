const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a DNS zone. Does not modify DNS records within the zone.
 *
 * @summary Creates or updates a DNS zone. Does not modify DNS records within the zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateZone.json
 */
async function createZone() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const parameters = { location: "Global", tags: { key1: "value1" } };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.zones.createOrUpdate(resourceGroupName, zoneName, parameters);
  console.log(result);
}

createZone().catch(console.error);
