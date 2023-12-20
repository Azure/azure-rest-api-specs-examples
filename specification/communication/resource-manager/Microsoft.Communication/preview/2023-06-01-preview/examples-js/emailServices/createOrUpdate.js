const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new EmailService or update an existing EmailService.
 *
 * @summary Create a new EmailService or update an existing EmailService.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-06-01-preview/examples/emailServices/createOrUpdate.json
 */
async function createOrUpdateEmailServiceResource() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const parameters = {
    dataLocation: "United States",
    location: "Global",
  };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.emailServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    emailServiceName,
    parameters
  );
  console.log(result);
}
