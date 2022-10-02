const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get tenant settings.
 *
 * @summary Get tenant settings.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetTenantSettings.json
 */
async function apiManagementGetTenantSettings() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const settingsType = "public";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tenantSettings.get(resourceGroupName, serviceName, settingsType);
  console.log(result);
}

apiManagementGetTenantSettings().catch(console.error);
