Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerinstance_8.1.0/sdk/containerinstance/arm-containerinstance/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of container groups in a specified subscription and resource group. This operation returns properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
 *
 * @summary Get a list of container groups in a specified subscription and resource group. This operation returns properties of each container group including containers, image registry credentials, restart policy, IP address type, OS type, state, and volumes.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsListByResourceGroup.json
 */
async function containerGroupsListByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerGroups.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

containerGroupsListByResourceGroup().catch(console.error);
```
