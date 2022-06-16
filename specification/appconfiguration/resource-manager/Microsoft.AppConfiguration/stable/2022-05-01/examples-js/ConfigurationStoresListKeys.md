```javascript
const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the access key for the specified configuration store.
 *
 * @summary Lists the access key for the specified configuration store.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresListKeys.json
 */
async function configurationStoresListKeys() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationStores.listKeys(resourceGroupName, configStoreName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

configurationStoresListKeys().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appconfiguration_3.0.0/sdk/appconfiguration/arm-appconfiguration/README.md) on how to add the SDK to your project and authenticate.
