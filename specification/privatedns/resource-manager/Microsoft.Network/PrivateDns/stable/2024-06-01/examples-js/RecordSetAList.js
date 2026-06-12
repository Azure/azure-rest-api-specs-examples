const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the record sets of a specified type in a Private DNS zone.
 *
 * @summary lists the record sets of a specified type in a Private DNS zone.
 * x-ms-original-file: 2024-06-01/RecordSetAList.json
 */
async function getPrivateDNSZoneARecordSets() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.recordSets.listByType(
    "resourceGroup1",
    "privatezone1.com",
    "A",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
