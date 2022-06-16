const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateLinkScopeUpdateTagsOnly() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const privateLinkScopeTags = {
    tags: { tag1: "Value1", tag2: "Value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.updateTags(
    resourceGroupName,
    scopeName,
    privateLinkScopeTags
  );
  console.log(result);
}

privateLinkScopeUpdateTagsOnly().catch(console.error);
