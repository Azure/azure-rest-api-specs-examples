const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a BuildResource
 *
 * @summary Create a BuildResource
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Builds_CreateOrUpdate_NoConfig.json
 */
async function buildsCreateOrUpdateNoConfig() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const builderName = "testBuilder";
  const buildName = "testBuild";
  const buildEnvelope = {};
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.builds.beginCreateOrUpdateAndWait(
    resourceGroupName,
    builderName,
    buildName,
    buildEnvelope,
  );
  console.log(result);
}
