const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a gallery resource.
 *
 * @summary Deletes a gallery resource.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Galleries_Delete.json
 */
async function galleriesDelete() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const galleryName = "{galleryName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.galleries.beginDeleteAndWait(
    resourceGroupName,
    devCenterName,
    galleryName
  );
  console.log(result);
}

galleriesDelete().catch(console.error);
