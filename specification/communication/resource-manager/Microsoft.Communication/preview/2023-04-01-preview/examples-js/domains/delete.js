const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a Domains resource.
 *
 * @summary Operation to delete a Domains resource.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/domains/delete.json
 */
async function deleteDomainsResource() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const domainName = "mydomain.com";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.domains.beginDeleteAndWait(
    resourceGroupName,
    emailServiceName,
    domainName
  );
  console.log(result);
}
