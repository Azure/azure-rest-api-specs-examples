const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a DNS zone. Retrieves the zone properties, but not the record sets within the zone.
 *
 * @summary Gets a DNS zone. Retrieves the zone properties, but not the record sets within the zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetZone.json
 */
async function getZone() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.zones.get(resourceGroupName, zoneName);
  console.log(result);
}

getZone().catch(console.error);
