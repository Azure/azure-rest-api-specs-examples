const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the vault setting. This includes the Migration Hub connection settings.
 *
 * @summary Gets the vault setting. This includes the Migration Hub connection settings.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationVaultSetting_Get.json
 */
async function getsTheVaultSetting() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const vaultSettingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationVaultSetting.get(
    resourceName,
    resourceGroupName,
    vaultSettingName
  );
  console.log(result);
}

getsTheVaultSetting().catch(console.error);
