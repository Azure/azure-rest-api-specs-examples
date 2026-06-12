const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a record set within a Private DNS zone.
 *
 * @summary creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetPTRPut.json
 */
async function putPrivateDNSZonePTRRecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    "resourceGroup1",
    "0.0.127.in-addr.arpa",
    "PTR",
    "1",
    { metadata: { key1: "value1" }, ptrRecords: [{ ptrdname: "localhost" }], ttl: 3600 },
  );
  console.log(result);
}
