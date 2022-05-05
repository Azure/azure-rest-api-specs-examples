Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletePrivateDnsZoneVirtualNetworkLink() {
async function deletePrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = "subscriptionId";
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const privateZoneName = "privatezone1.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginDeleteAndWait(
  const result = await client.virtualNetworkLinks.beginDeleteAndWait(
    resourceGroupName,
    resourceGroupName,
    privateZoneName,
    privateZoneName,
    virtualNetworkLinkName
    virtualNetworkLinkName
  );
  );
  console.log(result);
  console.log(result);
}
}

deletePrivateDnsZoneVirtualNetworkLink().catch(console.error);
deletePrivateDnsZoneVirtualNetworkLink().catch(console.error);
```
