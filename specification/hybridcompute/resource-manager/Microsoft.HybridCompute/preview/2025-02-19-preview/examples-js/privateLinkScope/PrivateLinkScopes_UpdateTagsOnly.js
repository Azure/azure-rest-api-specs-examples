const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
 *
 * @summary Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/privateLinkScope/PrivateLinkScopes_UpdateTagsOnly.json
 */
async function privateLinkScopeUpdateTagsOnly() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const privateLinkScopeTags = {
    tags: { tag1: "Value1", tag2: "Value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.updateTags(
    resourceGroupName,
    scopeName,
    privateLinkScopeTags,
  );
  console.log(result);
}
