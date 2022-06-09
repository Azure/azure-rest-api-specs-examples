```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove all linked integration runtimes under specific data factory in a self-hosted integration runtime.
 *
 * @summary Remove all linked integration runtimes under specific data factory in a self-hosted integration runtime.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_RemoveLinks.json
 */
async function integrationRuntimesUpgrade() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const linkedIntegrationRuntimeRequest = {
    linkedFactoryName: "exampleFactoryName-linked",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.removeLinks(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    linkedIntegrationRuntimeRequest
  );
  console.log(result);
}

integrationRuntimesUpgrade().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.6.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.
