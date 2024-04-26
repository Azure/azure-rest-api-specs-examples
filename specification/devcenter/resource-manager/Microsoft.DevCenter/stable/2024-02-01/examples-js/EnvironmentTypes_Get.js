const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an environment type.
 *
 * @summary Gets an environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/EnvironmentTypes_Get.json
 */
async function environmentTypesGet() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const devCenterName = "Contoso";
  const environmentTypeName = "DevTest";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.environmentTypes.get(
    resourceGroupName,
    devCenterName,
    environmentTypeName,
  );
  console.log(result);
}
