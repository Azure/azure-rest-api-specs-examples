const { ManagementLockClient } = require("@azure/arm-locks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the management locks for a resource or any level below resource.
 *
 * @summary Gets all the management locks for a resource or any level below resource.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_ListAtResourceLevel.json
 */
async function listManagementLocksAtResourceLevel() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourcegroupname";
  const resourceProviderNamespace = "Microsoft.Storage";
  const parentResourcePath = "parentResourcePath";
  const resourceType = "storageAccounts";
  const resourceName = "teststorageaccount";
  const credential = new DefaultAzureCredential();
  const client = new ManagementLockClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managementLocks.listAtResourceLevel(
    resourceGroupName,
    resourceProviderNamespace,
    parentResourcePath,
    resourceType,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagementLocksAtResourceLevel().catch(console.error);
