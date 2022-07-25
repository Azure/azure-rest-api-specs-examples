const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all dataNetworks in the mobile network.
 *
 * @summary Lists all dataNetworks in the mobile network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/DataNetworkListByMobileNetwork.json
 */
async function listMobileNetworkDataNetworksInAMobileNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataNetworks.listByMobileNetwork(
    resourceGroupName,
    mobileNetworkName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listMobileNetworkDataNetworksInAMobileNetwork().catch(console.error);
