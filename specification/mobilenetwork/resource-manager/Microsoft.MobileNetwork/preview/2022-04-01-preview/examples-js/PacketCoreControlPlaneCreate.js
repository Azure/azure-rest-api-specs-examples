const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a packet core control plane.
 *
 * @summary Creates or updates a packet core control plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/PacketCoreControlPlaneCreate.json
 */
async function createPacketCoreControlPlane() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const parameters = {
    controlPlaneAccessInterface: { name: "N2" },
    coreNetworkTechnology: "5GC",
    localDiagnosticsAccess: {
      httpsServerCertificate: {
        certificateUrl: "https://contosovault.vault.azure.net/certificates/ingress",
      },
    },
    location: "eastus",
    mobileNetwork: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork",
    },
    platform: {
      type: "AKS-HCI",
      azureStackEdgeDevice: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/TestAzureStackEdgeDevice",
      },
      connectedCluster: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/TestConnectedCluster",
      },
      customLocation: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation",
      },
    },
    sku: "testSku",
    version: "0.2.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreControlPlanes.beginCreateOrUpdateAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    parameters
  );
  console.log(result);
}

createPacketCoreControlPlane().catch(console.error);
