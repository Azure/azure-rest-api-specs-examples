Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get available Web app frameworks and their versions for location
 *
 * @summary Description for Get available Web app frameworks and their versions for location
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetWebAppStacksForLocation.json
 */
async function getLocationsWebAppStacks() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.provider.listWebAppStacksForLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getLocationsWebAppStacks().catch(console.error);
```
