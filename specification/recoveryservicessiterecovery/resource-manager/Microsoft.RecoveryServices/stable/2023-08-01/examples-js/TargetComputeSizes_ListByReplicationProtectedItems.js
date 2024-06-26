const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available target compute sizes for a replication protected item.
 *
 * @summary Lists the available target compute sizes for a replication protected item.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/TargetComputeSizes_ListByReplicationProtectedItems.json
 */
async function getsTheListOfTargetComputeSizesForTheReplicationProtectedItem() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "6808dbbc-98c7-431f-a1b1-9580902423b7";
  const resourceName = "avraiMgDiskVault";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "avraiMgDiskVaultRG";
  const fabricName = "asr-a2a-default-centraluseuap";
  const protectionContainerName = "asr-a2a-default-centraluseuap-container";
  const replicatedProtectedItemName = "468c912d-b1ab-4ea2-97eb-4b5095155db2";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.targetComputeSizes.listByReplicationProtectedItems(
    resourceName,
    resourceGroupName,
    fabricName,
    protectionContainerName,
    replicatedProtectedItemName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
