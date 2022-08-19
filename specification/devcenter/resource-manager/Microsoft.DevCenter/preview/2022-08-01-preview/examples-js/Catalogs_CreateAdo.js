const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a catalog.
 *
 * @summary Creates or updates a catalog.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Catalogs_CreateAdo.json
 */
async function catalogsCreateOrUpdateAdo() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const catalogName = "{catalogName}";
  const body = {
    adoGit: {
      path: "/templates",
      branch: "main",
      secretIdentifier: "https://contosokv.vault.azure.net/secrets/CentralRepoPat",
      uri: "https://contoso@dev.azure.com/contoso/contosoOrg/_git/centralrepo-fakecontoso",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.catalogs.beginCreateOrUpdateAndWait(
    resourceGroupName,
    devCenterName,
    catalogName,
    body
  );
  console.log(result);
}

catalogsCreateOrUpdateAdo().catch(console.error);
