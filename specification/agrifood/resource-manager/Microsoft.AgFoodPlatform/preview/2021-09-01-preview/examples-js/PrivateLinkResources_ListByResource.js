const { AgriFoodMgmtClient } = require("@azure/arm-agrifood");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get list of Private link resources.
 *
 * @summary Get list of Private link resources.
 * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateLinkResources_ListByResource.json
 */
async function privateLinkResourcesListByResource() {
  const subscriptionId = "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = "examples-rg";
  const farmBeatsResourceName = "examples-farmbeatsResourceName";
  const credential = new DefaultAzureCredential();
  const client = new AgriFoodMgmtClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listByResource(
    resourceGroupName,
    farmBeatsResourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateLinkResourcesListByResource().catch(console.error);
