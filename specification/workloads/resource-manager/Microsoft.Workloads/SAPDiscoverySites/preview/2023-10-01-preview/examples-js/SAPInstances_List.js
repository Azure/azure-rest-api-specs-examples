const { WorkloadsClient } = require("@azure/arm-migrationdiscoverysap");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the SAP Instance resources for the given SAP Migration discovery site resource.
 *
 * @summary Lists the SAP Instance resources for the given SAP Migration discovery site resource.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPInstances_List.json
 */
async function listsTheSapInstanceResourcesForTheGivenSapMigrationDiscoverySiteResource() {
  const subscriptionId =
    process.env["MIGRATIONDISCOVERY_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["MIGRATIONDISCOVERY_RESOURCE_GROUP"] || "test-rg";
  const sapDiscoverySiteName = "SampleSite";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sapInstances.listBySapDiscoverySite(
    resourceGroupName,
    sapDiscoverySiteName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
