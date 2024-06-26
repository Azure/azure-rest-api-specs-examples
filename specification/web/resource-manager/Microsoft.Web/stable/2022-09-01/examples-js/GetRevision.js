const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a revision of a Container App.
 *
 * @summary Get a revision of a Container App.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/GetRevision.json
 */
async function getContainerAppRevision() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testcontainerApp0";
  const name = "testcontainerApp0-pjxhsye";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.containerAppsRevisions.getRevision(
    resourceGroupName,
    containerAppName,
    name
  );
  console.log(result);
}
