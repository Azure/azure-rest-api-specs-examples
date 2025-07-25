const { PrivateLinkScopesClient } = require("@azure/arm-kubernetesconfiguration-privatelinkscopes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
 *
 * @summary Updates an existing PrivateLinkScope's tags. To update other fields use the CreateOrUpdate method.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateLinkScopesUpdateTagsOnly.json
 */
async function privateLinkScopeUpdateTagsOnly() {
  const subscriptionId = process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const privateLinkScopeTags = {
    tags: { tag1: "Value1", tag2: "Value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateLinkScopesClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.updateTags(
    resourceGroupName,
    scopeName,
    privateLinkScopeTags,
  );
  console.log(result);
}
