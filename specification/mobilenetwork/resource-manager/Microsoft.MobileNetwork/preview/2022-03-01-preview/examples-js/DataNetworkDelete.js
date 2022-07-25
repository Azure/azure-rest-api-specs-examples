const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified mobile network dataNetwork.
 *
 * @summary Deletes the specified mobile network dataNetwork.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/DataNetworkDelete.json
 */
async function deleteMobileNetworkDataNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const dataNetworkName = "testDataNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.dataNetworks.beginDeleteAndWait(
    resourceGroupName,
    mobileNetworkName,
    dataNetworkName
  );
  console.log(result);
}

deleteMobileNetworkDataNetwork().catch(console.error);
