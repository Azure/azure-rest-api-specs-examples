const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a mobile network.
 *
 * @summary Creates or updates a mobile network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/MobileNetworkCreate.json
 */
async function createMobileNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const parameters = {
    location: "eastus",
    publicLandMobileNetworkIdentifier: { mcc: "001", mnc: "01" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.mobileNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    mobileNetworkName,
    parameters
  );
  console.log(result);
}

createMobileNetwork().catch(console.error);
