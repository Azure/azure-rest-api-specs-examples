```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a SSIS integration runtime object metadata by specified path. The return is pageable metadata list.
 *
 * @summary Get a SSIS integration runtime object metadata by specified path. The return is pageable metadata list.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimeObjectMetadata_Get.json
 */
async function integrationRuntimeObjectMetadataGet() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "testactivityv2";
  const getMetadataRequest = {
    metadataPath: "ssisFolders",
  };
  const options = {
    getMetadataRequest,
  };
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.6.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.
