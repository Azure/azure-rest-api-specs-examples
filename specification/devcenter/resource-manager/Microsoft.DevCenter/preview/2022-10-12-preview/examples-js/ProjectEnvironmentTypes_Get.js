const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a project environment type.
 *
 * @summary Gets a project environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/ProjectEnvironmentTypes_Get.json
 */
async function projectEnvironmentTypesGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "ContosoProj";
  const environmentTypeName = "{environmentTypeName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projectEnvironmentTypes.get(
    resourceGroupName,
    projectName,
    environmentTypeName
  );
  console.log(result);
}

projectEnvironmentTypesGet().catch(console.error);
