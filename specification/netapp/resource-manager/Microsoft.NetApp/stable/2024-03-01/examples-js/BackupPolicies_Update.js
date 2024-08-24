const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a backup policy for Netapp Account
 *
 * @summary Patch a backup policy for Netapp Account
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/BackupPolicies_Update.json
 */
async function backupPoliciesUpdate() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const backupPolicyName = "backupPolicyName";
  const body = {
    dailyBackupsToKeep: 5,
    enabled: false,
    location: "westus",
    monthlyBackupsToKeep: 10,
    weeklyBackupsToKeep: 10,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backupPolicies.beginUpdateAndWait(
    resourceGroupName,
    accountName,
    backupPolicyName,
    body,
  );
  console.log(result);
}
