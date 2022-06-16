const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List gallery images in a given lab.
 *
 * @summary List gallery images in a given lab.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GalleryImages_List.json
 */
async function galleryImagesList() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleryImages.list(resourceGroupName, labName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

galleryImagesList().catch(console.error);
