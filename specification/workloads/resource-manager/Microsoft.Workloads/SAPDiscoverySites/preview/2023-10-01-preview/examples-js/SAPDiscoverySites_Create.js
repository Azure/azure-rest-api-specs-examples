const { WorkloadsClient } = require("@azure/arm-migrationdiscoverysap");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a discovery site resource for SAP Migration. This resource will be used to run system discovery and assessment with Azure Migrate.
 *
 * @summary Creates a discovery site resource for SAP Migration. This resource will be used to run system discovery and assessment with Azure Migrate.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/SAPDiscoverySites_Create.json
 */
async function createResourceForImportBasedInput() {
  const subscriptionId =
    process.env["MIGRATIONDISCOVERY_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["MIGRATIONDISCOVERY_RESOURCE_GROUP"] || "test-rg";
  const sapDiscoverySiteName = "SampleSite";
  const resource = {
    location: "eastus",
    properties: {
      masterSiteId: "MasterSiteIdResourceId",
      migrateProjectId: "MigrateProjectId",
    },
    tags: { property1: "value1", property2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapDiscoverySites.beginCreateAndWait(
    resourceGroupName,
    sapDiscoverySiteName,
    resource,
  );
  console.log(result);
}
