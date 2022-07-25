const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified sim.
 *
 * @summary Deletes the specified sim.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimDelete.json
 */
async function deleteSim() {
  const subscriptionId = "subid";
  const resourceGroupName = "testResourceGroupName";
  const simName = "testSim";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sims.beginDeleteAndWait(resourceGroupName, simName);
  console.log(result);
}

deleteSim().catch(console.error);
