const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to initiate pause replication of the item.
 *
 * @summary The operation to initiate pause replication of the item.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationMigrationItems_PauseReplication.json
 */
async function pauseReplication() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "cb53d0c3-bd59-4721-89bc-06916a9147ef";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourcegroup1";
  const resourceName = "migrationvault";
  const fabricName = "vmwarefabric1";
  const protectionContainerName = "vmwareContainer1";
  const migrationItemName = "virtualmachine1";
  const pauseReplicationInput = {
    properties: { instanceType: "VMwareCbt" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationMigrationItems.beginPauseReplicationAndWait(
    resourceGroupName,
    resourceName,
    fabricName,
    protectionContainerName,
    migrationItemName,
    pauseReplicationInput,
  );
  console.log(result);
}
