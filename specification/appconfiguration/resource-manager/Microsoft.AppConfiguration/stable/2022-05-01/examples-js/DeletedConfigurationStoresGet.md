```javascript
const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a deleted Azure app configuration store.
 *
 * @summary Gets a deleted Azure app configuration store.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/DeletedConfigurationStoresGet.json
 */
async function deletedConfigurationStoresGet() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const location = "westus";
  const configStoreName = "contoso";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.getDeleted(location, configStoreName);
  console.log(result);
}

deletedConfigurationStoresGet().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appconfiguration_3.0.0/sdk/appconfiguration/arm-appconfiguration/README.md) on how to add the SDK to your project and authenticate.
