const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details for specific Container Apps Patch by patch name.
 *
 * @summary Get details for specific Container Apps Patch by patch name.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsPatches_Get.json
 */
async function containerAppsPatchesGet0() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "test-app";
  const patchName = "testPatch-25fe4b";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsPatches.get(
    resourceGroupName,
    containerAppName,
    patchName,
  );
  console.log(result);
}
