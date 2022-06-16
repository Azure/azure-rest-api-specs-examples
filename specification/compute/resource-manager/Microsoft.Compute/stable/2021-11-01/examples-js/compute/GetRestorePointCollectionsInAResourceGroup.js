const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of restore point collections in a resource group.
 *
 * @summary Gets the list of restore point collections in a resource group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GetRestorePointCollectionsInAResourceGroup.json
 */
async function getsTheListOfRestorePointCollectionsInAResourceGroup() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.restorePointCollections.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheListOfRestorePointCollectionsInAResourceGroup().catch(console.error);
