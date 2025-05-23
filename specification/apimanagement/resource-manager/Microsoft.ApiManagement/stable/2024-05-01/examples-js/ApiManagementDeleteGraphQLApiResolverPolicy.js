const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the policy configuration at the GraphQL Api Resolver.
 *
 * @summary Deletes the policy configuration at the GraphQL Api Resolver.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteGraphQLApiResolverPolicy.json
 */
async function apiManagementDeleteGraphQlApiResolverPolicy() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "testapi";
  const resolverId = "testResolver";
  const policyId = "policy";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.graphQLApiResolverPolicy.delete(
    resourceGroupName,
    serviceName,
    apiId,
    resolverId,
    policyId,
    ifMatch,
  );
  console.log(result);
}
