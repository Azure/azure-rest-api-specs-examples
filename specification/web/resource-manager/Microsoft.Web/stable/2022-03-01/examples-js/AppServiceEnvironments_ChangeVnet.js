const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Move an App Service Environment to a different VNET.
 *
 * @summary Description for Move an App Service Environment to a different VNET.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_ChangeVnet.json
 */
async function moveAnAppServiceEnvironmentToADifferentVnet() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-rg";
  const name = "test-ase";
  const vnetInfo = {
    id: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.appServiceEnvironments.beginListChangeVnetAndWait(
    resourceGroupName,
    name,
    vnetInfo
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

moveAnAppServiceEnvironmentToADifferentVnet().catch(console.error);
