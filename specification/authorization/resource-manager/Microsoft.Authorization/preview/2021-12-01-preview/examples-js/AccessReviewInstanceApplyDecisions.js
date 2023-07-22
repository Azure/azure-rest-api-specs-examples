const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to An action to apply all decisions for an access review instance.
 *
 * @summary An action to apply all decisions for an access review instance.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/AccessReviewInstanceApplyDecisions.json
 */
async function getAccessReview() {
  const scope = "subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d";
  const scheduleDefinitionId = "fa73e90b-5bf1-45fd-a182-35ce5fc0674d";
  const id = "d9b9e056-7004-470b-bf21-1635e98487da";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.scopeAccessReviewInstance.applyDecisions(
    scope,
    scheduleDefinitionId,
    id
  );
  console.log(result);
}
