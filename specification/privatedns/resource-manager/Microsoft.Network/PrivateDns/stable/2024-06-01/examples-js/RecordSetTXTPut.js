const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a record set within a Private DNS zone.
 *
 * @summary creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetTXTPut.json
 */
async function putPrivateDNSZoneTXTRecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    "resourceGroup1",
    "privatezone1.com",
    "TXT",
    "recordTXT",
    { metadata: { key1: "value1" }, ttl: 3600, txtRecords: [{ value: ["string1", "string2"] }] },
  );
  console.log(result);
}
