const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a DNS zone.
 *
 * @summary Creates or updates a record set within a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateSOARecordset.json
 */
async function createSoaRecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "@";
  const recordType = "SOA";
  const parameters = {
    soaRecord: {
      email: "hostmaster.contoso.com",
      expireTime: 2419200,
      host: "ns1.contoso.com",
      minimumTtl: 300,
      refreshTime: 3600,
      retryTime: 300,
      serialNumber: 1,
    },
    ttl: 3600,
    metadata: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    zoneName,
    relativeRecordSetName,
    recordType,
    parameters
  );
  console.log(result);
}

createSoaRecordset().catch(console.error);
