Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerinstance_8.1.0/sdk/containerinstance/arm-containerinstance/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Attach to the output stream of a specific container instance in a specified resource group and container group.
 *
 * @summary Attach to the output stream of a specific container instance in a specified resource group and container group.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerAttach.json
 */
async function containerAttach() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const containerName = "container1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containers.attach(
    resourceGroupName,
    containerGroupName,
    containerName
  );
  console.log(result);
}

containerAttach().catch(console.error);
```
