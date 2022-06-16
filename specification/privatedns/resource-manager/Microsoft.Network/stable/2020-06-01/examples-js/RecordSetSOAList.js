const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function getPrivateDnsZoneSoaRecordSets() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "SOA";
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

getPrivateDnsZoneSoaRecordSets().catch(console.error);
