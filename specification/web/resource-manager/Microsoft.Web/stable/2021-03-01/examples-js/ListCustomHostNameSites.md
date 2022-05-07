Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get custom hostnames under this subscription
 *
 * @summary Get custom hostnames under this subscription
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListCustomHostNameSites.json
 */
async function getCustomHostnamesUnderSubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.listCustomHostNameSites()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getCustomHostnamesUnderSubscription().catch(console.error);
```
