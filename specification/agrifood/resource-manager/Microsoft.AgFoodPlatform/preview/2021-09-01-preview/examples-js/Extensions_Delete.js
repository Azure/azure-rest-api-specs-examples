const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Uninstall extension.
 *
 * @summary Uninstall extension.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Extensions_Delete.json
 */
async function extensionsDelete() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "examples-rg";
  const farmBeatsResourceName = "examples-farmbeatsResourceName";
  const extensionId = "provider.extension";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const result = await client.extensions.delete(
    resourceGroupName,
    farmBeatsResourceName,
    extensionId
  );
  console.log(result);
}

extensionsDelete().catch(console.error);
