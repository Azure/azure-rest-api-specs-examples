const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update sharing profile of a gallery.
 *
 * @summary Update sharing profile of a gallery.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/Gallery_AddToSharingProfile.json
 */
async function addSharingIdToTheSharingProfileOfAGallery() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const sharingUpdate = {
    groups: [
      {
        type: "Subscriptions",
        ids: ["34a4ab42-0d72-47d9-bd1a-aed207386dac", "380fd389-260b-41aa-bad9-0a83108c370b"],
      },
      { type: "AADTenants", ids: ["c24c76aa-8897-4027-9b03-8f7928b54ff6"] },
    ],
    operationType: "Add",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.gallerySharingProfile.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    sharingUpdate
  );
  console.log(result);
}

addSharingIdToTheSharingProfileOfAGallery().catch(console.error);
