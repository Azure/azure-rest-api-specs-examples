const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a record set from a Private DNS zone. This operation cannot be undone.
 *
 * @summary deletes a record set from a Private DNS zone. This operation cannot be undone.
 * x-ms-original-file: 2024-06-01/RecordSetMXDelete.json
 */
async function deletePrivateDNSZoneMXRecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  await client.recordSets.delete("resourceGroup1", "privatezone1.com", "MX", "recordMX");
}
