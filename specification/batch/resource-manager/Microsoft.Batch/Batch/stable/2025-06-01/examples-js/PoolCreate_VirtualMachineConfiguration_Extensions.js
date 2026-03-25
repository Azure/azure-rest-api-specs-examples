const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_VirtualMachineConfiguration_Extensions.json
 */
async function createPoolVirtualMachineConfigurationExtensions() {
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
            offer: "0001-com-ubuntu-server-focal",
            publisher: "Canonical",
            sku: "20_04-lts",
          },
          nodeAgentSkuId: "batch.node.ubuntu 20.04",
          extensions: [
            {
              name: "batchextension1",
              type: "KeyVaultForLinux",
              autoUpgradeMinorVersion: true,
              enableAutomaticUpgrade: true,
              publisher: "Microsoft.Azure.KeyVault",
              settings: {
                authenticationSettingsKey: "authenticationSettingsValue",
                secretsManagementSettingsKey: "secretsManagementSettingsValue",
              },
              typeHandlerVersion: "2.0",
            },
          ],
        },
      },
      scaleSettings: {
        autoScale: { evaluationInterval: "PT5M", formula: "$TargetDedicatedNodes=1" },
      },
      vmSize: "STANDARD_D4",
    },
  );
  console.log(result);
}
