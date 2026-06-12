const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a record set within a Private DNS zone.
 *
 * @summary creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetMXPut.json
 */
async function putPrivateDNSZoneMXRecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    "resourceGroup1",
    "privatezone1.com",
    "MX",
    "recordMX",
    {
      metadata: { key1: "value1" },
      mxRecords: [{ exchange: "mail.privatezone1.com", preference: 0 }],
      ttl: 3600,
    },
  );
  console.log(result);
}
