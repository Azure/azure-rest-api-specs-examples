const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a community gallery image version.
 *
 * @summary Get a community gallery image version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/communityGalleryExamples/CommunityGalleryImageVersion_Get.json
 */
async function getACommunityGalleryImageVersion() {
  const subscriptionId = "{subscription-id}";
  const location = "myLocation";
  const publicGalleryName = "publicGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImageVersionName = "myGalleryImageVersionName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.communityGalleryImageVersions.get(
    location,
    publicGalleryName,
    galleryImageName,
    galleryImageVersionName
  );
  console.log(result);
}

getACommunityGalleryImageVersion().catch(console.error);
