```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

async function putPrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const privateZoneName = "privatezone1.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const parameters = {
    location: "Global",
    registrationEnabled: false,
    tags: { key1: "value1" },
    virtualNetwork: {
      id: "/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateZoneName,
    virtualNetworkLinkName,
    parameters
  );
  console.log(result);
}

putPrivateDnsZoneVirtualNetworkLink().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.
