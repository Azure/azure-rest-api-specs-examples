const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Create or update a worker pool.
 *
 * @summary Description for Create or update a worker pool.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/AppServiceEnvironments_CreateOrUpdateWorkerPool.json
 */
async function getPropertiesOfAWorkerPool() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "test-rg";
  const name = "test-ase";
  const workerPoolName = "0";
  const workerPoolEnvelope = {
    workerCount: 3,
    workerSize: "Small",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceEnvironments.beginCreateOrUpdateWorkerPoolAndWait(
    resourceGroupName,
    name,
    workerPoolName,
    workerPoolEnvelope,
  );
  console.log(result);
}
