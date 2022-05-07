Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateASimpleGalleryWithSoftDeletionEnabled() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const gallery = {
    description: "This is the gallery description.",
    location: "West US",
    softDeletePolicy: { isSoftDeleteEnabled: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleries.beginCreateOrUpdateAndWait(
    resourceGroupName,
    galleryName,
    gallery
  );
  console.log(result);
}

createOrUpdateASimpleGalleryWithSoftDeletionEnabled().catch(console.error);
```
