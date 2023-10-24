const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Dev Box definitions in the catalog.
 *
 * @summary List Dev Box definitions in the catalog.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_ListByCatalog.json
 */
async function catalogDevBoxDefinitionsListByCatalog() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const devCenterName = "Contoso";
  const catalogName = "CentralCatalog";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.catalogDevBoxDefinitions.listByCatalog(
    resourceGroupName,
    devCenterName,
    catalogName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
