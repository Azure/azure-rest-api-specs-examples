const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an image version.
 *
 * @summary Gets an image version.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/ImageVersions_Get.json
 */
async function versionsGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const galleryName = "DefaultDevGallery";
  const imageName = "Win11";
  const versionName = "{versionName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.imageVersions.get(
    resourceGroupName,
    devCenterName,
    galleryName,
    imageName,
    versionName
  );
  console.log(result);
}

versionsGet().catch(console.error);
