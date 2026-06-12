const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a record set within a Private DNS zone.
 *
 * @summary updates a record set within a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetTXTPatch.json
 */
async function patchPrivateDNSZoneTXTRecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update(
    "resourceGroup1",
    "privatezone1.com",
    "TXT",
    "recordTXT",
    { metadata: { key2: "value2" } },
  );
  console.log(result);
}
