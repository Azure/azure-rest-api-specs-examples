const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all attachments for the Issue associated with the specified API.
 *
 * @summary Lists all attachments for the Issue associated with the specified API.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiIssueAttachments.json
 */
async function apiManagementListApiIssueAttachments() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "57d1f7558aa04f15146d9d8a";
  const issueId = "57d2ef278aa04f0ad01d6cdc";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.apiIssueAttachment.listByService(
    resourceGroupName,
    serviceName,
    apiId,
    issueId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
