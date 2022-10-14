const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists allowed environment types for a project.
 *
 * @summary Lists allowed environment types for a project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/ProjectAllowedEnvironmentTypes_List.json
 */
async function projectAllowedEnvironmentTypesList() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "Contoso";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.projectAllowedEnvironmentTypes.list(
    resourceGroupName,
    projectName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

projectAllowedEnvironmentTypesList().catch(console.error);
