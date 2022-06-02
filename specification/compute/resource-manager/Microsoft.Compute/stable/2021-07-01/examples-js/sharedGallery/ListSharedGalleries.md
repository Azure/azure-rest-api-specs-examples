Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List shared galleries by subscription id or tenant id.
 *
 * @summary List shared galleries by subscription id or tenant id.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/sharedGallery/ListSharedGalleries.json
 */
async function getAGallery() {
  const subscriptionId = "{subscription-id}";
  const location = "myLocation";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sharedGalleries.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAGallery().catch(console.error);
```
