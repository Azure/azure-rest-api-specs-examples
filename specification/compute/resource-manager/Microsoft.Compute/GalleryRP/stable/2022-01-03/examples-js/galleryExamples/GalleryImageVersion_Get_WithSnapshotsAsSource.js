const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about a gallery image version.
 *
 * @summary Retrieves information about a gallery image version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryImageVersion_Get_WithSnapshotsAsSource.json
 */
async function getAGalleryImageVersionWithSnapshotsAsASource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImageVersionName = "1.0.0";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryImageVersions.get(
    resourceGroupName,
    galleryName,
    galleryImageName,
    galleryImageVersionName
  );
  console.log(result);
}

getAGalleryImageVersionWithSnapshotsAsASource().catch(console.error);
