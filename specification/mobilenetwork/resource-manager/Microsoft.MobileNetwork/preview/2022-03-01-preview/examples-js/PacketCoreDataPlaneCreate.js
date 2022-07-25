const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a PacketCoreDataPlane.
 *
 * @summary Creates or updates a PacketCoreDataPlane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/PacketCoreDataPlaneCreate.json
 */
async function createPacketCoreDataPlane() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "testPacketCoreCP";
  const packetCoreDataPlaneName = "testPacketCoreDP";
  const parameters = {
    location: "eastus",
    userPlaneAccessInterface: { name: "N3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreDataPlanes.beginCreateOrUpdateAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName,
    parameters
  );
  console.log(result);
}

createPacketCoreDataPlane().catch(console.error);
