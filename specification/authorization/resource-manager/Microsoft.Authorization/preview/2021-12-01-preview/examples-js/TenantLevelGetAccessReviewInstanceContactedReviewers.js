const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get access review instance contacted reviewers
 *
 * @summary Get access review instance contacted reviewers
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/TenantLevelGetAccessReviewInstanceContactedReviewers.json
 */
async function getAccessReviews() {
  const scheduleDefinitionId = "265785a7-a81f-4201-8a18-bb0db95982b7";
  const id = "f25ed880-9c31-4101-bc57-825d8df3b58c";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.tenantLevelAccessReviewInstanceContactedReviewers.list(
    scheduleDefinitionId,
    id
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
