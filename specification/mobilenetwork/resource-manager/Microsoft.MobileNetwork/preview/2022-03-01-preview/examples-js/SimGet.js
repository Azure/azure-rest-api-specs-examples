const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified sim.
 *
 * @summary Gets information about the specified sim.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimGet.json
 */
async function getSim() {
  const subscriptionId = "subid";
  const resourceGroupName = "testResourceGroupName";
  const simName = "testSimName";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sims.get(resourceGroupName, simName);
  console.log(result);
}

getSim().catch(console.error);
