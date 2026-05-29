const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the record sets of a specified type in a DNS zone.
 *
 * @summary lists the record sets of a specified type in a DNS zone.
 * x-ms-original-file: 2023-07-01-preview/ListCaaRecordset.json
 */
async function listCAARecordsets() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new DnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.recordSets.listByType("rg1", "zone1", "CAA")) {
    resArray.push(item);
  }

  console.log(resArray);
}
