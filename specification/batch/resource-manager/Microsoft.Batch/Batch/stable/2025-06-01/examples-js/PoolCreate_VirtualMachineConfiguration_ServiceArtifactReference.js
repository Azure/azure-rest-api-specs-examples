const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_VirtualMachineConfiguration_ServiceArtifactReference.json
 */
async function createPoolVirtualMachineConfigurationServiceArtifactReference() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            offer: "WindowsServer",
            publisher: "MicrosoftWindowsServer",
            sku: "2019-datacenter-smalldisk",
            version: "latest",
          },
          nodeAgentSkuId: "batch.node.windows amd64",
          serviceArtifactReference: {
            id: "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/vmArtifactsProfile",
          },
          windowsConfiguration: { enableAutomaticUpdates: false },
        },
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 2, targetLowPriorityNodes: 0 } },
      upgradePolicy: {
        automaticOSUpgradePolicy: { enableAutomaticOSUpgrade: true },
        mode: "automatic",
      },
      vmSize: "Standard_d4s_v3",
    },
  );
  console.log(result);
}
