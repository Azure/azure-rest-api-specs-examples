const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific project.
 *
 * @summary Gets a specific project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Projects_Get.json
 */
async function projectsGet() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const projectName = "DevProject";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projects.get(resourceGroupName, projectName);
  console.log(result);
}

projectsGet().catch(console.error);
