const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a gallery.
 *
 * @summary Creates or updates a gallery.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Galleries_Create.json
 */
async function galleriesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const galleryName = "{galleryName}";
  const body = {
    galleryResourceId:
      "/subscriptions/{subscriptionId}/resourceGroups/rg1/providers/Microsoft.Compute/galleries/{galleryName}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.galleries.beginCreateOrUpdateAndWait(
    resourceGroupName,
    devCenterName,
    galleryName,
    body
  );
  console.log(result);
}

galleriesCreateOrUpdate().catch(console.error);
