Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");
const { DefaultAzureCredential } = require("@azure/identity");

async function patchPrivateDnsZoneVirtualNetworkLink() {
async function patchPrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = "subscriptionId";
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const privateZoneName = "privatezone1.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const parameters = {
  const parameters = {
    registrationEnabled: true,
    registrationEnabled: true,
    tags: { key2: "value2" },
    tags: { key2: "value2" },
  };
  };
  const credential = new DefaultAzureCredential();
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginUpdateAndWait(
  const result = await client.virtualNetworkLinks.beginUpdateAndWait(
    resourceGroupName,
    resourceGroupName,
    privateZoneName,
    privateZoneName,
    virtualNetworkLinkName,
    virtualNetworkLinkName,
    parameters
    parameters
  );
  );
  console.log(result);
  console.log(result);
}
}

patchPrivateDnsZoneVirtualNetworkLink().catch(console.error);
patchPrivateDnsZoneVirtualNetworkLink().catch(console.error);
```
