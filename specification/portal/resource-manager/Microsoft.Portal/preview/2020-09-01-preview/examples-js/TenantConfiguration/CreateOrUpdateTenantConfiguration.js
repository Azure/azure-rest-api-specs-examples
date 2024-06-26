const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the tenant configuration. If configuration already exists - update it. User has to be a Tenant Admin for this operation.
 *
 * @summary Create the tenant configuration. If configuration already exists - update it. User has to be a Tenant Admin for this operation.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/TenantConfiguration/CreateOrUpdateTenantConfiguration.json
 */
async function createOrUpdateTenantConfiguration() {
  const subscriptionId =
    process.env["PORTAL_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const configurationName = "default";
  const tenantConfiguration = {
    enforcePrivateMarkdownStorage: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const result = await client.tenantConfigurations.create(configurationName, tenantConfiguration);
  console.log(result);
}
