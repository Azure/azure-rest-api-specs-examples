const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about a gallery image version.
 *
 * @summary Retrieves information about a gallery image version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryImageVersionWithSnapshotsAsSource.json
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
