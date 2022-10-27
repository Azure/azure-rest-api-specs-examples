const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates a disk encryption set
 *
 * @summary Creates or updates a disk encryption set
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentSubscription.json
 */
async function createADiskEncryptionSetWithKeyVaultFromADifferentSubscription() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const options = {
    body: {
      identity: { type: "SystemAssigned" },
      location: "West US",
      properties: {
        activeKey: {
          keyUrl: "https://myvaultdifferentsub.vault-int.azure-int.net/keys/{key}",
        },
        encryptionType: "EncryptionAtRestWithCustomerKey",
      },
    },
    queryParameters: { "api-version": "2022-07-02" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}",
      subscriptionId,
      resourceGroupName,
      diskEncryptionSetName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createADiskEncryptionSetWithKeyVaultFromADifferentSubscription().catch(console.error);
