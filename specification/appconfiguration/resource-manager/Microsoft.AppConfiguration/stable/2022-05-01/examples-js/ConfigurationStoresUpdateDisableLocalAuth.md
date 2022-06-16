```javascript
const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a configuration store with the specified parameters.
 *
 * @summary Updates a configuration store with the specified parameters.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresUpdateDisableLocalAuth.json
 */
async function configurationStoresUpdateDisableLocalAuth() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const configStoreUpdateParameters = {
    disableLocalAuth: true,
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.beginUpdateAndWait(
    resourceGroupName,
    configStoreName,
    configStoreUpdateParameters
  );
  console.log(result);
}

configurationStoresUpdateDisableLocalAuth().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appconfiguration_3.0.0/sdk/appconfiguration/arm-appconfiguration/README.md) on how to add the SDK to your project and authenticate.
