const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the policy restriction configuration of the Api Management service.
 *
 * @summary Updates the policy restriction configuration of the Api Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdatePolicyRestriction.json
 */
async function apiManagementUpdatePolicyRestriction() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const policyRestrictionId = "policyRestriction1";
  const ifMatch = "*";
  const parameters = {
    scope: "Sample Path 2 to the policy document.",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.policyRestriction.update(
    resourceGroupName,
    serviceName,
    policyRestrictionId,
    ifMatch,
    parameters,
  );
  console.log(result);
}
