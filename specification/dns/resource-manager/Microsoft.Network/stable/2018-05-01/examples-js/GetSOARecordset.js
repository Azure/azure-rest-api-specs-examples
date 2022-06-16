const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a record set.
 *
 * @summary Gets a record set.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetSOARecordset.json
 */
async function getSoaRecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "@";
  const recordType = "SOA";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.get(
    resourceGroupName,
    zoneName,
    relativeRecordSetName,
    recordType
  );
  console.log(result);
}

getSoaRecordset().catch(console.error);
