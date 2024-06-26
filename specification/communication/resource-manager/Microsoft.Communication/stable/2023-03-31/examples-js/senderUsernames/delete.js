const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a SenderUsernames resource.
 *
 * @summary Operation to delete a SenderUsernames resource.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/senderUsernames/delete.json
 */
async function deleteSenderUsernamesResource() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const domainName = "mydomain.com";
  const senderUsername = "contosoNewsAlerts";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.senderUsernames.delete(
    resourceGroupName,
    emailServiceName,
    domainName,
    senderUsername
  );
  console.log(result);
}
