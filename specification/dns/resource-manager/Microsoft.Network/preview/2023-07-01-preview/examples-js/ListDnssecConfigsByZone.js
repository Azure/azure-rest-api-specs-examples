const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the DNSSEC configurations in a DNS zone.
 *
 * @summary Lists the DNSSEC configurations in a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/ListDnssecConfigsByZone.json
 */
async function listDnssecConfigs() {
  const subscriptionId = process.env["DNS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["DNS_RESOURCE_GROUP"] || "rg1";
  const zoneName = "zone1";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dnssecConfigs.listByDnsZone(resourceGroupName, zoneName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
