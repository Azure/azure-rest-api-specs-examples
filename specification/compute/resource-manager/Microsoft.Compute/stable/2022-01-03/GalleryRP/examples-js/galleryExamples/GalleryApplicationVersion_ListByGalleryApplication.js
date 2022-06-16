const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List gallery Application Versions in a gallery Application Definition.
 *
 * @summary List gallery Application Versions in a gallery Application Definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/galleryExamples/GalleryApplicationVersion_ListByGalleryApplication.json
 */
async function listGalleryApplicationVersionsInAGalleryApplicationDefinition() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryApplicationName = "myGalleryApplicationName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleryApplicationVersions.listByGalleryApplication(
    resourceGroupName,
    galleryName,
    galleryApplicationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listGalleryApplicationVersionsInAGalleryApplicationDefinition().catch(console.error);
