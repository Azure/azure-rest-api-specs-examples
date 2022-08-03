const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified network slice.
 *
 * @summary Deletes the specified network slice.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SliceDelete.json
 */
async function deleteNetworkSlice() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const sliceName = "testSlice";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.slices.beginDeleteAndWait(
    resourceGroupName,
    mobileNetworkName,
    sliceName
  );
  console.log(result);
}

deleteNetworkSlice().catch(console.error);
