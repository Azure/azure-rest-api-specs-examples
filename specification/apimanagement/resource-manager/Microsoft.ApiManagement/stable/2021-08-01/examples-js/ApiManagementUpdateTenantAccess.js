const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update tenant access information details.
 *
 * @summary Update tenant access information details.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateTenantAccess.json
 */
async function apiManagementUpdateTenantAccess() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const accessName = "access";
  const ifMatch = "*";
  const parameters = { enabled: true };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tenantAccess.update(
    resourceGroupName,
    serviceName,
    accessName,
    ifMatch,
    parameters
  );
  console.log(result);
}

apiManagementUpdateTenantAccess().catch(console.error);
