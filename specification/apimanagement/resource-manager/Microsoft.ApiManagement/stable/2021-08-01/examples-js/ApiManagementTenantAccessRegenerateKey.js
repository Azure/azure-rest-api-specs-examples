const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate primary access key for GIT.
 *
 * @summary Regenerate primary access key for GIT.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementTenantAccessRegenerateKey.json
 */
async function apiManagementTenantAccessRegenerateKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const accessName = "access";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tenantAccessGit.regeneratePrimaryKey(
    resourceGroupName,
    serviceName,
    accessName
  );
  console.log(result);
}

apiManagementTenantAccessRegenerateKey().catch(console.error);
