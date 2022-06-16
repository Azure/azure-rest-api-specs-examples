const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a DNS zone.
 *
 * @summary Creates or updates a record set within a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateARecordset.json
 */
async function createARecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "record1";
  const recordType = "A";
  const parameters = {
    aRecords: [{ ipv4Address: "127.0.0.1" }],
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

createARecordset().catch(console.error);
