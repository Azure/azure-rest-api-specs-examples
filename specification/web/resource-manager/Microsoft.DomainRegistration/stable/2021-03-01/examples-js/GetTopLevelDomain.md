Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get details of a top-level domain.
 *
 * @summary Description for Get details of a top-level domain.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2021-03-01/examples/GetTopLevelDomain.json
 */
async function getTopLevelDomain() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const name = "com";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.topLevelDomains.get(name);
  console.log(result);
}

getTopLevelDomain().catch(console.error);
```
