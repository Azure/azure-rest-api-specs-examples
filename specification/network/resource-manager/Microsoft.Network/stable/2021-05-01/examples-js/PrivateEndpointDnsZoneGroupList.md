Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all private dns zone groups in a private endpoint.
 *
 * @summary Gets all private dns zone groups in a private endpoint.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointDnsZoneGroupList.json
 */
async function listPrivateEndpointsInResourceGroup() {
  const subscriptionId = "subId";
  const privateEndpointName = "testPe";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateDnsZoneGroups.list(privateEndpointName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPrivateEndpointsInResourceGroup().catch(console.error);
```
