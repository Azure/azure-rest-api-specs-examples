const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the list of migration items in the vault.
 *
 * @summary Gets the list of migration items in the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationMigrationItems_List.json
 */
async function getsTheListOfMigrationItemsInTheVault() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "cb53d0c3-bd59-4721-89bc-06916a9147ef";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourcegroup1";
  const resourceName = "migrationvault";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.replicationMigrationItems.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
