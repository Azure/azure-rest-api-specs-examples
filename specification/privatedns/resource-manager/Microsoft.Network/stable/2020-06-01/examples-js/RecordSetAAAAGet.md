Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

async function getPrivateDnsZoneAaaaRecordSet() {
async function getPrivateDnsZoneAaaaRecordSet() {
  const subscriptionId = "subscriptionId";
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const privateZoneName = "privatezone1.com";
  const recordType = "AAAA";
  const recordType = "AAAA";
  const relativeRecordSetName = "recordAAAA";
  const relativeRecordSetName = "recordAAAA";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.get(
  const result = await client.recordSets.get(
    resourceGroupName,
    resourceGroupName,
    privateZoneName,
    privateZoneName,
    recordType,
    recordType,
    relativeRecordSetName
    relativeRecordSetName
  );
  );
  console.log(result);
  console.log(result);
}
}

getPrivateDnsZoneAaaaRecordSet().catch(console.error);
getPrivateDnsZoneAaaaRecordSet().catch(console.error);
```
