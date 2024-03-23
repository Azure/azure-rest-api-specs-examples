const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PoolCreate_SecurityProfile.json
 */
async function createPoolSecurityProfile() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        imageReference: {
          offer: "UbuntuServer",
          publisher: "Canonical",
          sku: "18_04-lts-gen2",
          version: "latest",
        },
        nodeAgentSkuId: "batch.node.ubuntu 18.04",
        securityProfile: {
          encryptionAtHost: true,
          securityType: "trustedLaunch",
          uefiSettings: { secureBootEnabled: undefined, vTpmEnabled: false },
        },
      },
    },
    scaleSettings: {
      fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 },
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
