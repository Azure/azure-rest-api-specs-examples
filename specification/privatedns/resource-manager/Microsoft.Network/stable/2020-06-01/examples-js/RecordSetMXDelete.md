Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a record set from a Private DNS zone. This operation cannot be undone.
 *
 * @summary Deletes a record set from a Private DNS zone. This operation cannot be undone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetMXDelete.json
 */
async function deletePrivateDnsZoneMxRecordSet() {
async function deletePrivateDnsZoneMxRecordSet() {
  const subscriptionId = "subscriptionId";
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const privateZoneName = "privatezone1.com";
  const recordType = "MX";
  const recordType = "MX";
  const relativeRecordSetName = "recordMX";
  const relativeRecordSetName = "recordMX";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.delete(
  const result = await client.recordSets.delete(
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

deletePrivateDnsZoneMxRecordSet().catch(console.error);
deletePrivateDnsZoneMxRecordSet().catch(console.error);
```
