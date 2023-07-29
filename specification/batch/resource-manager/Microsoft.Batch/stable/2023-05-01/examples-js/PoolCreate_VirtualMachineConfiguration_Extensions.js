const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolCreate_VirtualMachineConfiguration_Extensions.json
 */
async function createPoolVirtualMachineConfigurationExtensions() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
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
      autoScale: {
        evaluationInterval: "PT5M",
        formula: "$TargetDedicatedNodes=1",
      },
    },
    targetNodeCommunicationMode: "Default",
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
