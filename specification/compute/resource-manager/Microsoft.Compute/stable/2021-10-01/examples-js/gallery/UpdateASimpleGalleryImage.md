Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a gallery image definition.
 *
 * @summary Update a gallery image definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/UpdateASimpleGalleryImage.json
 */
async function updateASimpleGalleryImage() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImage = {
    hyperVGeneration: "V1",
    identifier: {
      offer: "myOfferName",
      publisher: "myPublisherName",
      sku: "mySkuName",
    },
    osState: "Generalized",
    osType: "Windows",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImages.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryImageName,
    galleryImage
  );
  console.log(result);
}

updateASimpleGalleryImage().catch(console.error);
```
