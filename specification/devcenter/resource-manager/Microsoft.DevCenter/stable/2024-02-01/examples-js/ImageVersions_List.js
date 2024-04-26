const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists versions for an image.
 *
 * @summary Lists versions for an image.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ImageVersions_List.json
 */
async function imageVersionsListByImage() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
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
    imageName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
