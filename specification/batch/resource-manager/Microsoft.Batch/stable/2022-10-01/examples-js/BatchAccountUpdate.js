const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing Batch account.
 *
 * @summary Updates the properties of an existing Batch account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/BatchAccountUpdate.json
 */
async function batchAccountUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const parameters = {
    autoStorage: {
      storageAccountId:
        "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.batchAccountOperations.update(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

batchAccountUpdate().catch(console.error);
