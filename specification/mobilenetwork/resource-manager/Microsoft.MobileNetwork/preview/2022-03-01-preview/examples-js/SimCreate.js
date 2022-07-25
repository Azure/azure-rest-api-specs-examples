const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Sim.
 *
 * @summary Creates or updates a Sim.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimCreate.json
 */
async function createSim() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const simName = "testSim";
  const parameters = {
    authenticationKey: "00000000000000000000000000000000",
    deviceType: "Video camera",
    integratedCircuitCardIdentifier: "8900000000000000000",
    internationalMobileSubscriberIdentity: "00000",
    location: "testLocation",
    mobileNetwork: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork",
    },
    operatorKeyCode: "00000000000000000000000000000000",
    simPolicy: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy",
    },
    staticIpConfiguration: [
      {
        attachedDataNetwork: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork",
        },
        slice: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice",
        },
        staticIp: { ipv4Address: "2.4.0.1" },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sims.beginCreateOrUpdateAndWait(
    resourceGroupName,
    simName,
    parameters
  );
  console.log(result);
}

createSim().catch(console.error);
