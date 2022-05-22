Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicelinker_2.0.0/sdk/servicelinker/arm-servicelinker/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list source configurations for a linker.
 *
 * @summary list source configurations for a linker.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/GetConfigurations.json
 */
async function getConfiguration() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const linkerName = "linkName";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linker.listConfigurations(resourceUri, linkerName);
  console.log(result);
}

getConfiguration().catch(console.error);
```
