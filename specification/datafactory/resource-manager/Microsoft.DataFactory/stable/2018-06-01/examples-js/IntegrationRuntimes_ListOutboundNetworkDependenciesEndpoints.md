Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.3.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function integrationRuntimesOutboundNetworkDependenciesEndpoints() {
  const subscriptionId = "7ad7c73b-38b8-4df3-84ee-52ff91092f61";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.listOutboundNetworkDependenciesEndpoints(
    resourceGroupName,
    factoryName,
    integrationRuntimeName
  );
  console.log(result);
}

integrationRuntimesOutboundNetworkDependenciesEndpoints().catch(console.error);
```
