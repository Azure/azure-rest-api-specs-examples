const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a record set.
 *
 * @summary gets a record set.
 * x-ms-original-file: 2023-07-01-preview/GetSRVRecordset.json
 */
async function getSRVRecordset() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.get("rg1", "zone1", "record1", "SRV");
  console.log(result);
}
