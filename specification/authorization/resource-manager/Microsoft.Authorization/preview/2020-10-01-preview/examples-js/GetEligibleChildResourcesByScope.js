const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the child resources of a resource on which user has eligible access
 *
 * @summary Get the child resources of a resource on which user has eligible access
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-10-01-preview/examples/GetEligibleChildResourcesByScope.json
 */
async function getEligibleChildResourcesByScope() {
  const scope =
    "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
  const filter = "resourceType+eq+'resourcegroup'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.eligibleChildResources.list(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
