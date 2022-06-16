const { DomainServicesResourceProvider } = require("@azure/arm-domainservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Get Domain Service operation retrieves a json representation of the Domain Service.
 *
 * @summary The Get Domain Service operation retrieves a json representation of the Domain Service.
 * x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetDomainService.json
 */
async function getDomainService() {
  const subscriptionId = "1639790a-76a2-4ac4-98d9-8562f5dfcb4d";
  const resourceGroupName = "TestResourceGroup";
  const domainServiceName = "TestDomainService.com";
  const credential = new DefaultAzureCredential();
  const client = new DomainServicesResourceProvider(credential, subscriptionId);
  const result = await client.domainServices.get(resourceGroupName, domainServiceName);
  console.log(result);
}

getDomainService().catch(console.error);
