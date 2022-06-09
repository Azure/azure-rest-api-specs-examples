```javascript
const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a resource group.
 *
 * @summary Creates or updates a resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/CreateResourceGroup.json
 */
async function createOrUpdateAResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "my-resource-group";
  const parameters = { location: "eastus" };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.resourceGroups.createOrUpdate(resourceGroupName, parameters);
  console.log(result);
}

createOrUpdateAResourceGroup().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources_5.0.1/sdk/resources/arm-resources/README.md) on how to add the SDK to your project and authenticate.
