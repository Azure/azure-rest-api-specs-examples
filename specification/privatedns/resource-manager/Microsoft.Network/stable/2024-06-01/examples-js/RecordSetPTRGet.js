const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a record set.
 *
 * @summary Gets a record set.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/RecordSetPTRGet.json
 */
async function getPrivateDnsZonePtrRecordSet() {
  const subscriptionId = process.env["PRIVATEDNS_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["PRIVATEDNS_RESOURCE_GROUP"] || "resourceGroup1";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const recordType = "PTR";
  const relativeRecordSetName = "1";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.get(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
  );
  console.log(result);
}
