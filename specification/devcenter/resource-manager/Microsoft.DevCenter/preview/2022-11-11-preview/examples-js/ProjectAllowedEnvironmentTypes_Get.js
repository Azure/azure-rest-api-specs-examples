const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an allowed environment type.
 *
 * @summary Gets an allowed environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectAllowedEnvironmentTypes_Get.json
 */
async function projectAllowedEnvironmentTypesGet() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const projectName = "Contoso";
  const environmentTypeName = "DevTest";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projectAllowedEnvironmentTypes.get(
    resourceGroupName,
    projectName,
    environmentTypeName
  );
  console.log(result);
}

projectAllowedEnvironmentTypesGet().catch(console.error);
