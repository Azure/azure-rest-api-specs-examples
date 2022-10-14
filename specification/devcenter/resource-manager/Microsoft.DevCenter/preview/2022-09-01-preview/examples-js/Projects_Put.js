const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a project.
 *
 * @summary Creates or updates a project.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Projects_Put.json
 */
async function projectsCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const body = {
    description: "This is my first project.",
    devCenterId:
      "/subscriptions/{subscriptionId}/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/{devCenterName}",
    location: "centralus",
    tags: { costCenter: "R&D" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projects.beginCreateOrUpdateAndWait(
    resourceGroupName,
    projectName,
    body
  );
  console.log(result);
}

projectsCreateOrUpdate().catch(console.error);
