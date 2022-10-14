const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Partially updates a catalog.
 *
 * @summary Partially updates a catalog.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Catalogs_Patch.json
 */
async function catalogsUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const catalogName = "{catalogName}";
  const body = { gitHub: { path: "/environments" } };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.catalogs.beginUpdateAndWait(
    resourceGroupName,
    devCenterName,
    catalogName,
    body
  );
  console.log(result);
}

catalogsUpdate().catch(console.error);
