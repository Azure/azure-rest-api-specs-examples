const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add a new Domains resource under the parent EmailService resource or update an existing Domains resource.
 *
 * @summary Add a new Domains resource under the parent EmailService resource or update an existing Domains resource.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/domains/createOrUpdate.json
 */
async function createOrUpdateDomainsResource() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const domainName = "mydomain.com";
  const parameters = {
    domainManagement: "CustomerManaged",
    location: "Global",
  };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.domains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    emailServiceName,
    domainName,
    parameters
  );
  console.log(result);
}

createOrUpdateDomainsResource().catch(console.error);
