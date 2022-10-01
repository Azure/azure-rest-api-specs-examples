const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update tenant access information details.
 *
 * @summary Update tenant access information details.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateTenantAccess.json
 */
async function apiManagementCreateTenantAccess() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const accessName = "access";
  const ifMatch = "*";
  const parameters = { enabled: true };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tenantAccess.create(
    resourceGroupName,
    serviceName,
    accessName,
    ifMatch,
    parameters
  );
  console.log(result);
}

apiManagementCreateTenantAccess().catch(console.error);
