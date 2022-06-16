const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolCreate_VirtualMachineConfiguration_Extensions.json
 */
async function createPoolVirtualMachineConfigurationExtensions() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        imageReference: {
          offer: "0001-com-ubuntu-server-focal",
          publisher: "Canonical",
          sku: "20_04-lts",
        },
        nodeAgentSkuId: "batch.node.ubuntu 20.04",
        extensions: [
          {
            name: "batchextension1",
            type: "SecurityMonitoringForLinux",
            autoUpgradeMinorVersion: true,
            protectedSettings: {
              protectedSettingsKey: "protectedSettingsValue",
            },
            publisher: "Microsoft.Azure.Security.Monitoring",
            settings: { settingsKey: "settingsValue" },
            typeHandlerVersion: "1.0",
          },
        ],
      },
    },
    scaleSettings: {
      autoScale: {
        evaluationInterval: "PT5M",
        formula: "$TargetDedicatedNodes=1",
      },
    },
    vmSize: "STANDARD_D4",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.create(
    resourceGroupName,
    accountName,
    poolName,
    parameters
  );
  console.log(result);
}

createPoolVirtualMachineConfigurationExtensions().catch(console.error);
