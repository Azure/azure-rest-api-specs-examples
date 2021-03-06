const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a Private DNS zone.
 *
 * @summary Creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetTXTPut.json
 */
async function putPrivateDnsZoneTxtRecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "TXT";
  const relativeRecordSetName = "recordTXT";
  const parameters = {
    metadata: { key1: "value1" },
    ttl: 3600,
    txtRecords: [{ value: ["string1", "string2"] }],
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
    parameters
  );
  console.log(result);
}

putPrivateDnsZoneTxtRecordSet().catch(console.error);
