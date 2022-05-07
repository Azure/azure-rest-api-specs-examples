Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerinstance_8.1.0/sdk/containerinstance/arm-containerinstance/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified container group in the specified subscription and resource group. The operation returns the properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
 *
 * @summary Gets the properties of the specified container group in the specified subscription and resource group. The operation returns the properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsGet_Failed.json
 */
async function containerGroupsGetFailed() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.get(resourceGroupName, containerGroupName);
  console.log(result);
}

containerGroupsGetFailed().catch(console.error);
```
