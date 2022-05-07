Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineImagesEdgeZoneListMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaa";
  const edgeZone = "aaaaaaaaaaaaaaaaaaaaaaaaa";
  const publisherName = "aaaa";
  const offer = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const skus = "aaaaaaaaaaaaaaaaaaaaaaa";
  const expand = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const top = 12;
  const orderby = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const options = { expand: expand, top: top, orderby: orderby };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.list(
    location,
    edgeZone,
    publisherName,
    offer,
    skus,
    options
  );
  console.log(result);
}

virtualMachineImagesEdgeZoneListMaximumSetGen().catch(console.error);
```
