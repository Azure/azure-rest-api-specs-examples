const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a record set within a DNS zone.
 *
 * @summary updates a record set within a DNS zone.
 * x-ms-original-file: 2023-07-01-preview/PatchTLSARecordset.json
 */
async function patchTlsaRecordset() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update("rg1", "zone1", "record1", "TLSA", {
    metadata: { key2: "value2" },
  });
  console.log(result);
}
