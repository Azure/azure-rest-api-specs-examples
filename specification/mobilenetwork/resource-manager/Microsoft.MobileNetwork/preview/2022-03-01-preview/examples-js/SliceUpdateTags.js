const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update slice tags.
 *
 * @summary Update slice tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SliceUpdateTags.json
 */
async function updateMobileNetworkSliceTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const sliceName = "testSlice";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.slices.updateTags(
    resourceGroupName,
    mobileNetworkName,
    sliceName,
    parameters
  );
  console.log(result);
}

updateMobileNetworkSliceTags().catch(console.error);
