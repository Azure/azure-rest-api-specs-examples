const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Operation to add disks(s) to the replication protected item.
 *
 * @summary Operation to add disks(s) to the replication protected item.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectedItems_AddDisks.json
 */
async function addDiskSForProtection() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const resourceName = "vault1";
  const fabricName = "cloud1";
  const protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
  const replicatedProtectedItemName = "f8491e4f-817a-40dd-a90c-af773978c75b";
  const addDisksInput = {
    properties: {
      providerSpecificDetails: {
        instanceType: "A2A",
        vmDisks: [
          {
            diskUri: "https://vmstorage.blob.core.windows.net/vhds/datadisk1.vhd",
            primaryStagingAzureStorageAccountId:
              "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/primaryResource/providers/Microsoft.Storage/storageAccounts/vmcachestorage",
            recoveryAzureStorageAccountId:
              "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourcegroups/recoveryResource/providers/Microsoft.Storage/storageAccounts/recoverystorage",
          },
        ],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectedItems.beginAddDisksAndWait(
    resourceGroupName,
    resourceName,
    fabricName,
    protectionContainerName,
    replicatedProtectedItemName,
    addDisksInput,
  );
  console.log(result);
}
