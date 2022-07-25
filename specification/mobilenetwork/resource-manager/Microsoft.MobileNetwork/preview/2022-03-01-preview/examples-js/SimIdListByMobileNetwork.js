const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List sim ids under a mobile network.
 *
 * @summary List sim ids under a mobile network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimIdListByMobileNetwork.json
 */
async function listSimProfileIdsByNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const mobileNetworkName = "testMobileNetworkName";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.mobileNetworks.beginListSimIdsAndWait(
    resourceGroupName,
    mobileNetworkName
  );
  console.log(result);
}

listSimProfileIdsByNetwork().catch(console.error);
