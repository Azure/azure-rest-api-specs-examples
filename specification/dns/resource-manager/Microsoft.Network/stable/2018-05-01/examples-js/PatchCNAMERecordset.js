const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a record set within a DNS zone.
 *
 * @summary Updates a record set within a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/PatchCNAMERecordset.json
 */
async function patchCnameRecordset() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "record1";
  const recordType = "CNAME";
  const parameters = { metadata: { key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update(
    resourceGroupName,
    zoneName,
    relativeRecordSetName,
    recordType,
    parameters
  );
  console.log(result);
}

patchCnameRecordset().catch(console.error);
