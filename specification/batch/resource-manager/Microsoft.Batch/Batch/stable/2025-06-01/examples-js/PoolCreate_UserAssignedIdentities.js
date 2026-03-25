const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_UserAssignedIdentities.json
 */
async function createPoolUserAssignedIdentities() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1":
            {},
          "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2":
            {},
        },
      },
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            offer: "UbuntuServer",
            publisher: "Canonical",
            sku: "18.04-LTS",
            version: "latest",
          },
          nodeAgentSkuId: "batch.node.ubuntu 18.04",
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
