const { PrivateLinkScopesClient } = require("@azure/arm-kubernetesconfiguration-privatelinkscopes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates (or updates) a Azure Arc PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 *
 * @summary Creates (or updates) a Azure Arc PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateLinkScopesCreate.json
 */
async function privateLinkScopeCreate() {
  const subscriptionId =
    process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] ||
    "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName =
    process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const parameters = {
    location: "westus",
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateLinkScopesClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.createOrUpdate(
    resourceGroupName,
    scopeName,
    parameters,
  );
  console.log(result);
}
