const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the issue Comment for an API specified by its identifier.
 *
 * @summary Gets the entity state (Etag) version of the issue Comment for an API specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiIssueComment.json
 */
async function apiManagementHeadApiIssueComment() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "57d2ef278aa04f0888cba3f3";
  const issueId = "57d2ef278aa04f0ad01d6cdc";
  const commentId = "599e29ab193c3c0bd0b3e2fb";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiIssueComment.getEntityTag(
    resourceGroupName,
    serviceName,
    apiId,
    issueId,
    commentId
  );
  console.log(result);
}

apiManagementHeadApiIssueComment().catch(console.error);
