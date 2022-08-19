const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a gallery image.
 *
 * @summary Gets a gallery image.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Images_Get.json
 */
async function imagesGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const galleryName = "DefaultDevGallery";
  const imageName = "{imageName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.images.get(resourceGroupName, devCenterName, galleryName, imageName);
  console.log(result);
}

imagesGet().catch(console.error);
