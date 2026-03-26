const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_CustomerManagedKey_ForBatchManagedAccounts.json
 */
async function createPoolCustomerManagedKeyForBatchManagedAccounts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      vmSize: "Standard_D4ds_v5",
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            sku: "2022-Datacenter",
            publisher: "MicrosoftWindowsServer",
            version: "latest",
            offer: "WindowsServer",
          },
          nodeAgentSkuId: "batch.node.windows amd64",
          diskEncryptionConfiguration: {
            targets: ["OsDisk"],
            customerManagedKey: {
              keyUrl: "http://sample.vault.azure.net//keys/cmk/bb60031a6d4545d3a60d3f94588538c9",
              identityReference: {
                resourceId:
                  "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
              },
            },
          },
        },
      },
      scaleSettings: { fixedScale: { targetDedicatedNodes: 1, resizeTimeout: "PT15M" } },
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1":
            {},
        },
      },
    },
  );
  console.log(result);
}
