const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates SIM group tags.
 *
 * @summary Updates SIM group tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimGroupUpdateTags.json
 */
async function updateSimGroupTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const simGroupName = "testSimGroup";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.simGroups.updateTags(resourceGroupName, simGroupName, parameters);
  console.log(result);
}

updateSimGroupTags().catch(console.error);
