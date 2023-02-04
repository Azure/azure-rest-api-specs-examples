const { AzureMigrateV2 } = require("@azure/arm-migrate");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of Import collector.
 *
 * @summary Get a list of Import collector.
 * x-ms-original-file: specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ImportCollectors_ListByProject.json
 */
async function importCollectorsListByProject() {
  const subscriptionId =
    process.env["MIGRATE_SUBSCRIPTION_ID"] || "31be0ff4-c932-4cb3-8efc-efa411d79280";
  const resourceGroupName = process.env["MIGRATE_RESOURCE_GROUP"] || "markusavstestrg";
  const projectName = "rajoshCCY9671project";
  const credential = new DefaultAzureCredential();
  const client = new AzureMigrateV2(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.importCollectors.listByProject(resourceGroupName, projectName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
