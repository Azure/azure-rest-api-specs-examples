const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a Private DNS zone.
 *
 * @summary Creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/RecordSetPTRPut.json
 */
async function putPrivateDnsZonePtrRecordSet() {
  const subscriptionId = process.env["PRIVATEDNS_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["PRIVATEDNS_RESOURCE_GROUP"] || "resourceGroup1";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const recordType = "PTR";
  const relativeRecordSetName = "1";
  const parameters = {
    metadata: { key1: "value1" },
    ptrRecords: [{ ptrdname: "localhost" }],
    ttl: 3600,
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
    parameters,
  );
  console.log(result);
}
