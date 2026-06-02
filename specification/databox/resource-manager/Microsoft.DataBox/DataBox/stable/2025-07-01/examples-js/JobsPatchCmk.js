const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the properties of an existing job.
 *
 * @summary updates the properties of an existing job.
 * x-ms-original-file: 2025-07-01/JobsPatchCmk.json
 */
async function jobsPatchCmk() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "YourSubscriptionId";
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.update("YourResourceGroupName", "TestJobName1", {
    details: {
      keyEncryptionKey: {
        kekType: "CustomerManaged",
        kekUrl: "https://xxx.xxx.xx",
        kekVaultResourceID:
          "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.KeyVault/vaults/YourKeyVaultName",
      },
    },
  });
  console.log(result);
}
