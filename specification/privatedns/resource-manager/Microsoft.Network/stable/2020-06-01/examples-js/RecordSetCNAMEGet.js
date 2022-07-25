const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a record set.
 *
 * @summary Gets a record set.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetCNAMEGet.json
 */
async function getPrivateDnsZoneCnameRecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "CNAME";
  const relativeRecordSetName = "recordCNAME";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.get(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName
  );
  console.log(result);
}

getPrivateDnsZoneCnameRecordSet().catch(console.error);
