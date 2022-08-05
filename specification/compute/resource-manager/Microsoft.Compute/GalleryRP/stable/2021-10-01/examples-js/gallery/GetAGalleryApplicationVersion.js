const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about a gallery Application Version.
 *
 * @summary Retrieves information about a gallery Application Version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/GetAGalleryApplicationVersion.json
 */
async function getAGalleryApplicationVersion() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryApplicationName = "myGalleryApplicationName";
  const galleryApplicationVersionName = "1.0.0";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryApplicationVersions.get(
    resourceGroupName,
    galleryName,
    galleryApplicationName,
    galleryApplicationVersionName
  );
  console.log(result);
}

getAGalleryApplicationVersion().catch(console.error);
