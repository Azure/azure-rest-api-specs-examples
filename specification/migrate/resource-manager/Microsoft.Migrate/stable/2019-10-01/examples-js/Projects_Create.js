const { AzureMigrateV2 } = require("@azure/arm-migrate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a project with specified name. If a project already exists, update it.
 *
 * @summary Create a project with specified name. If a project already exists, update it.
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Projects_Create.json
 */
async function projectsCreate() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "6393a73f-8d55-47ef-b6dd-179b3e0c7910";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "abgoyal-westEurope";
  const projectName = "abGoyalProject2";
  const project = {
    eTag: "",
    location: "West Europe",
    properties: {
      assessmentSolutionId:
        "/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.migrate/migrateprojects/abgoyalweselfhost/Solutions/Servers-Assessment-ServerAssessment",
      customerWorkspaceId: undefined,
      customerWorkspaceLocation: undefined,
      projectStatus: "Active",
    },
    tags: {},
  };
  const options = { project };
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateV2(credential, subscriptionId);
  const result = await client.projects.create(resourceGroupName, projectName, options);
  console.log(result);
}
