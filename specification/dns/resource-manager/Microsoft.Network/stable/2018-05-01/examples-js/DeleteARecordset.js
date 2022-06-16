const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a record set from a DNS zone. This operation cannot be undone.
 *
 * @summary Deletes a record set from a DNS zone. This operation cannot be undone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeleteARecordset.json
 */
async function deleteNsRecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "record1";
  const recordType = "A";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.delete(
    resourceGroupName,
    zoneName,
    relativeRecordSetName,
    recordType
  );
  console.log(result);
}

deleteNsRecordset().catch(console.error);
