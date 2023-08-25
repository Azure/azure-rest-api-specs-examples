const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new resolver in the GraphQL API or updates an existing one.
 *
 * @summary Creates a new resolver in the GraphQL API or updates an existing one.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGraphQLApiResolver.json
 */
async function apiManagementCreateGraphQlApiResolver() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "someAPI";
  const resolverId = "newResolver";
  const parameters = {
    path: "Query/users",
    description: "A GraphQL Resolver example",
    displayName: "Query Users",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.graphQLApiResolver.createOrUpdate(
    resourceGroupName,
    serviceName,
    apiId,
    resolverId,
    parameters
  );
  console.log(result);
}
