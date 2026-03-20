const { PrivateLinkScopesClient } = require("@azure/arm-kubernetesconfiguration-privatelinkscopes");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates (or updates) a Azure Arc PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 *
 * @summary creates (or updates) a Azure Arc PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 * x-ms-original-file: 2024-11-01-preview/PrivateLinkScopesCreate.json
 */
async function privateLinkScopeCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const client = new PrivateLinkScopesClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.createOrUpdate(
    "my-resource-group",
    "my-privatelinkscope",
    { location: "westus" },
  );
  console.log(result);
}
