const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
 *
 * @summary Gets the private link resources that need to be created for a Azure Monitor PrivateLinkScope.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/privateLinkScope/PrivateLinkScopePrivateLinkResource_Get.json
 */
async function getsPrivateEndpointConnection() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const groupName = "hybridcompute";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, scopeName, groupName);
  console.log(result);
}
