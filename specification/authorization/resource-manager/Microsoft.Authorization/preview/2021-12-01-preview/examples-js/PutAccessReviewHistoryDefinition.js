const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a scheduled or one-time Access Review History Definition
 *
 * @summary Create a scheduled or one-time Access Review History Definition
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/PutAccessReviewHistoryDefinition.json
 */
async function putAccessReviewHistoryDefinition() {
  const scope = "subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a";
  const historyDefinitionId = "44724910-d7a5-4c29-b28f-db73e717165a";
  const properties = {};
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.scopeAccessReviewHistoryDefinition.create(
    scope,
    historyDefinitionId,
    properties
  );
  console.log(result);
}
