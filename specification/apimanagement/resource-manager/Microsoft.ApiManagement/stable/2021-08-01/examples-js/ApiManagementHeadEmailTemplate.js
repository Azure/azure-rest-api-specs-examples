const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the email template specified by its identifier.
 *
 * @summary Gets the entity state (Etag) version of the email template specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadEmailTemplate.json
 */
async function apiManagementHeadEmailTemplate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const templateName = "newIssueNotificationMessage";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.emailTemplate.getEntityTag(
    resourceGroupName,
    serviceName,
    templateName
  );
  console.log(result);
}

apiManagementHeadEmailTemplate().catch(console.error);
