const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a record set from a DNS zone. This operation cannot be undone.
 *
 * @summary Deletes a record set from a DNS zone. This operation cannot be undone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/DeletePTRRecordset.json
 */
async function deletePtrRecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "0.0.127.in-addr.arpa";
  const relativeRecordSetName = "1";
  const recordType = "PTR";
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

deletePtrRecordset().catch(console.error);
