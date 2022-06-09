```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the backend health of the specified application gateway in a resource group.
 *
 * @summary Gets the backend health of the specified application gateway in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ApplicationGatewayBackendHealthGet.json
 */
async function getBackendHealth() {
  const subscriptionId = "subid";
  const resourceGroupName = "appgw";
  const applicationGatewayName = "appgw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGateways.beginBackendHealthAndWait(
    resourceGroupName,
    applicationGatewayName
  );
  console.log(result);
}

getBackendHealth().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
