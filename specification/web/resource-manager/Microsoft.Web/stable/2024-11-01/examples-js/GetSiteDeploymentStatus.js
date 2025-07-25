const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the deployment status for an app (or deployment slot, if specified).
 *
 * @summary Gets the deployment status for an app (or deployment slot, if specified).
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetSiteDeploymentStatus.json
 */
async function getDeploymentStatus() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testSite";
  const deploymentStatusId = "eacfd68b-3bbd-4ad9-99c5-98614d89c8e5";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.beginGetProductionSiteDeploymentStatusAndWait(
    resourceGroupName,
    name,
    deploymentStatusId,
  );
  console.log(result);
}
