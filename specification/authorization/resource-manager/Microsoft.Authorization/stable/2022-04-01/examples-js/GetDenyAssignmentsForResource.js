const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets deny assignments for a resource.
 *
 * @summary Gets deny assignments for a resource.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetDenyAssignmentsForResource.json
 */
async function listDenyAssignmentsForResource() {
  const subscriptionId = process.env["AUTHORIZATION_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["AUTHORIZATION_RESOURCE_GROUP"] || "rgname";
  const resourceProviderNamespace = "resourceProviderNamespace";
  const parentResourcePath = "parentResourcePath";
  const resourceType = "resourceType";
  const resourceName = "resourceName";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.denyAssignments.listForResource(
    resourceGroupName,
    resourceProviderNamespace,
    parentResourcePath,
    resourceType,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
