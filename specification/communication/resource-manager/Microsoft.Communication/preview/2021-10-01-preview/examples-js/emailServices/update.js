const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an existing EmailService.
 *
 * @summary Operation to update an existing EmailService.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/emailServices/update.json
 */
async function updateEmailServiceResource() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const parameters = { tags: { newTag: "newVal" } };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.emailServices.beginUpdateAndWait(
    resourceGroupName,
    emailServiceName,
    parameters
  );
  console.log(result);
}

updateEmailServiceResource().catch(console.error);
