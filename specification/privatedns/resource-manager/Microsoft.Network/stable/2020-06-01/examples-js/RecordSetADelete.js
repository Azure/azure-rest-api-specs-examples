const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletePrivateDnsZoneARecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "A";
  const relativeRecordSetName = "recordA";
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

deletePrivateDnsZoneARecordSet().catch(console.error);
