const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PoolCreate_VirtualMachineConfiguration_ServiceArtifactReference.json
 */
async function createPoolVirtualMachineConfigurationServiceArtifactReference() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
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
          id: "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/vmArtifactsProfile",
        },
        windowsConfiguration: { enableAutomaticUpdates: false },
      },
    },
    scaleSettings: {
      fixedScale: { targetDedicatedNodes: 2, targetLowPriorityNodes: 0 },
    },
    upgradePolicy: {
      automaticOSUpgradePolicy: { enableAutomaticOSUpgrade: true },
      mode: "automatic",
    },
    vmSize: "Standard_d4s_v3",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.create(
    resourceGroupName,
    accountName,
    poolName,
    parameters,
  );
  console.log(result);
}
