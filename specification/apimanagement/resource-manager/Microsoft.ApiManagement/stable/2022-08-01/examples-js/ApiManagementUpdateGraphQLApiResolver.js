const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the details of the resolver in the GraphQL API specified by its identifier.
 *
 * @summary Updates the details of the resolver in the GraphQL API specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateGraphQLApiResolver.json
 */
async function apiManagementUpdateGraphQlApiResolver() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "echo-api";
  const resolverId = "resolverId";
  const ifMatch = "*";
  const parameters = {
    path: "Query/adminUsers",
    description: "A GraphQL Resolver example",
    displayName: "Query AdminUsers",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.graphQLApiResolver.update(
    resourceGroupName,
    serviceName,
    apiId,
    resolverId,
    ifMatch,
    parameters
  );
  console.log(result);
}
