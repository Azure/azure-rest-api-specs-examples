```javascript
const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolCreate_UserAssignedIdentities.json
 */
async function createPoolUserAssignedIdentities() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
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
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subid/resourceGroups/defaultAzurebatchJapaneast/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
        "/subscriptions/subid/resourceGroups/defaultAzurebatchJapaneast/providers/MicrosoftManagedIdentity/userAssignedIdentities/id2":
          {},
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

createPoolUserAssignedIdentities().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-batch_7.1.0/sdk/batch/arm-batch/README.md) on how to add the SDK to your project and authenticate.
