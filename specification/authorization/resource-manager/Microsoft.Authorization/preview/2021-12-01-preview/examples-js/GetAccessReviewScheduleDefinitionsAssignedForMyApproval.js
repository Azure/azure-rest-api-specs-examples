const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get access review instances assigned for my approval.
 *
 * @summary Get access review instances assigned for my approval.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewScheduleDefinitionsAssignedForMyApproval.json
 */
async function getAccessReviews() {
  const filter = "assignedToMeToReview()";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.accessReviewScheduleDefinitionsAssignedForMyApproval.list(
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
