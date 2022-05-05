Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

async function putPrivateDnsZonePtrRecordSet() {
async function putPrivateDnsZonePtrRecordSet() {
  const subscriptionId = "subscriptionId";
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const recordType = "PTR";
  const recordType = "PTR";
  const relativeRecordSetName = "1";
  const relativeRecordSetName = "1";
  const parameters = {
  const parameters = {
    metadata: { key1: "value1" },
    metadata: { key1: "value1" },
    ptrRecords: [{ ptrdname: "localhost" }],
    ptrRecords: [{ ptrdname: "localhost" }],
    ttl: 3600,
    ttl: 3600,
  };
  };
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    resourceGroupName,
    privateZoneName,
    privateZoneName,
    recordType,
    recordType,
    relativeRecordSetName,
    relativeRecordSetName,
    parameters
    parameters
  );
  );
  console.log(result);
  console.log(result);
}
}

putPrivateDnsZonePtrRecordSet().catch(console.error);
putPrivateDnsZonePtrRecordSet().catch(console.error);
```
