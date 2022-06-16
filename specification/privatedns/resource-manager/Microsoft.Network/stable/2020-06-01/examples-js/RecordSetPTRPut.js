const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function putPrivateDnsZonePtrRecordSet() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const recordType = "PTR";
  const relativeRecordSetName = "1";
  const parameters = {
    metadata: { key1: "value1" },
    ptrRecords: [{ ptrdname: "localhost" }],
    ttl: 3600,
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    privateZoneName,
    recordType,
    relativeRecordSetName,
    parameters
  );
  console.log(result);
}

putPrivateDnsZonePtrRecordSet().catch(console.error);
