const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete backup policy
 *
 * @summary Delete backup policy
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/BackupPolicies_Delete.json
 */
async function backupsDelete() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "resourceGroup";
  const accountName = "accountName";
  const backupPolicyName = "backupPolicyName";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.backupPolicies.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    backupPolicyName
  );
  console.log(result);
}

backupsDelete().catch(console.error);
