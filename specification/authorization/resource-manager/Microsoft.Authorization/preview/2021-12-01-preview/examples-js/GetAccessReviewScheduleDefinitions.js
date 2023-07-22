const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get access review schedule definitions
 *
 * @summary Get access review schedule definitions
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewScheduleDefinitions.json
 */
async function getAccessReviews() {
  const scope = "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.scopeAccessReviewScheduleDefinitions.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
