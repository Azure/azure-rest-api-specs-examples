const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of a migration item.
 *
 * @summary Gets the details of a migration item.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationMigrationItems_Get.json
 */
async function getsTheDetailsOfAMigrationItem() {
  const subscriptionId = "cb53d0c3-bd59-4721-89bc-06916a9147ef";
  const resourceName = "migrationvault";
  const resourceGroupName = "resourcegroup1";
  const fabricName = "vmwarefabric1";
  const protectionContainerName = "vmwareContainer1";
  const migrationItemName = "virtualmachine1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationMigrationItems.get(
    resourceName,
    resourceGroupName,
    fabricName,
    protectionContainerName,
    migrationItemName
  );
  console.log(result);
}

getsTheDetailsOfAMigrationItem().catch(console.error);
