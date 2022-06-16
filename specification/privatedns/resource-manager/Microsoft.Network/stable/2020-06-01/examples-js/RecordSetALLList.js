const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all record sets in a Private DNS zone.
 *
 * @summary Lists all record sets in a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetALLList.json
 */
async function getPrivateDnsZoneAllRecordSets() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recordSets.list(resourceGroupName, privateZoneName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getPrivateDnsZoneAllRecordSets().catch(console.error);
