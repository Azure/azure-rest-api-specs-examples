const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Dev Box definition configured for a project
 *
 * @summary Gets a Dev Box definition configured for a project
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/DevBoxDefinitions_GetByProject.json
 */
async function devBoxDefinitionsGetByProject() {
  const subscriptionId = "{subscriptionId}";
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
