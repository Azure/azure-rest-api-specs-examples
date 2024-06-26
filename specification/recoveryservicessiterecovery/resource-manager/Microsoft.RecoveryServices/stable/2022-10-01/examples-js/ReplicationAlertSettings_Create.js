const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an email notification(alert) configuration.
 *
 * @summary Create or update an email notification(alert) configuration.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationAlertSettings_Create.json
 */
async function configuresEmailNotificationsForThisVault() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const alertSettingName = "defaultAlertSetting";
  const request = {
    properties: {
      customEmailAddresses: ["ronehr@microsoft.com"],
      locale: "",
      sendToOwners: "false",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationAlertSettings.create(
    resourceName,
    resourceGroupName,
    alertSettingName,
    request
  );
  console.log(result);
}

configuresEmailNotificationsForThisVault().catch(console.error);
