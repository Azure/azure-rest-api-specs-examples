const { AzureMigrateAssessmentService } = require("@azure/arm-migrationassessment");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List AssessedSqlDatabaseV2 resources by SqlAssessmentV2
 *
 * @summary List AssessedSqlDatabaseV2 resources by SqlAssessmentV2
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/stable/2023-03-15/examples/AssessedSqlDatabaseV2Operations_ListBySqlAssessmentV2_MaximumSet_Gen.json
 */
async function assessedSqlDatabaseV2OperationsListBySqlAssessmentV2MaximumSetGen() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "4bd2aa0f-2bd2-4d67-91a8-5a4533d58600";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "rgmigrate";
  const filter = "(contains(Properties/DatabaseName,'adv130'))";
  const pageSize = 23;
  const continuationToken = undefined;
  const totalRecordCount = 1;
  const projectName = "fci-test6904project";
  const groupName = "test_fci_hadr";
  const assessmentName = "test_swagger_1";
  const options = { filter, pageSize, continuationToken, totalRecordCount };
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateAssessmentService(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assessedSqlDatabaseV2Operations.listBySqlAssessmentV2(
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
