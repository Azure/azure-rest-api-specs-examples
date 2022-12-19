const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Dev Box definitions configured for a project.
 *
 * @summary List Dev Box definitions configured for a project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/DevBoxDefinitions_ListByProject.json
 */
async function devBoxDefinitionsListByProject() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const projectName = "ContosoProject";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.devBoxDefinitions.listByProject(resourceGroupName, projectName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

devBoxDefinitionsListByProject().catch(console.error);
