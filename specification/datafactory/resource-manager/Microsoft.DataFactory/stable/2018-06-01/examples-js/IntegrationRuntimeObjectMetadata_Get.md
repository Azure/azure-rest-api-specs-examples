Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.3.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function integrationRuntimeObjectMetadataGet() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "testactivityv2";
  const getMetadataRequest = {
    metadataPath: "ssisFolders",
  };
  const options = { getMetadataRequest: getMetadataRequest };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimeObjectMetadata.get(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    options
  );
  console.log(result);
}

integrationRuntimeObjectMetadataGet().catch(console.error);
```
