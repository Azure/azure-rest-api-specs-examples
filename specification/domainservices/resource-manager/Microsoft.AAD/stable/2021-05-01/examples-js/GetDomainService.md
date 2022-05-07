Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-domainservices_4.0.1/sdk/domainservices/arm-domainservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
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
```
