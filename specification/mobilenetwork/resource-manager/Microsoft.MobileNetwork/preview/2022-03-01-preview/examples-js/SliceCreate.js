const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a mobile network slice.
 *
 * @summary Creates or updates a mobile network slice.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SliceCreate.json
 */
async function createMobileNetworkSlice() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const sliceName = "testSlice";
  const parameters = {
    description: "myFavouriteSlice",
    location: "eastus",
    snssai: { sd: "1abcde", sst: 1 },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.slices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    mobileNetworkName,
    sliceName,
    parameters
  );
  console.log(result);
}

createMobileNetworkSlice().catch(console.error);
