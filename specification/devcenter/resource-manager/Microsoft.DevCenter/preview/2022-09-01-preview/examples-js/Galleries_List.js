const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists galleries for a devcenter.
 *
 * @summary Lists galleries for a devcenter.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/Galleries_List.json
 */
async function galleriesListByDevCenter() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.galleries.listByDevCenter(resourceGroupName, devCenterName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

galleriesListByDevCenter().catch(console.error);
