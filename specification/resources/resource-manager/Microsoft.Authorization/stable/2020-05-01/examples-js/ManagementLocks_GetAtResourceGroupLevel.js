const { ManagementLockClient } = require("@azure/arm-locks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a management lock at the resource group level.
 *
 * @summary Gets a management lock at the resource group level.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_GetAtResourceGroupLevel.json
 */
async function getManagementLockAtResourceGroupLevel() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourcegroupname";
  const lockName = "testlock";
  const credential = new DefaultAzureCredential();
  const client = new ManagementLockClient(credential, subscriptionId);
  const result = await client.managementLocks.getAtResourceGroupLevel(resourceGroupName, lockName);
  console.log(result);
}

getManagementLockAtResourceGroupLevel().catch(console.error);
