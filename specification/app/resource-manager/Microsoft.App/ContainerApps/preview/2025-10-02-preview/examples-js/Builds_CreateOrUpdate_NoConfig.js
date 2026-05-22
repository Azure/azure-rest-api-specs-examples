const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a BuildResource
 *
 * @summary create a BuildResource
 * x-ms-original-file: 2025-10-02-preview/Builds_CreateOrUpdate_NoConfig.json
 */
async function buildsCreateOrUpdateNoConfig() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.builds.createOrUpdate("rg", "testBuilder", "testBuild", {});
  console.log(result);
}
