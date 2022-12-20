const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of migration items in the vault.
 *
 * @summary Gets the list of migration items in the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationMigrationItems_List.json
 */
async function getsTheListOfMigrationItemsInTheVault() {
  const subscriptionId = "cb53d0c3-bd59-4721-89bc-06916a9147ef";
  const resourceName = "migrationvault";
  const resourceGroupName = "resourcegroup1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.replicationMigrationItems.list(resourceName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfMigrationItemsInTheVault().catch(console.error);
