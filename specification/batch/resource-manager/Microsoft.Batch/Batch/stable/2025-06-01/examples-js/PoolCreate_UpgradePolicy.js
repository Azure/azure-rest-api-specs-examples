const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_UpgradePolicy.json
 */
async function createPoolUpgradePolicy() {
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
          nodePlacementConfiguration: { policy: "Zonal" },
          windowsConfiguration: { enableAutomaticUpdates: false },
        },
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 2, targetLowPriorityNodes: 0 } },
      upgradePolicy: {
        automaticOSUpgradePolicy: {
          disableAutomaticRollback: true,
          enableAutomaticOSUpgrade: true,
          osRollingUpgradeDeferral: true,
          useRollingUpgradePolicy: true,
        },
        mode: "automatic",
        rollingUpgradePolicy: {
          enableCrossZoneUpgrade: true,
          maxBatchInstancePercent: 20,
          maxUnhealthyInstancePercent: 20,
          maxUnhealthyUpgradedInstancePercent: 20,
          pauseTimeBetweenBatches: "PT0S",
          prioritizeUnhealthyInstances: false,
          rollbackFailedInstancesOnPolicyBreach: false,
        },
      },
      vmSize: "Standard_d4s_v3",
    },
  );
  console.log(result);
}
