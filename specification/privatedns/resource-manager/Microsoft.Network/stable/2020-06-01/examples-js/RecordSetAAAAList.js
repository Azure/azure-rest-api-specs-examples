const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the record sets of a specified type in a Private DNS zone.
 *
 * @summary Lists the record sets of a specified type in a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetAAAAList.json
 */
async function getPrivateDnsZoneAaaaRecordSets() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "AAAA";
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

getPrivateDnsZoneAaaaRecordSets().catch(console.error);
