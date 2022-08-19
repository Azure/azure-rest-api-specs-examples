const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a project environment type.
 *
 * @summary Deletes a project environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/ProjectEnvironmentTypes_Delete.json
 */
async function projectEnvironmentTypesDelete() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "ContosoProj";
  const environmentTypeName = "{environmentTypeName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projectEnvironmentTypes.delete(
    resourceGroupName,
    projectName,
    environmentTypeName
  );
  console.log(result);
}

projectEnvironmentTypesDelete().catch(console.error);
