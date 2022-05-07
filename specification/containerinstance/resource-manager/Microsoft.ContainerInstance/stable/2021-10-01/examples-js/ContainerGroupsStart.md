Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerinstance_8.1.0/sdk/containerinstance/arm-containerinstance/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts all containers in a container group. Compute resources will be allocated and billing will start.
 *
 * @summary Starts all containers in a container group. Compute resources will be allocated and billing will start.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsStart.json
 */
async function containerStart() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.beginStartAndWait(
    resourceGroupName,
    containerGroupName
  );
  console.log(result);
}

containerStart().catch(console.error);
```
