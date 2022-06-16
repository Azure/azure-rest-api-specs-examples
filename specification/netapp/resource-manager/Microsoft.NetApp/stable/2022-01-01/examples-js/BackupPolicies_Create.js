const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a backup policy for Netapp Account
 *
 * @summary Create a backup policy for Netapp Account
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/BackupPolicies_Create.json
 */
async function backupPoliciesCreate() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const backupPolicyName = "backupPolicyName";
  const body = {
    dailyBackupsToKeep: 10,
    enabled: true,
    location: "westus",
    monthlyBackupsToKeep: 10,
    weeklyBackupsToKeep: 10,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backupPolicies.beginCreateAndWait(
    resourceGroupName,
    accountName,
    backupPolicyName,
    body
  );
  console.log(result);
}

backupPoliciesCreate().catch(console.error);
