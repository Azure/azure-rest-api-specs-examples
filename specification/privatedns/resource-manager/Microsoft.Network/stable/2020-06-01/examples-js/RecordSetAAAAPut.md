Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a Private DNS zone.
 *
 * @summary Creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetAAAAPut.json
 */
async function putPrivateDnsZoneAaaaRecordSet() {
async function putPrivateDnsZoneAaaaRecordSet() {
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
  const parameters = {
  const parameters = {
    aaaaRecords: [{ ipv6Address: "::1" }],
    aaaaRecords: [{ ipv6Address: "::1" }],
    metadata: { key1: "value1" },
    metadata: { key1: "value1" },
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

putPrivateDnsZoneAaaaRecordSet().catch(console.error);
putPrivateDnsZoneAaaaRecordSet().catch(console.error);
```
