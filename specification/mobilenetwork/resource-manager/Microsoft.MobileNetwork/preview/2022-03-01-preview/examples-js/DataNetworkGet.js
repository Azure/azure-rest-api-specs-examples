const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified mobile network dataNetwork.
 *
 * @summary Gets information about the specified mobile network dataNetwork.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/DataNetworkGet.json
 */
async function getMobileNetworkDataNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const dataNetworkName = "testDataNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.dataNetworks.get(
    resourceGroupName,
    mobileNetworkName,
    dataNetworkName
  );
  console.log(result);
}

getMobileNetworkDataNetwork().catch(console.error);
