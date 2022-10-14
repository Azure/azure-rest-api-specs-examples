const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Partially updates a project.
 *
 * @summary Partially updates a project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Projects_Patch.json
 */
async function projectsUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const body = {
    description: "This is my first project.",
    tags: { costCenter: "R&D" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projects.beginUpdateAndWait(resourceGroupName, projectName, body);
  console.log(result);
}

projectsUpdate().catch(console.error);
