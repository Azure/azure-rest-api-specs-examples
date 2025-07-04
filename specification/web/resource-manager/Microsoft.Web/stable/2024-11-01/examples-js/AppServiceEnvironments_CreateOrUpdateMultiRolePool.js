const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Create or update a multi-role pool.
 *
 * @summary Description for Create or update a multi-role pool.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/AppServiceEnvironments_CreateOrUpdateMultiRolePool.json
 */
async function createOrUpdateAMultiRolePool() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "test-rg";
  const name = "test-ase";
  const multiRolePoolEnvelope = {
    workerCount: 3,
    workerSize: "Medium",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceEnvironments.beginCreateOrUpdateMultiRolePoolAndWait(
    resourceGroupName,
    name,
    multiRolePoolEnvelope,
  );
  console.log(result);
}
