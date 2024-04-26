const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a project catalog.
 *
 * @summary Creates or updates a project catalog.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogs_CreateGitHub.json
 */
async function projectCatalogsCreateOrUpdateGitHub() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const projectName = "DevProject";
  const catalogName = "CentralCatalog";
  const body = {
    gitHub: {
      path: "/templates",
      branch: "main",
      secretIdentifier: "https://contosokv.vault.azure.net/secrets/CentralRepoPat",
      uri: "https://github.com/Contoso/centralrepo-fake.git",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.projectCatalogs.beginCreateOrUpdateAndWait(
    resourceGroupName,
    projectName,
    catalogName,
    body,
  );
  console.log(result);
}
