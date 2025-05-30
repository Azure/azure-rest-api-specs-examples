const { AzureMigrateAssessmentService } = require("@azure/arm-migrationassessment");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List AssessedMachine resources by Assessment
 *
 * @summary List AssessedMachine resources by Assessment
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessedMachinesOperations_ListByAssessment_MaximumSet_Gen.json
 */
async function assessedMachinesOperationsListByAssessmentMaximumSetGen() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "D8E1C413-E65F-40C0-8A7E-743D6B7A6AE9";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "rgopenapi";
  const filter = "sbkdovsfqldhdb";
  const pageSize = 10;
  const continuationToken = "hbyseetshbplfkjmpjhsiurqgt";
  const totalRecordCount = 25;
  const projectName = "sloqixzfjk";
  const groupName = "kjuepxerwseq";
  const assessmentName = "rhzcmubwrrkhtocsibu";
  const options = {
    filter,
    pageSize,
    continuationToken,
    totalRecordCount,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateAssessmentService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assessedMachinesOperations.listByAssessment(
    resourceGroupName,
    projectName,
    groupName,
    assessmentName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
