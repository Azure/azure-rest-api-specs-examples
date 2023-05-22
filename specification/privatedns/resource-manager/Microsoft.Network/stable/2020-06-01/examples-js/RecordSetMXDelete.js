const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a record set from a Private DNS zone. This operation cannot be undone.
 *
 * @summary Deletes a record set from a Private DNS zone. This operation cannot be undone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetMXDelete.json
 */
async function deletePrivateDnsZoneMxRecordSet() {
  const subscriptionId = process.env["PRIVATEDNS_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["PRIVATEDNS_RESOURCE_GROUP"] || "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "MX";
  const relativeRecordSetName = "recordMX";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.delete(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName
  );
  console.log(result);
}
