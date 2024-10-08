const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a Private DNS zone.
 *
 * @summary Creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/RecordSetSOAPut.json
 */
async function putPrivateDnsZoneSoaRecordSet() {
  const subscriptionId = process.env["PRIVATEDNS_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["PRIVATEDNS_RESOURCE_GROUP"] || "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "SOA";
  const relativeRecordSetName = "@";
  const parameters = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
    parameters,
  );
  console.log(result);
}
