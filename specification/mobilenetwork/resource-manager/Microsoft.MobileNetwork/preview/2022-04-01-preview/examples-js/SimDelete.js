const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified SIM.
 *
 * @summary Deletes the specified SIM.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimDelete.json
 */
async function deleteSim() {
  const subscriptionId = "subid";
  const resourceGroupName = "testResourceGroupName";
  const simGroupName = "testSimGroup";
  const simName = "testSim";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sims.beginDeleteAndWait(resourceGroupName, simGroupName, simName);
  console.log(result);
}

deleteSim().catch(console.error);
