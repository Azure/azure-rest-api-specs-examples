const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a record set within a Private DNS zone.
 *
 * @summary Updates a record set within a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetSOAPatch.json
 */
async function patchPrivateDnsZoneSoaRecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "SOA";
  const relativeRecordSetName = "@";
  const parameters = { metadata: { key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
    parameters
  );
  console.log(result);
}

patchPrivateDnsZoneSoaRecordSet().catch(console.error);
