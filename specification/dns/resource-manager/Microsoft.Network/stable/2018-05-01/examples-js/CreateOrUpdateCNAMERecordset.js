const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a DNS zone.
 *
 * @summary Creates or updates a record set within a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateCNAMERecordset.json
 */
async function createCnameRecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "record1";
  const recordType = "CNAME";
  const parameters = {
    cnameRecord: { cname: "contoso.com" },
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

createCnameRecordset().catch(console.error);
