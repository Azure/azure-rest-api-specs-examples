const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the policy configuration at the GraphQL API Resolver level.
 *
 * @summary Get the policy configuration at the GraphQL API Resolver level.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetGraphQLApiResolverPolicy.json
 */
async function apiManagementGetGraphQlApiResolverPolicy() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "5600b539c53f5b0062040001";
  const resolverId = "5600b53ac53f5b0062080006";
  const policyId = "policy";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.graphQLApiResolverPolicy.get(
    resourceGroupName,
    serviceName,
    apiId,
    resolverId,
    policyId
  );
  console.log(result);
}
