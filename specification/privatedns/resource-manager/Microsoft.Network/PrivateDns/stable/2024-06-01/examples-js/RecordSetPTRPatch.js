const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a record set within a Private DNS zone.
 *
 * @summary updates a record set within a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetPTRPatch.json
 */
async function patchPrivateDNSZonePTRRecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update(
    "resourceGroup1",
    "0.0.127.in-addr.arpa",
    "PTR",
    "1",
    { metadata: { key2: "value2" } },
  );
  console.log(result);
}
