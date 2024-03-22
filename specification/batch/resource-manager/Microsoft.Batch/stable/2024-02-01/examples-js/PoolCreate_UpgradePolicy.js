const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PoolCreate_UpgradePolicy.json
 */
async function createPoolUpgradePolicy() {
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
        nodePlacementConfiguration: { policy: "Zonal" },
        windowsConfiguration: { enableAutomaticUpdates: false },
      },
    },
    scaleSettings: {
      fixedScale: { targetDedicatedNodes: 2, targetLowPriorityNodes: 0 },
    },
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
