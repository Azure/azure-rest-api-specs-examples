const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all sites in the mobile network.
 *
 * @summary Lists all sites in the mobile network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SiteListByMobileNetwork.json
 */
async function listMobileNetworkSitesInAMobileNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sites.listByMobileNetwork(resourceGroupName, mobileNetworkName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listMobileNetworkSitesInAMobileNetwork().catch(console.error);
