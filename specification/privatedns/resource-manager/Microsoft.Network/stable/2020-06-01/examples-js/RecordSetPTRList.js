const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the record sets of a specified type in a Private DNS zone.
 *
 * @summary Lists the record sets of a specified type in a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetPTRList.json
 */
async function getPrivateDnsZonePtrRecordSets() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const recordType = "PTR";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recordSets.listByType(
    resourceGroupName,
    privateZoneName,
    recordType
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getPrivateDnsZonePtrRecordSets().catch(console.error);
