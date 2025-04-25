const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the policy restriction configuration of the Api Management Service.
 *
 * @summary Deletes the policy restriction configuration of the Api Management Service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeletePolicyRestriction.json
 */
async function apiManagementDeletePolicyRestriction() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const policyRestrictionId = "policyRestriction1";
  const ifMatch = "*";
  const options = { ifMatch };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.policyRestriction.delete(
    resourceGroupName,
    serviceName,
    policyRestrictionId,
    options,
  );
  console.log(result);
}
