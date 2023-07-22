const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get access review instances
 *
 * @summary Get access review instances
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewInstances.json
 */
async function getAccessReviews() {
  const scope = "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d";
  const scheduleDefinitionId = "265785a7-a81f-4201-8a18-bb0db95982b7";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.scopeAccessReviewInstances.list(scope, scheduleDefinitionId)) {
    resArray.push(item);
  }
  console.log(resArray);
}
