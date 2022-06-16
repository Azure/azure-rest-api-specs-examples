const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List community gallery images inside a gallery.
 *
 * @summary List community gallery images inside a gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-01-03/GalleryRP/examples/communityGalleryExamples/CommunityGalleryImage_List.json
 */
async function listCommunityGalleryImages() {
  const subscriptionId = "{subscription-id}";
  const location = "myLocation";
  const publicGalleryName = "publicGalleryName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communityGalleryImages.list(location, publicGalleryName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCommunityGalleryImages().catch(console.error);
