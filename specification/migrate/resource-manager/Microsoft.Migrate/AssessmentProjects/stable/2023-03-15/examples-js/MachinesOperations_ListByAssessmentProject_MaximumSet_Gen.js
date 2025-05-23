const { AzureMigrateAssessmentService } = require("@azure/arm-migrationassessment");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List Machine resources by AssessmentProject
 *
 * @summary List Machine resources by AssessmentProject
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/MachinesOperations_ListByAssessmentProject_MaximumSet_Gen.json
 */
async function machinesOperationsListByAssessmentProjectMaximumSetGen() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "ayagrawrg";
  const filter = undefined;
  const pageSize = 1;
  const continuationToken = undefined;
  const totalRecordCount = 1;
  const projectName = "app18700project";
  const options = {
    filter,
    pageSize,
    continuationToken,
    totalRecordCount,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateAssessmentService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.machinesOperations.listByAssessmentProject(
    resourceGroupName,
    projectName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
