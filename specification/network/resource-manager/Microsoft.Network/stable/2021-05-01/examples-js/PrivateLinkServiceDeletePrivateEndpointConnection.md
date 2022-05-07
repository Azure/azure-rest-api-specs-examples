Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete private end point connection for a private link service in a subscription.
 *
 * @summary Delete private end point connection for a private link service in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceDeletePrivateEndpointConnection.json
 */
async function deletePrivateEndPointConnectionForAPrivateLinkService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const serviceName = "testPls";
  const peConnectionName = "testPlePeConnection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.beginDeletePrivateEndpointConnectionAndWait(
    resourceGroupName,
    serviceName,
    peConnectionName
  );
  console.log(result);
}

deletePrivateEndPointConnectionForAPrivateLinkService().catch(console.error);
```
