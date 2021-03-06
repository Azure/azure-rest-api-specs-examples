const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a community gallery image.
 *
 * @summary Get a community gallery image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/communityGallery/GetACommunityGalleryImage.json
 */
async function getAGallery() {
  const subscriptionId = "{subscription-id}";
  const location = "myLocation";
  const publicGalleryName = "publicGalleryName";
  const galleryImageName = "myGalleryImageName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.communityGalleryImages.get(
    location,
    publicGalleryName,
    galleryImageName
  );
  console.log(result);
}

getAGallery().catch(console.error);
