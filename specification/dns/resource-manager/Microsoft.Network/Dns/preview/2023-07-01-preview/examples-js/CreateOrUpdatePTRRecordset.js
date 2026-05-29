const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a record set within a DNS zone. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
 *
 * @summary creates or updates a record set within a DNS zone. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
 * x-ms-original-file: 2023-07-01-preview/CreateOrUpdatePTRRecordset.json
 */
async function createPTRRecordset() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate("rg1", "0.0.127.in-addr.arpa", "1", "PTR", {
    ptrRecords: [{ ptrdname: "localhost" }],
    ttl: 3600,
    metadata: { key1: "value1" },
  });
  console.log(result);
}
