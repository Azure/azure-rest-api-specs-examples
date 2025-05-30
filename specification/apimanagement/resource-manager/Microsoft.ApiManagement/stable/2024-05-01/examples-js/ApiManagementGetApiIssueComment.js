const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the details of the issue Comment for an API specified by its identifier.
 *
 * @summary Gets the details of the issue Comment for an API specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetApiIssueComment.json
 */
async function apiManagementGetApiIssueComment() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "57d2ef278aa04f0888cba3f3";
  const issueId = "57d2ef278aa04f0ad01d6cdc";
  const commentId = "599e29ab193c3c0bd0b3e2fb";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiIssueComment.get(
    resourceGroupName,
    serviceName,
    apiId,
    issueId,
    commentId,
  );
  console.log(result);
}
