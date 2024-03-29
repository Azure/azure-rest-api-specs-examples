const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Azure Arc PrivateLinkScopes within a resource group.
 *
 * @summary Gets a list of Azure Arc PrivateLinkScopes within a resource group.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/PrivateLinkScopesListByResourceGroup.json
 */
async function privateLinkScopeListByResourceGroup() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName = "my-resource-group";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkScopes.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateLinkScopeListByResourceGroup().catch(console.error);
