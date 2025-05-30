const { AzureMigrateAssessmentService } = require("@azure/arm-migrationassessment");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a AvsAssessmentOptions
 *
 * @summary Get a AvsAssessmentOptions
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AvsAssessmentOptionsOperations_Get_MaximumSet_Gen.json
 */
async function avsAssessmentOptionsOperationsGetMaximumSetGen() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "ayagrawrg";
  const projectName = "app18700project";
  const avsAssessmentOptionsName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateAssessmentService(credential, subscriptionId);
  const result = await client.avsAssessmentOptionsOperations.get(
    resourceGroupName,
    projectName,
    avsAssessmentOptionsName,
  );
  console.log(result);
}
