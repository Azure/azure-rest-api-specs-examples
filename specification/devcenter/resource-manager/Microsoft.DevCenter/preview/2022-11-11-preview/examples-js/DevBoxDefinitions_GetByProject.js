const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Dev Box definition configured for a project
 *
 * @summary Gets a Dev Box definition configured for a project
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/DevBoxDefinitions_GetByProject.json
 */
async function devBoxDefinitionsGetByProject() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const projectName = "ContosoProject";
  const devBoxDefinitionName = "WebDevBox";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devBoxDefinitions.getByProject(
    resourceGroupName,
    projectName,
    devBoxDefinitionName
  );
  console.log(result);
}

devBoxDefinitionsGetByProject().catch(console.error);
