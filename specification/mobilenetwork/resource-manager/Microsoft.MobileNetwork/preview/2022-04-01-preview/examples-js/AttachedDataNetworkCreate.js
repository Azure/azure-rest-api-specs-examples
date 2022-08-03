const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an attached data network.
 *
 * @summary Creates or updates an attached data network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/AttachedDataNetworkCreate.json
 */
async function createAttachedDataNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const packetCoreDataPlaneName = "TestPacketCoreDP";
  const attachedDataNetworkName = "TestAttachedDataNetwork";
  const parameters = {
    dnsAddresses: ["1.1.1.1"],
    location: "eastus",
    naptConfiguration: {
      enabled: "Enabled",
      pinholeLimits: 65536,
      pinholeTimeouts: { icmp: 30, tcp: 180, udp: 30 },
      portRange: { maxPort: 49999, minPort: 1024 },
      portReuseHoldTime: { tcp: 120, udp: 60 },
    },
    userEquipmentAddressPoolPrefix: ["2.2.0.0/16"],
    userEquipmentStaticAddressPoolPrefix: ["2.4.0.0/16"],
    userPlaneDataInterface: { name: "N6" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.attachedDataNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName,
    attachedDataNetworkName,
    parameters
  );
  console.log(result);
}

createAttachedDataNetwork().catch(console.error);
