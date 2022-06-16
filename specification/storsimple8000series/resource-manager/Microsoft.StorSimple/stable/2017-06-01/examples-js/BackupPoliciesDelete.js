const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the backup policy.
 *
 * @summary Deletes the backup policy.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupPoliciesDelete.json
 */
async function backupPoliciesDelete() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const backupPolicyName = "BkUpPolicy01ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.backupPolicies.beginDeleteAndWait(
    deviceName,
    backupPolicyName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

backupPoliciesDelete().catch(console.error);
