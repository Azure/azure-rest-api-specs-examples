const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the tenant configuration. User has to be a Tenant Admin for this operation.
 *
 * @summary Delete the tenant configuration. User has to be a Tenant Admin for this operation.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/TenantConfiguration/DeleteTenantConfiguration.json
 */
async function deleteTenantConfiguration() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const result = await client.tenantConfigurations.delete(configurationName);
  console.log(result);
}

deleteTenantConfiguration().catch(console.error);
