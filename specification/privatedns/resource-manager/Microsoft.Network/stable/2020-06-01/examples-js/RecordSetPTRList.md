Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the record sets of a specified type in a Private DNS zone.
 *
 * @summary Lists the record sets of a specified type in a Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetPTRList.json
 */
async function getPrivateDnsZonePtrRecordSets() {
async function getPrivateDnsZonePtrRecordSets() {
  const subscriptionId = "subscriptionId";
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const privateZoneName = "0.0.127.in-addr.arpa";
  const recordType = "PTR";
  const recordType = "PTR";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  const resArray = new Array();
  for await (let item of client.recordSets.listByType(
  for await (let item of client.recordSets.listByType(
    resourceGroupName,
    resourceGroupName,
    privateZoneName,
    privateZoneName,
    recordType
    recordType
  )) {
  )) {
    resArray.push(item);
    resArray.push(item);
  }
  }
  console.log(resArray);
  console.log(resArray);
}
}

getPrivateDnsZonePtrRecordSets().catch(console.error);
getPrivateDnsZonePtrRecordSets().catch(console.error);
```
