const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the tenant configuration.
 *
 * @summary Gets the tenant configuration.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/TenantConfiguration/GetTenantConfiguration.json
 */
async function getTenantConfiguration() {
  const subscriptionId =
    process.env["PORTAL_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const result = await client.tenantConfigurations.get(configurationName);
  console.log(result);
}
