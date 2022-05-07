Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the specified ExpressRouteCrossConnection.
 *
 * @summary Update the specified ExpressRouteCrossConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCrossConnectionUpdate.json
 */
async function updateExpressRouteCrossConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "CrossConnection-SiliconValley";
  const crossConnectionName = "<circuitServiceKey>";
  const parameters = {
    serviceProviderProvisioningState: "NotProvisioned",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCrossConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    crossConnectionName,
    parameters
  );
  console.log(result);
}

updateExpressRouteCrossConnection().catch(console.error);
```
