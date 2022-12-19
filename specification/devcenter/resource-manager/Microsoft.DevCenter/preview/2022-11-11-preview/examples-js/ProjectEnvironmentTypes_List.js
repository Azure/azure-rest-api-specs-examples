const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists environment types for a project.
 *
 * @summary Lists environment types for a project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_List.json
 */
async function projectEnvironmentTypesList() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const projectName = "ContosoProj";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.projectEnvironmentTypes.list(resourceGroupName, projectName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

projectEnvironmentTypesList().catch(console.error);
