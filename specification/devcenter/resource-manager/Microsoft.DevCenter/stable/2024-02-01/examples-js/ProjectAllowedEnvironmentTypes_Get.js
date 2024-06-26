const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an allowed environment type.
 *
 * @summary Gets an allowed environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectAllowedEnvironmentTypes_Get.json
 */
async function projectAllowedEnvironmentTypesGet() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "Contoso";
  const environmentTypeName = "DevTest";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projectAllowedEnvironmentTypes.get(
    resourceGroupName,
    projectName,
    environmentTypeName,
  );
  console.log(result);
}
