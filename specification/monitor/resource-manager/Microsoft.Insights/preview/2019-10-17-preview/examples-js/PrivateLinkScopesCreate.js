const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates (or updates) a Azure Monitor PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 *
 * @summary Creates (or updates) a Azure Monitor PrivateLinkScope. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2019-10-17-preview/examples/PrivateLinkScopesCreate.json
 */
async function privateLinkScopeCreate() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName = "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const azureMonitorPrivateLinkScopePayload = {
    location: "Global",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.createOrUpdate(
    resourceGroupName,
    scopeName,
    azureMonitorPrivateLinkScopePayload
  );
  console.log(result);
}

privateLinkScopeCreate().catch(console.error);
