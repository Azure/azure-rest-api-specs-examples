const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists versions for an image.
 *
 * @summary Lists versions for an image.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ImageVersions_List.json
 */
async function imageVersionsListByImage() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const galleryName = "DefaultDevGallery";
  const imageName = "Win11";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.imageVersions.listByImage(
    resourceGroupName,
    devCenterName,
    galleryName,
    imageName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

imageVersionsListByImage().catch(console.error);
