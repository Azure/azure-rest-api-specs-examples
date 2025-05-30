const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update tenant access information details.
 *
 * @summary Update tenant access information details.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateTenantAccess.json
 */
async function apiManagementCreateTenantAccess() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
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
    parameters,
  );
  console.log(result);
}
