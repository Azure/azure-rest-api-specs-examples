const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists the storage classification mappings in the vault.
 *
 * @summary Lists the storage classification mappings in the vault.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationStorageClassificationMappings_List.json
 */
async function getsTheListOfStorageClassificationMappingsObjectsUnderAVault() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "9112a37f-0f3e-46ec-9c00-060c6edca071";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const resourceName = "vault1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.replicationStorageClassificationMappings.list(
    resourceGroupName,
    resourceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
