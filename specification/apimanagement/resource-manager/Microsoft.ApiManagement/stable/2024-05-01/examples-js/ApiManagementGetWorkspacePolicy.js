const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the policy configuration at the API level.
 *
 * @summary Get the policy configuration at the API level.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspacePolicy.json
 */
async function apiManagementGetWorkspacePolicy() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const policyId = "policy";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspacePolicy.get(
    resourceGroupName,
    serviceName,
    workspaceId,
    policyId,
  );
  console.log(result);
}
