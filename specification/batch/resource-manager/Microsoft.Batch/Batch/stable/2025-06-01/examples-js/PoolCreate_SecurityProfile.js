const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_SecurityProfile.json
 */
async function createPoolSecurityProfile() {
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
            offer: "UbuntuServer",
            publisher: "Canonical",
            sku: "18_04-lts-gen2",
            version: "latest",
          },
          nodeAgentSkuId: "batch.node.ubuntu 18.04",
          securityProfile: {
            encryptionAtHost: true,
            securityType: "trustedLaunch",
            uefiSettings: { vTpmEnabled: false },
          },
        },
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 1, targetLowPriorityNodes: 0 } },
      vmSize: "Standard_d4s_v3",
    },
  );
  console.log(result);
}
