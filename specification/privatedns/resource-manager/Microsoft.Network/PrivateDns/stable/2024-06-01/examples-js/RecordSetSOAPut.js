const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a record set within a Private DNS zone.
 *
 * @summary creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetSOAPut.json
 */
async function putPrivateDNSZoneSOARecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    "resourceGroup1",
    "privatezone1.com",
    "SOA",
    "@",
    {
      metadata: { key1: "value1" },
      soaRecord: {
        email: "azureprivatedns-hostmaster.microsoft.com",
        expireTime: 2419200,
        host: "azureprivatedns.net",
        minimumTtl: 300,
        refreshTime: 3600,
        retryTime: 300,
        serialNumber: 1,
      },
      ttl: 3600,
    },
  );
  console.log(result);
}
