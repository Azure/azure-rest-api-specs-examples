const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists catalogs for a devcenter.
 *
 * @summary Lists catalogs for a devcenter.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Catalogs_List.json
 */
async function catalogsListByDevCenter() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.catalogs.listByDevCenter(resourceGroupName, devCenterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

catalogsListByDevCenter().catch(console.error);
