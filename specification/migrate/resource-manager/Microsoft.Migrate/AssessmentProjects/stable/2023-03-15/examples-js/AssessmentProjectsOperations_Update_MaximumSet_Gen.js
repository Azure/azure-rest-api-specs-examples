const { AzureMigrateAssessmentService } = require("@azure/arm-migrationassessment");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update a AssessmentProject
 *
 * @summary Update a AssessmentProject
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentProjectsOperations_Update_MaximumSet_Gen.json
 */
async function assessmentProjectsOperationsUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "sakanwar";
  const projectName = "sakanwar1204project";
  const properties = {
    assessmentSolutionId:
      "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa",
    customerStorageAccountArmId:
      "/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sakanwar/providers/Microsoft.Storage/storageAccounts/sakanwar1204usa",
    customerWorkspaceId: undefined,
    customerWorkspaceLocation: undefined,
    projectStatus: "Active",
    provisioningState: "Succeeded",
    publicNetworkAccess: "Disabled",
    tags: { migrateProject: "sakanwar-PE-SEA" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateAssessmentService(credential, subscriptionId);
  const result = await client.assessmentProjectsOperations.beginUpdateAndWait(
    resourceGroupName,
    projectName,
    properties,
  );
  console.log(result);
}
