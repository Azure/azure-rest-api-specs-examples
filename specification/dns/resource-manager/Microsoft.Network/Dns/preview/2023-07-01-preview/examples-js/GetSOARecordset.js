const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a record set.
 *
 * @summary gets a record set.
 * x-ms-original-file: 2023-07-01-preview/GetSOARecordset.json
 */
async function getSOARecordset() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.get("rg1", "zone1", "@", "SOA");
  console.log(result);
}
