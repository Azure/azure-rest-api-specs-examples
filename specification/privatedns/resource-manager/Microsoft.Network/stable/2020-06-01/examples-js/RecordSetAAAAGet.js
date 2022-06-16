const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function getPrivateDnsZoneAaaaRecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "AAAA";
  const relativeRecordSetName = "recordAAAA";
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

getPrivateDnsZoneAaaaRecordSet().catch(console.error);
