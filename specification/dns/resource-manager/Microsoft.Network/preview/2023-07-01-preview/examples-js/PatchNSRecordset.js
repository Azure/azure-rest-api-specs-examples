const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a record set within a DNS zone.
 *
 * @summary Updates a record set within a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/PatchNSRecordset.json
 */
async function patchNsRecordset() {
  const subscriptionId = process.env["DNS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["DNS_RESOURCE_GROUP"] || "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "record1";
  const recordType = "NS";
  const parameters = { metadata: { key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update(
    resourceGroupName,
    zoneName,
    relativeRecordSetName,
    recordType,
    parameters,
  );
  console.log(result);
}
