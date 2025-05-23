const { AzureMigrateAssessmentService } = require("@azure/arm-migrationassessment");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List AssessmentProject resources by subscription ID
 *
 * @summary List AssessmentProject resources by subscription ID
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessmentProjectsOperations_ListBySubscription_MaximumSet_Gen.json
 */
async function assessmentProjectsOperationsListBySubscriptionMaximumSetGen() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateAssessmentService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assessmentProjectsOperations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
