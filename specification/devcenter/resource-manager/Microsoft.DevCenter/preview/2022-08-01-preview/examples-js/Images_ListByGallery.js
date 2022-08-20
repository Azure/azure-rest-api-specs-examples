const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists images for a gallery.
 *
 * @summary Lists images for a gallery.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Images_ListByGallery.json
 */
async function imagesListByGallery() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const galleryName = "DevGallery";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.images.listByGallery(
    resourceGroupName,
    devCenterName,
    galleryName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

imagesListByGallery().catch(console.error);
