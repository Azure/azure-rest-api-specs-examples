const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all Domains resources under the parent EmailServices resource.
 *
 * @summary Handles requests to list all Domains resources under the parent EmailServices resource.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/domains/listByEmailService.json
 */
async function listDomainsResourcesByEmailServiceName() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const emailServiceName = "MyEmailServiceResource";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.domains.listByEmailServiceResource(
    resourceGroupName,
    emailServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDomainsResourcesByEmailServiceName().catch(console.error);
