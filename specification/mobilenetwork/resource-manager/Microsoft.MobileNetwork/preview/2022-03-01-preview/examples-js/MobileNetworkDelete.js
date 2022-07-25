const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified mobile network.
 *
 * @summary Deletes the specified mobile network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/MobileNetworkDelete.json
 */
async function deleteMobileNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.mobileNetworks.beginDeleteAndWait(
    resourceGroupName,
    mobileNetworkName
  );
  console.log(result);
}

deleteMobileNetwork().catch(console.error);
