const { WorkloadsClient } = require("@azure/arm-migrationdiscoverysap");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Server Instance resource. This operation on a resource by end user will return a Bad Request error.
 *
 * @summary Updates the Server Instance resource. This operation on a resource by end user will return a Bad Request error.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/ServerInstances_Update.json
 */
async function updatesTheServerInstanceResource() {
  const subscriptionId =
    process.env["MIGRATIONDISCOVERY_SUBSCRIPTION_ID"] || "6d875e77-e412-4d7d-9af4-8895278b4443";
  const resourceGroupName = process.env["MIGRATIONDISCOVERY_RESOURCE_GROUP"] || "test-rg";
  const sapDiscoverySiteName = "SampleSite";
  const sapInstanceName = "MPP_MPP";
  const serverInstanceName = "APP_SapServer1";
  const properties = { properties: {} };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.serverInstances.update(
    resourceGroupName,
    sapDiscoverySiteName,
    sapInstanceName,
    serverInstanceName,
    properties,
  );
  console.log(result);
}
