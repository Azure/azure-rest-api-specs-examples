const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a project.
 *
 * @summary creates or updates a project.
 * x-ms-original-file: 2026-01-01-preview/Projects_PutWithMaxDevBoxPerUser.json
 */
async function projectsCreateOrUpdateWithLimitsPerDev() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58fffff";
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projects.createOrUpdate("rg1", "DevProject", {
    location: "centralus",
    description: "This is my first project.",
    devCenterId:
      "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso",
    maxDevBoxesPerUser: 3,
    tags: { CostCenter: "R&D" },
  });
  console.log(result);
}
