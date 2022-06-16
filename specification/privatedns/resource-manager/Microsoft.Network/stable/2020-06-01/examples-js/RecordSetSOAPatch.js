const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function patchPrivateDnsZoneSoaRecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const recordType = "SOA";
  const relativeRecordSetName = "@";
  const parameters = { metadata: { key2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.update(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
    parameters
  );
  console.log(result);
}

patchPrivateDnsZoneSoaRecordSet().catch(console.error);
