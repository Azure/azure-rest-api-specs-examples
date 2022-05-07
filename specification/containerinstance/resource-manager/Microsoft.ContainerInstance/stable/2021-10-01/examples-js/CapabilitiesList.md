Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerinstance_8.1.0/sdk/containerinstance/arm-containerinstance/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the list of CPU/memory/GPU capabilities of a region.
 *
 * @summary Get the list of CPU/memory/GPU capabilities of a region.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/CapabilitiesList.json
 */
async function getCapabilities() {
  const subscriptionId = "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listCapabilities(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getCapabilities().catch(console.error);
```
