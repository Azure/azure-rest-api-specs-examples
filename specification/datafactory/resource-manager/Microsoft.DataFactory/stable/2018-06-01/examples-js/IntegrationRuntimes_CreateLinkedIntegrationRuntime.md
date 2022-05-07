Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.3.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function integrationRuntimesCreateLinkedIntegrationRuntime() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const createLinkedIntegrationRuntimeRequest = {
    name: "bfa92911-9fb6-4fbe-8f23-beae87bc1c83",
    dataFactoryLocation: "West US",
    dataFactoryName: "e9955d6d-56ea-4be3-841c-52a12c1a9981",
    subscriptionId: "061774c7-4b5a-4159-a55b-365581830283",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.createLinkedIntegrationRuntime(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    createLinkedIntegrationRuntimeRequest
  );
  console.log(result);
}

integrationRuntimesCreateLinkedIntegrationRuntime().catch(console.error);
```
